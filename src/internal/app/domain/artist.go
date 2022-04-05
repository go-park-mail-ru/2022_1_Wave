package domain

import (
	"errors"
	"fmt"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/interfaces"
)

type Artist struct {
	Id             uint64 `json:"id" example:"6"`
	Name           string `json:"name" example:"Imagine Dragons"`
	PhotoId        uint64 `json:"photoId" example:"6"`
	CountFollowers uint64 `json:"countFollowers" example:"1001"`
	CountListening uint64 `json:"countListening" example:"7654"`
}

func (artist Artist) GetId() uint64 {
	return artist.Id
}

func (artist Artist) GetIdRef() *uint64 {
	return &artist.Id
}

func (artist Artist) SetId(id uint64) (utilsInterfaces.Domain, error) {
	artist.Id = id
	artist.PhotoId = artist.Id
	return artist, nil
}

func (artist Artist) Check() error {
	if artist.Id < 0 {
		return errors.New(constants.ErrorArtistPhotoIdIsNegative)
	}

	if artist.PhotoId < 0 {
		return errors.New(constants.ErrorArtistsMaxPhotoLinkLen)
	}

	if len(artist.Name) > constants.ArtistNameLen {
		return errors.New(constants.ErrorArtistMaxNameLen)
	}

	if artist.CountFollowers < 0 {
		return errors.New(constants.ErrorArtistCountFollowersIsNegative)
	}

	if artist.CountListening < 0 {
		return errors.New(constants.ErrorArtistCountListeningIsNegative)
	}

	return nil

}

func (artist Artist) GetCountListening() uint64 {
	return artist.CountListening
}

func (artist Artist) CastDomainToDataTransferObject(utilsInterfaces.Domain) (utilsInterfaces.DataTransfer, error) {
	return ArtistDataTransfer{
		Name:  artist.Name,
		Cover: constants.AssetsPrefix + constants.ArtistPreName + fmt.Sprint(artist.PhotoId) + constants.PngFormat,
	}, nil
}

type ArtistDataTransfer struct {
	Name  string `json:"name" example:"Mercury"`
	Cover string `json:"cover" example:"assets/artist_1.png"`
}

func (artist ArtistDataTransfer) CreateDataTransferFromInterface(data interface{}) (utilsInterfaces.DataTransfer, error) {
	temp := data.(map[string]interface{})
	return ArtistDataTransfer{
		Name:  temp["name"].(string),
		Cover: temp["cover"].(string),
	}, nil
}
