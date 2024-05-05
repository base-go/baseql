package baseqlpb

//go:generate sh -c "docker run -v `pwd`:/defs namely/protoc-all:1.11 -d . -l gogo && mv gen/pb-gogo/github.com/base-go/baseql/baseqlpb/* . && rm -rf gen"
