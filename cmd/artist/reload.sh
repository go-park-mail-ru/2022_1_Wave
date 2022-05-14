#!/bin/bash

# generate grpc
bash gRPC.sh

#build
go build -o ./tmp/artist ./cmd/artist/artist.go
