package domain

import (
	"errors"
	"fmt"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/interfaces"
)

type Album struct {
	Id             uint64 `json:"id" example:"8"`
	Title          string `json:"title" example:"Mercury"`
	ArtistId       uint64 `json:"authorId" example:"4"`
	CountLikes     uint64 `json:"countLikes" example:"54"`
	CountListening uint64 `json:"countListening" example:"15632"`
	Date           int    `json:"date" example:"0"`
	CoverId        uint64 `json:"coverId" example:"8"`
}

func (album Album) GetId() uint64 {
	return album.Id
}

func (album Album) SetId(id uint64) (utilsInterfaces.Domain, error) {
	album.Id = id
	return album, nil
}

func (album Album) Check() error {
	if album.Id < 0 {
		return errors.New(constants.ErrorAlbumIdIsNegative)
	}

	if album.CoverId < 0 {
		return errors.New(constants.ErrorAlbumCoverIdIsNegative)
	}

	if album.ArtistId < 0 {
		return errors.New(constants.ErrorArtistIdIsNegative)
	}

	if len(album.Title) > constants.AlbumTitleLen {
		return errors.New(constants.ErrorAlbumMaxTitleLen)
	}

	if album.CountLikes < 0 {
		return errors.New(constants.ErrorAlbumCountLikesIsNegative)
	}

	if album.CountListening < 0 {
		return errors.New(constants.ErrorAlbumCountListeningIsNegative)
	}

	return nil
}

func (album Album) GetCountListening() uint64 {
	return album.CountListening
}

func (album Album) CastDomainToDataTransferObject(artist utilsInterfaces.Domain) (utilsInterfaces.DataTransfer, error) {
	return AlbumDataTransfer{
		Title:  album.Title,
		Artist: artist.(Artist).Name,
		Cover:  constants.AssetsPrefix + constants.AlbumPreName + fmt.Sprint(album.CoverId) + constants.PngFormat,
	}, nil
}

type AlbumDataTransfer struct {
	Title  string `json:"title" example:"Mercury"`
	Artist string `json:"artist" example:"Hexed"`
	Cover  string `json:"cover" example:"assets/album_1.png"`
}

func (album AlbumDataTransfer) CreateDataTransferFromInterface(data interface{}) (utilsInterfaces.DataTransfer, error) {
	temp := data.(map[string]interface{})
	return AlbumDataTransfer{
		Title:  temp["title"].(string),
		Artist: temp["artist"].(string),
		Cover:  temp["cover"].(string),
	}, nil
}
