#!/bin/bash

# generate grpc
bash cmd/playlist/gRPC.sh

#build
go build -o ./tmp/playlist ./cmd/playlist/playlist.go
