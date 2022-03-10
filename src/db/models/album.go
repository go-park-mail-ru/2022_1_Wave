package models

import (
	"errors"
)

type Album struct {
	Id             uint64 `json:"id" example:"777"`
	Title          string `json:"title" example:"Mercury"`
	AuthorId       uint64 `json:"authorId" example:"121"`
	CountLikes     uint64 `json:"countLikes" example:"54"`
	CountListening uint64 `json:"countListening" example:"15632"`
	Date           int    `json:"Date" example:"0"`
	CoverId        uint64 `json:"coverId" example:"254"`
}

func (album *Album) CheckAlbum() error {
	if len(album.Title) > AlbumTitleLen {
		return errors.New(ErrorAlbumMaxTitleLen)
	}
	return nil
}
