#!/bin/bash

# generate grpc
bash cmd/user/gRPC.sh

#build
go build -o ./tmp/user ./cmd/user/user.go
