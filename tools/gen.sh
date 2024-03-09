#!/bin/bash

protoc $1 -I=$GOPATH/src/{project-name} --go_out=$GOPATH/src/{project-name} --go-grpc_out=$GOPATH/src/{project-name} --proto_path=$GOPATH/src/{project-name}/subpath-to-protofile