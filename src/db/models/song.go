package models

import (
	"errors"
)

type Song struct {
	Id             uint64 `json:"id" example:"777"`
	AlbumId        uint64 `json:"albumId" example:"143"`
	AuthorId       uint64 `json:"authorId" example:"121"`
	Title          string `json:"title" example:"Rain"`
	Duration       uint64 `json:"duration" example:"180"`
	Mp4            string `json:"mp4" example:"/public/songs/mp4/track1.mp4"`
	CountLikes     uint64 `json:"countLikes" example:"54"`
	CountListening uint64 `json:"countListening" example:"15632"`
}

func (song *Song) CheckSong() error {
	if len(song.Title) > SongTitleLen {
		return errors.New(ErrorSongMaxNameLen)
	}

	if len(song.Mp4) > SongLinkLen {
		return errors.New(ErrorSongMaxPhotoLinkLen)
	}

	return nil
}
