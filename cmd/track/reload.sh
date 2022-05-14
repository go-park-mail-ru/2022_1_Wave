#!/bin/bash

# generate grpc
bash gRPC.sh

#build
go build -o ./tmp/track ./cmd/track/track.go
