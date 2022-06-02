#!/bin/bash

go get -u github.com/gorilla/websocket
go get -u github.com/labstack/echo/v4
go get -u golang.org/x/net/context
go get -u github.com/google/uuid
go run ./websocket-server/cmd/main.go
