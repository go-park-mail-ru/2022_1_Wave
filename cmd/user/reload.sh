#!/bin/bash

# generate grpc
bash gRPC.sh

#build
go build -o ./tmp/user ./cmd/user/user.go
