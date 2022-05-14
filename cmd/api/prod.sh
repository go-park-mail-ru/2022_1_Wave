#!/bin/bash
go get -u google.golang.org/protobuf
go get -u google.golang.org/grpc
go get -u golang.org/x/net/context
bash cmd/api/main.sh
bash -c "mkdir -p tmp && go build -o ./tmp/api cmd/api/main.go && ./tmp/api"

