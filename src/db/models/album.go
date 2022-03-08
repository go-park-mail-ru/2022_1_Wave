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
	Cover          string `json:"cover" example:"assets/album_1.png"`
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
	//if album.Cover < 0 {
	//	return errors.New(ErrorCoverIdIsNegative)
	//}

	if len(album.Title) > AlbumTitleLen {
		return errors.New(ErrorAlbumMaxTitleLen)
	}
	return nil
}
