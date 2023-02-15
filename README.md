## About download modules

    go env -w GO111MODULE=on
    go env -w GOPROXY=https://goproxy.cn,direct

    # protoc 
    apt install protobuf-compiler

    # protoc-gen-go
    go install google.golang.org/protobuf/cmd/protoc-gen-go
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

    # generate source codes
    # [document](https://protobuf.dev/reference/go/go-generated/)
    cd /violin-settings/api/proto
    protoc --go_out=./gen --go_opt=paths=source_relative --go-grpc_out=./gen --go-grpc_opt=paths=source_relative notice_service.proto

## 

    http://127.0.0.1:80/violin-api/v1/settings?tenant_id=1