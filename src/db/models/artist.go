package models

import (
	"errors"
)

type Artist struct {
	Id             uint64
	Name           string
	Photo          string
	CountFollowers uint64
	CountListening uint64
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
