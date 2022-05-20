#!/bin/bash

# generate grpc
bash cmd/auth/gRPC.sh

#build
go build -o ./tmp/auth ./cmd/auth/auth.go
