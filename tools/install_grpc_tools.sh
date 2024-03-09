#!/bin/bash

sudo apt install -y protobuf-compiler
protoc --version

# Requires Golang installed
# If these tools are not found, add to your $PATH $GOPATH/bin
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
