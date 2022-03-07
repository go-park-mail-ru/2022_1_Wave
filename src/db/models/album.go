package models

import (
	"errors"
)

type Album struct {
	Id             uint64
	Title          string
	AuthorId       uint64
	CountLikes     uint64
	CountListening uint64
	Date           int
	CoverId        uint64
}

func (album *Album) CheckAlbum() error {
	//if album.Id < 0 {
	//	return errors.New(ErrorAlbumIdIsNegative)
	//}
	//
	//if album.AuthorId < 0 {
	//	return errors.New(ErrorAuthorIdIsNegative)
	//}
	//
	//if album.CountLikes < 0 {
	//	return errors.New(ErrorCountLikesIsNegative)
	//}
	//
	//if album.CountListening < 0 {
	//	return errors.New(ErrorCountListeningIsNegative)
	//}
	//
	//if album.CoverId < 0 {
	//	return errors.New(ErrorCoverIdIsNegative)
	//}

	if len(album.Title) > AlbumTitleLen {
		return errors.New(ErrorAlbumMaxTitleLen)
	}
	return nil
}
