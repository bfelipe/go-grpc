#!/bin/bash

<<DOCS
    - \$1: name of protofile
    - -I: base directory of your project
    - --go_out: directory where generated go code will be placed
    - --proto_path: exactly path where your .proto file resides
DOCS

BASE_PATH=$GOPATH/src/{project-name}

protoc $1 -I=$BASE_PATH \
 --go_out=$BASE_PATH \
 --proto_path=$BASE_PATH/{subpath-to-protofile}
