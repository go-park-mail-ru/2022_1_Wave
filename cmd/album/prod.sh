#!/bin/bash
bash cmd/album/reload.sh
bash -c "mkdir -p tmp && go build -o ./tmp/album cmd/album/album.go && ./tmp/album"

