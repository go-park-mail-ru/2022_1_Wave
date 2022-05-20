#!/bin/bash

# generate grpc
bash cmd/album/gRPC.sh

#build
go build -o ./tmp/album ./cmd/album/album.go
