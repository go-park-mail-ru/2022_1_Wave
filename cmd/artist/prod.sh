#!/bin/bash
bash cmd/artist/reload.sh
bash -c "mkdir -p tmp && go build -o ./tmp/artist cmd/artist/artist.go && ./tmp/artist"

