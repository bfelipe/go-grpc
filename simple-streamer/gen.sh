#!/bin/bash

protoc $1 -I=$GOPATH/src/go-grpc --go_out=$GOPATH/src/go-grpc --go-grpc_out=$GOPATH/src/go-grpc --proto_path=$GOPATH/src/go-grpc/simple-streamer/pb