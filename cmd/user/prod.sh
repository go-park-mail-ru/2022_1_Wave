#!/bin/bash
bash cmd/user/reload.sh
bash -c "mkdir -p tmp && go build -o ./tmp/user cmd/user/user.go && ./tmp/user"

