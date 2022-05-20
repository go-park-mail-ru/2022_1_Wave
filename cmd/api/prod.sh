#!/bin/bash
go get -u google.golang.org/protobuf
go get -u google.golang.org/grpc
go get -u golang.org/x/net/context
bash gRPC.sh
go build -o ./tmp/wave.music ./cmd/api/main.go

