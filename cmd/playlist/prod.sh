#!/bin/bash
bash cmd/playlist/reload.sh
bash -c "mkdir -p tmp && go build -o ./tmp/playlist cmd/playlist/playlist.go && ./tmp/playlist"

