#!/bin/bash

sudo apt install -y protobuf-compiler
protoc --version

# Requires Golang installed
# If these tools are not found, add to your $PATH $GOPATH/bin
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

# CLI tool to perform requests over server
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest