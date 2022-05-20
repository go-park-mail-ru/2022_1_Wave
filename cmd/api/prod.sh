#!/bin/bash
go get -u google.golang.org/protobuf
go get -u google.golang.org/grpc
go get -u golang.org/x/net/context
bash gRPC.sh
swag init -g cmd/api/main.go
go build -o ./tmp/api ./cmd/api/main.go

