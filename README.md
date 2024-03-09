Protocol buffer compiler

link: https://grpc.io/docs/protoc-installation/

    sudo apt install -y protobuf-compiler
    protoc --version

Dict

    - protoc: cli to execute protobuf compiler commands
    - .proto: file extension for service and message(req/res) definitions
    - grpc: remote procedure call (rpc) framework, that rely on defining a service method which by passing some arguments(input) it will execute some action. Like if you are executing a cli service through your terminal.
    - protobuff: interface description language IDL, the equivalent of a payload for common webservices like(json, xml)

Installation proto generator code for golang

    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

Generating structs based on .proto definition

    protoc file.proto -I=$GOPATH/src/project-name --go_out=$GOPATH/src/project-name --proto-path=$GOPATH/src/project-name/sub-path-to-proto-file

    ex.

    protoc app.proto -I=$GOPATH/src/go-grpc --go_out=$GOPATH/src/go-grpc --proto_path=$GOPATH/src/go-grpc/simple-server

Dict

    -I=Directory of the project, if not provided its default will be the current directory
    --go_out=Where do you want the generated code be placed

    The last argument is the path for your proto file which will be used as schema to generate your code

#####################################################

rpc - execution of a function located on another process/server which transmit data as binary, serialized by a data format such as protocol buff or thrift

Serializing and deserializing data using protocol buffers are faster than json, since it process binary directly. And since you transmite binary data over the network, the size of this payload is smaller compared to the size of a json payload.

grpc-gateway is a reverse-proxy for clients that can't communicate with grpc server directly, clients still implementing restful apis calls, while grpc converts its request to a grpc call

#################################################

Go grpc tools

go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

Generate go files from .proto

go:generate mkdir -p pb
go:generate protoc --go_opt=paths=$GOPATH/src/project-name --go_out=$GOPATH/src/project-name/pb  --proto-path=$GOPATH/src/project-name/sub-path-to-proto-file