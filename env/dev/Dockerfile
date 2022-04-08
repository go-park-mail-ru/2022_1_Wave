FROM golang
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g cmd/api/main.go