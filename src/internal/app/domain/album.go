package domain

import (
	"fmt"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/interfaces"
	"gopkg.in/validator.v2"
)

type Album struct {
	Id             uint64 `json:"id" example:"8" db:"id" validate:"min=0,nonnil"`
	Title          string `json:"title" example:"Mercury" db:"title" validate:"max=256,nonnil"`
	ArtistId       uint64 `json:"artistId" example:"4" db:"artist_id" validate:"min=0,nonnil"`
	CountLikes     uint64 `json:"countLikes" example:"54" db:"count_likes" validate:"min=0,nonnil"`
	CountListening uint64 `json:"countListening" example:"15632" db:"count_listening" validate:"min=0,nonnil"`
	Date           int64  `json:"date" example:"0" db:"date,nonnil"`
}

func (album Album) GetId() uint64 {
	return album.Id
}

func (album Album) SetId(id uint64) (utilsInterfaces.Domain, error) {
	album.Id = id
	return album, nil
}

func (album Album) Check() error {
	return validator.Validate(album)

	//if album.Id < 0 {
	//	return errors.New(constants.ErrorAlbumIdIsNegative)
	//}
	//
	////if album.CoverId < 0 {
	////	return errors.New(constants.ErrorAlbumCoverIdIsNegative)
	////}
	//
	//if album.ArtistId < 0 {
	//	return errors.New(constants.ErrorArtistIdIsNegative)
	//}
	//
	//if len(album.Title) > constants.AlbumTitleLen {
	//	return errors.New(constants.ErrorAlbumMaxTitleLen)
	//}
	//
	//if album.CountLikes < 0 {
	//	return errors.New(constants.ErrorAlbumCountLikesIsNegative)
	//}
	//
	//if album.CountListening < 0 {
	//	return errors.New(constants.ErrorAlbumCountListeningIsNegative)
	//}
	//
	//return nil
}

func (album Album) GetCountListening() uint64 {
	return album.CountListening
}

func (album Album) CreatePath(fileFormat string) (string, error) {
	return constants.AssetsPrefix + constants.AlbumPreName + fmt.Sprint(album.Id) + fileFormat, nil
}

func (album Album) CastDomainToDataTransferObject(artist utilsInterfaces.Domain) (utilsInterfaces.DataTransfer, error) {
	pathToCover, err := album.CreatePath(constants.PngFormat)
	if err != nil {
		return nil, err
	}
	return AlbumDataTransfer{
		Title:  album.Title,
		Artist: artist.(Artist).Name,
		Cover:  pathToCover,
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
