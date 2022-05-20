#!/bin/bash

# generate grpc
cd internal/microservices || return
trackProto="./track/trackProto/"
protoFiles=$(find . -regex ".*\(album\|gateway\)+\.proto" )
# generate stubs and skeletons
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative -I='.' -I=$trackProto $protoFiles
# inject tags
find . -iname 'album.pb.go' -exec protoc-go-inject-tag -input={} \;
cd ../.. || return
