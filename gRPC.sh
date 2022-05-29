#!/bin/bash

# TODO need to run chmod +x gRPC.sh

cd internal/microservices || return
trackProto="./track/trackProto/"
protoFiles=$(find . -iname '*.proto')
# generate stubs and skeletons
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       -I=. -I=$trackProto $protoFiles \
# inject tags
find . -iname '*.pb.go' -exec protoc-go-inject-tag -input={} \;
cd ../.. || return