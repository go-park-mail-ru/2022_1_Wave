#!/bin/bash

# generate grpc
bash gRPC.sh

# swag
swag init -g cmd/api/main.go

#build
go build -o ./tmp/wave.music ./cmd/api/main.go
