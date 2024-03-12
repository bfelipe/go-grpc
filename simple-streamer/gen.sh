#!/bin/bash

BASE_PATH=$GOPATH/src/go-grpc

protoc $1 -I=$BASE_PATH \
 --go_out=$BASE_PATH \
 --proto_path=$BASE_PATH/simple-streamer/pb