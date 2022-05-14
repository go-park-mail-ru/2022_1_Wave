#!/bin/bash

# generate grpc
bash gRPC.sh

#build
go build -o ./tmp/album ./cmd/album/album.go
