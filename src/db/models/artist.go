package models

import (
	"errors"
)

type Artist struct {
	Id             uint64 `json:"id" example:"43"`
	Name           string `json:"name" example:"Imagine Dragons"`
	Photo          string `json:"photo" example:"/public/artists/photo/imagineDragons.png"`
	CountFollowers uint64 `json:"countFollowers" example:"1001"`
	CountListening uint64 `json:"countListening" example:"7654"`
}

func (artist *Artist) CheckArtist() error {
	if len(artist.Name) > ArtistNameLen {
		return errors.New(ErrorArtistMaxNameLen)
	}

	if len(artist.Photo) > ArtistPhotoLinkLen {
		return errors.New(ErrorArtistsMaxPhotoLinkLen)
	}

	return nil
}
