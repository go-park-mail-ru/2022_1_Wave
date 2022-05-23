#!/bin/bash

# generate grpc and swag
bash generateAll.sh

#build
go build -o ./tmp/run.album ./cmd/album/album.go
go build -o ./tmp/run.artist ./cmd/artist/artist.go
go build -o ./tmp/run.auth ./cmd/auth/auth.go
go build -o ./tmp/run.playlist ./cmd/playlist/playlist.go
go build -o ./tmp/run.track ./cmd/track/track.go
go build -o ./tmp/run.user ./cmd/user/user.go
go build -o ./tmp/wave.music ./cmd/api/main.go


