package livesql

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/base-go/baseql/internal/fields"
	"github.com/base-go/baseql/sqlgen"
	"github.com/base-go/baseql/baseqlpb"
)

// valueToField converts a driver.Value into a baseqlpb.Field.
// driver.Value must be one of the following:
//   int64
//   float64
//   bool
//   []byte
//   string
//   time.Time
func valueToField(value driver.Value) (*baseqlpb.Field, error) {
	switch column := value.(type) {
	case nil:
		return &baseqlpb.Field{Kind: baseqlpb.FieldKind_Null}, nil
	case int64:
		return &baseqlpb.Field{Kind: baseqlpb.FieldKind_Int, Value: &baseqlpb.Field_Int{Int: column}}, nil
	case float64:
		return &baseqlpb.Field{Kind: baseqlpb.FieldKind_Float64, Value: &baseqlpb.Field_Float64{Float64: column}}, nil
	case bool:
		return &baseqlpb.Field{Kind: baseqlpb.FieldKind_Bool, Value: &baseqlpb.Field_Bool{Bool: column}}, nil
	case []byte:
		return &baseqlpb.Field{Kind: baseqlpb.FieldKind_Bytes, Value: &baseqlpb.Field_Bytes{Bytes: column}}, nil
	case string:
		return &baseqlpb.Field{Kind: baseqlpb.FieldKind_String, Value: &baseqlpb.Field_String_{String_: column}}, nil
	case time.Time:
		return &baseqlpb.Field{Kind: baseqlpb.FieldKind_Time, Value: &baseqlpb.Field_Time{Time: &column}}, nil
	default:
		return nil, fmt.Errorf("unknown type %s", reflect.TypeOf(column))
	}
}

// FieldToValue converts baseqlpb.Field to driver.Value.
func FieldToValue(field *baseqlpb.Field) (driver.Value, error) {
	switch field.Kind {
	case baseqlpb.FieldKind_Null:
		return nil, nil
	case baseqlpb.FieldKind_Bool:
		return field.GetBool(), nil
	case baseqlpb.FieldKind_Int:
		return field.GetInt(), nil
	case baseqlpb.FieldKind_Uint:
		return field.GetUint(), nil
	case baseqlpb.FieldKind_String:
		return field.GetString_(), nil
	case baseqlpb.FieldKind_Bytes:
		return field.GetBytes(), nil
	case baseqlpb.FieldKind_Float64:
		return field.GetFloat64(), nil
	case baseqlpb.FieldKind_Time:
		ptr := field.GetTime()
		if ptr == nil {
			return time.Time{}, nil
		}
		return *ptr, nil
	default:
		return nil, fmt.Errorf("unknown kind %s", field.Kind.String())
	}
}

// FilterToProto takes a sqlgen.Filter, runs Valuer on each filter value, and returns a baseqlpb.SQLFilter.
func FilterToProto(schema *sqlgen.Schema, tableName string, filter sqlgen.Filter) (*baseqlpb.SQLFilter, error) {
	table, ok := schema.ByName[tableName]
	if !ok {
		return nil, fmt.Errorf("unknown table: %s", tableName)
	}

	if filter == nil {
		return &baseqlpb.SQLFilter{Table: tableName}, nil
	}

	fields := make(map[string]*baseqlpb.Field, len(filter))
	for col, val := range filter {
		column, ok := table.ColumnsByName[col]
		if !ok {
			return nil, fmt.Errorf("unknown column %s", col)
		}

		val, err := column.Descriptor.Valuer(reflect.ValueOf(val)).Value()
		if err != nil {
			return nil, err
		}

		field, err := valueToField(val)
		if err != nil {
			return nil, err
		}
		fields[col] = field
	}
	return &baseqlpb.SQLFilter{Table: tableName, Fields: fields}, nil
}

// FilterFromProto takes a baseqlpb.SQLFilter, runs Scanner on each field value, and returns a sqlgen.Filter.
func FilterFromProto(schema *sqlgen.Schema, proto *baseqlpb.SQLFilter) (string, sqlgen.Filter, error) {
	table, ok := schema.ByName[proto.Table]
	if !ok {
		return "", nil, fmt.Errorf("unknown table: %s", proto.Table)
	}

	scanners := table.Scanners.Get().([]interface{})
	defer table.Scanners.Put(scanners)

	filter := make(sqlgen.Filter, len(proto.Fields))
	for col, field := range proto.Fields {
		val, err := FieldToValue(field)
		if err != nil {
			return "", nil, err
		}

		column, ok := table.ColumnsByName[col]
		if !ok {
			return "", nil, fmt.Errorf("unknown column %s", col)
		}

		if !column.Descriptor.Ptr && val == nil {
			return "", nil, errors.New("cannot unmarshal nil into non-pointer type")
		}

		scanner := scanners[column.Order].(*fields.Scanner)

		// target is always a pointer.
		var target, ptrptr reflect.Value
		if column.Descriptor.Ptr {
			// We need to hold onto this pointer-pointer in order to make the value addressable.
			ptrptr = reflect.New(reflect.PtrTo(column.Descriptor.Type))
			target = ptrptr.Elem()
		} else {
			target = reflect.New(column.Descriptor.Type)
		}
		scanner.Target(target)

		if err := scanner.Scan(val); err != nil {
			return "", nil, err
		}

		if column.Descriptor.Ptr {
			filter[col] = target.Interface()
		} else {
			// Dereference pointer if column type is not a pointer.
			filter[col] = target.Elem().Interface()
		}
	}
	return proto.Table, filter, nil
}
