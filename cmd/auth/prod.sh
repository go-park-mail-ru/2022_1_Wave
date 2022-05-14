#!/bin/bash
bash cmd/auth/reload.sh
bash -c "mkdir -p tmp && go build -o ./tmp/auth cmd/auth/auth.go && ./tmp/auth"

