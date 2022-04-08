package domain

import (
	"errors"
	"fmt"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/interfaces"
)

//PhotoId        uint64 `json:"photoId" example:"6" db:"photo_id"`

type Artist struct {
	Id   uint64 `json:"id" example:"6" db:"id"`
	Name string `json:"name" example:"Imagine Dragons" db:"name"`
	//PhotoId        uint64 `db:"photo_id"`
	CountFollowers uint64 `json:"countFollowers" example:"1001" db:"count_followers"`
	CountListening uint64 `json:"countListening" example:"7654" db:"count_listening"`
}

func (artist Artist) GetId() uint64 {
	return artist.Id
}

func (artist Artist) GetIdRef() *uint64 {
	return &artist.Id
}

func (artist Artist) SetId(id uint64) (utilsInterfaces.Domain, error) {
	artist.Id = id
	//artist.PhotoId = artist.Id
	return artist, nil
}

func (artist Artist) Check() error {
	if artist.Id < 0 {
		return errors.New(constants.ErrorArtistPhotoIdIsNegative)
	}

	//if artist.PhotoId < 0 {
	//	return errors.New(constants.ErrorArtistsMaxPhotoLinkLen)
	//}

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

func (artist Artist) CreatePath(fileFormat string) (string, error) {
	return constants.AssetsPrefix + constants.ArtistPreName + fmt.Sprint(artist.Id) + fileFormat, nil
}

func (artist Artist) CastDomainToDataTransferObject(utilsInterfaces.Domain) (utilsInterfaces.DataTransfer, error) {
	pathToPhoto, err := artist.CreatePath(constants.PngFormat)
	if err != nil {
		return nil, err
	}
	return ArtistDataTransfer{
		Name:  artist.Name,
		Cover: pathToPhoto,
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
