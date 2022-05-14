#!/bin/bash

# generate grpc
bash gRPC.sh

#build
go build -o ./tmp/playlist ./cmd/playlist/playlist.go
