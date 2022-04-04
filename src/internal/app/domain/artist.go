package domain

import (
	"errors"
	"fmt"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/interfaces"
)

type Artist struct {
	Id             uint64 `json:"id" example:"43"`
	Name           string `json:"name" example:"Imagine Dragons"`
	Photo          string `json:"photo" example:"assets/artist_1.png"`
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
	return artist, nil
}

func (artist Artist) Check() error {
	if artist.Id < 0 {
		return errors.New(constants.ErrorArtistIdIsNegative)
	}

	if len(artist.Name) > constants.ArtistNameLen {
		return errors.New(constants.ErrorArtistMaxNameLen)
	}

	if len(artist.Photo) > constants.ArtistPhotoLinkLen {
		return errors.New(constants.ErrorArtistsMaxPhotoLinkLen)
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
		Cover: constants.AssetsPrefix + constants.ArtistPreName + fmt.Sprint(artist.Id) + constants.PngFormat,
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
