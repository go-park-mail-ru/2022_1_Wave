#!/bin/bash
bash cmd/track/reload.sh
bash -c "mkdir -p tmp && go build -o ./tmp/track cmd/track/track.go && ./tmp/track"

