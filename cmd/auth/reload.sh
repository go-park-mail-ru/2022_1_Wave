#!/bin/bash

# generate grpc
bash gRPC.sh

#build
go build -o ./tmp/auth ./cmd/auth/auth.go
