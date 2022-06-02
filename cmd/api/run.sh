#!/bin/bash
go get -u google.golang.org/protobuf
go get -u google.golang.org/grpc
go get -u golang.org/x/net/context
air -c cmd/api/.air.toml

