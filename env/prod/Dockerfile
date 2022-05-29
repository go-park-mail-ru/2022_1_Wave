FROM golang:latest
RUN apt-get update -y
RUN apt-get upgrade -y
RUN apt install -y protobuf-compiler
RUN apt install -y golang-goprotobuf-dev
RUN go install github.com/vektra/mockery/v2@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install github.com/cosmtrek/air@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install github.com/favadi/protoc-go-inject-tag@latest
