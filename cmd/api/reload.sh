#!/bin/bash

# generate grpc
bash cmd/api/gRPC.sh

# swag
swag init -g cmd/api/main.go

#build
go build -o ./tmp/wave.music ./cmd/api/main.go
