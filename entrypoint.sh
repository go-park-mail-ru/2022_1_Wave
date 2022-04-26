#!/bin/bash
go version
go install github.com/swaggo/swag/cmd/swag@latest
go install github.com/cosmtrek/air@latest
air
