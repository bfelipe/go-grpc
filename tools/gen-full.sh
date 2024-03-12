#!/bin/bash

<<DOCS
    - \$1: name of protofile
    - -I: base directory of your project
    - --go_out: directory where generated go code will be placed
    - --grpc-gateway_out: directory where generated grpc gateway code will be placed
    - --grpc-gateway_opt=logtostderr=true: enables log during code generation for debug purpose
    - --grpc-gateway_opt=generate_unbound_methods=true: generate unbound methods for flexible request handling
    - --proto_path: exactly path where your .proto file resides
DOCS

BASE_PATH=$GOPATH/src/{project-name}

protoc $1 -I=$BASE_PATH \
 --go_out=$BASE_PATH \
#  --grpc-gateway_out=$BASE_PATH \
#  --grpc-gateway_opt=logtostderr=true \
#  --grpc-gateway_opt=generate_unbound_methods=true \
 --proto_path=$BASE_PATH/{subpath-to-protofile}
