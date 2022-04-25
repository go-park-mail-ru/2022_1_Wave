#!/bin/bash

# generate grpc
cd internal/app/microservices || return
trackProto="./track/trackProto/"
protoFiles=$(find . -iname '*.proto')
# generate stubs and skeletons
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative -I='.' -I=$trackProto $protoFiles
# inject tags
find . -iname '*.pb.go' -exec protoc-go-inject-tag -input={} \;
cd ../../.. || return

# swag
swag init -g cmd/api/main.go

#build
go build -o ./tmp/wave.music ./cmd/api/main.go
