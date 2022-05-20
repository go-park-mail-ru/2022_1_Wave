#!/bin/bash

# generate grpc
bash cmd/track/gRPC.sh

#build
go build -o ./tmp/track ./cmd/track/track.go
