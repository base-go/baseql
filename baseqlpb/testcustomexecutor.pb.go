// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: testcustomexecutor.proto

package baseqlpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CustomExecutorRequest struct {
	Request *ExecuteRequest `protobuf:"bytes,1,opt,name=request" json:"request,omitempty"`
	Token   string          `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (m *CustomExecutorRequest) Reset()         { *m = CustomExecutorRequest{} }
func (m *CustomExecutorRequest) String() string { return proto.CompactTextString(m) }
func (*CustomExecutorRequest) ProtoMessage()    {}
func (*CustomExecutorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorTestcustomexecutor, []int{0}
}

func (m *CustomExecutorRequest) GetRequest() *ExecuteRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *CustomExecutorRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type CustomExecutorResponse struct {
	Response *ExecuteResponse `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
	Token    string           `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (m *CustomExecutorResponse) Reset()         { *m = CustomExecutorResponse{} }
func (m *CustomExecutorResponse) String() string { return proto.CompactTextString(m) }
func (*CustomExecutorResponse) ProtoMessage()    {}
func (*CustomExecutorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorTestcustomexecutor, []int{1}
}

func (m *CustomExecutorResponse) GetResponse() *ExecuteResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *CustomExecutorResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*CustomExecutorRequest)(nil), "baseqlpb.CustomExecutorRequest")
	proto.RegisterType((*CustomExecutorResponse)(nil), "baseqlpb.CustomExecutorResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CustomExecutor service

type CustomExecutorClient interface {
	Execute(ctx context.Context, in *CustomExecutorRequest, opts ...grpc.CallOption) (*CustomExecutorResponse, error)
}

type customExecutorClient struct {
	cc *grpc.ClientConn
}

func NewCustomExecutorClient(cc *grpc.ClientConn) CustomExecutorClient {
	return &customExecutorClient{cc}
}

func (c *customExecutorClient) Execute(ctx context.Context, in *CustomExecutorRequest, opts ...grpc.CallOption) (*CustomExecutorResponse, error) {
	out := new(CustomExecutorResponse)
	err := grpc.Invoke(ctx, "/baseqlpb.CustomExecutor/Execute", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CustomExecutor service

type CustomExecutorServer interface {
	Execute(context.Context, *CustomExecutorRequest) (*CustomExecutorResponse, error)
}

func RegisterCustomExecutorServer(s *grpc.Server, srv CustomExecutorServer) {
	s.RegisterService(&_CustomExecutor_serviceDesc, srv)
}

func _CustomExecutor_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomExecutorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomExecutorServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/baseqlpb.CustomExecutor/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomExecutorServer).Execute(ctx, req.(*CustomExecutorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomExecutor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "baseqlpb.CustomExecutor",
	HandlerType: (*CustomExecutorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _CustomExecutor_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "testcustomexecutor.proto",
}

func (m *CustomExecutorRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CustomExecutorRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Request != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTestcustomexecutor(dAtA, i, uint64(m.Request.Size()))
		n1, err := m.Request.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Token) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTestcustomexecutor(dAtA, i, uint64(len(m.Token)))
		i += copy(dAtA[i:], m.Token)
	}
	return i, nil
}

func (m *CustomExecutorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CustomExecutorResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Response != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTestcustomexecutor(dAtA, i, uint64(m.Response.Size()))
		n2, err := m.Response.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Token) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTestcustomexecutor(dAtA, i, uint64(len(m.Token)))
		i += copy(dAtA[i:], m.Token)
	}
	return i, nil
}

func encodeVarintTestcustomexecutor(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CustomExecutorRequest) Size() (n int) {
	var l int
	_ = l
	if m.Request != nil {
		l = m.Request.Size()
		n += 1 + l + sovTestcustomexecutor(uint64(l))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovTestcustomexecutor(uint64(l))
	}
	return n
}

func (m *CustomExecutorResponse) Size() (n int) {
	var l int
	_ = l
	if m.Response != nil {
		l = m.Response.Size()
		n += 1 + l + sovTestcustomexecutor(uint64(l))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovTestcustomexecutor(uint64(l))
	}
	return n
}

func sovTestcustomexecutor(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTestcustomexecutor(x uint64) (n int) {
	return sovTestcustomexecutor(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CustomExecutorRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTestcustomexecutor
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CustomExecutorRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CustomExecutorRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Request", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestcustomexecutor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTestcustomexecutor
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Request == nil {
				m.Request = &ExecuteRequest{}
			}
			if err := m.Request.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestcustomexecutor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTestcustomexecutor
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTestcustomexecutor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTestcustomexecutor
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CustomExecutorResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTestcustomexecutor
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CustomExecutorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CustomExecutorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Response", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestcustomexecutor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTestcustomexecutor
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Response == nil {
				m.Response = &ExecuteResponse{}
			}
			if err := m.Response.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestcustomexecutor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTestcustomexecutor
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTestcustomexecutor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTestcustomexecutor
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTestcustomexecutor(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTestcustomexecutor
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTestcustomexecutor
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTestcustomexecutor
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTestcustomexecutor
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTestcustomexecutor
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTestcustomexecutor(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTestcustomexecutor = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTestcustomexecutor   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("testcustomexecutor.proto", fileDescriptorTestcustomexecutor) }

var fileDescriptorTestcustomexecutor = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x14, 0x45, 0x8d, 0xa0, 0xe3, 0x44, 0x10, 0x09, 0x2a, 0xb5, 0x8b, 0x52, 0xbb, 0x90, 0xd9, 0xd8,
	0xc2, 0x8c, 0xf8, 0x01, 0x8a, 0x7b, 0xe9, 0xd2, 0x5d, 0xd3, 0x79, 0x6d, 0x07, 0x69, 0x5f, 0x27,
	0x79, 0x01, 0x3f, 0xd1, 0xa5, 0x9f, 0x20, 0xfd, 0x12, 0x21, 0x89, 0xa5, 0x4a, 0x71, 0x77, 0x6f,
	0x39, 0xbd, 0x87, 0x3c, 0x1e, 0x10, 0x68, 0x2a, 0x8d, 0x26, 0x6c, 0xe1, 0x1d, 0x4a, 0x43, 0xa8,
	0xd2, 0x5e, 0x21, 0xa1, 0x58, 0x52, 0x63, 0xba, 0x2d, 0xa8, 0x5e, 0x86, 0x77, 0xf5, 0x8e, 0x1a,
	0x23, 0xd3, 0x12, 0xdb, 0xac, 0xc6, 0x1a, 0x33, 0x4b, 0x48, 0x53, 0xd9, 0x66, 0x8b, 0x4d, 0xee,
	0xcf, 0xf0, 0xbc, 0x82, 0x2d, 0xa8, 0x82, 0x76, 0xd8, 0xb9, 0x2f, 0x89, 0xe4, 0x97, 0x4f, 0xd6,
	0xf1, 0xec, 0x1d, 0x39, 0xec, 0x0d, 0x68, 0x12, 0x1b, 0xbe, 0x50, 0x2e, 0x06, 0x2c, 0x66, 0xab,
	0xd3, 0xf5, 0x75, 0x3a, 0x6a, 0x53, 0x07, 0x83, 0x67, 0xf3, 0x1f, 0x52, 0x5c, 0xf0, 0x23, 0xc2,
	0x37, 0xe8, 0x82, 0xc3, 0x98, 0xad, 0x96, 0xb9, 0x2b, 0x49, 0xc5, 0xaf, 0xfe, 0x3a, 0x74, 0x8f,
	0x9d, 0x06, 0xf1, 0xc0, 0x4f, 0x94, 0xcf, 0xde, 0x12, 0xce, 0x59, 0x1c, 0x91, 0x8f, 0xec, 0xbc,
	0x67, 0x2d, 0xf9, 0xd9, 0x6f, 0x8f, 0x78, 0xe1, 0x0b, 0x3f, 0x22, 0xe2, 0xc9, 0xf0, 0xec, 0x8b,
	0xc3, 0x9b, 0x7f, 0x08, 0xe7, 0x4d, 0x0e, 0x1e, 0xef, 0x3f, 0x86, 0x88, 0x7d, 0x0e, 0x11, 0xfb,
	0x1a, 0x22, 0xf6, 0x7a, 0x3b, 0x39, 0xbf, 0x2e, 0x5a, 0x5d, 0xa8, 0xa2, 0xd9, 0x67, 0x7e, 0x26,
	0x1b, 0xe7, 0xe4, 0xb1, 0x3d, 0xf6, 0xe6, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x93, 0x93, 0x7e,
	0xd4, 0x01, 0x00, 0x00,
}
