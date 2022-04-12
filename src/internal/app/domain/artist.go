package domain

import (
	"fmt"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/interfaces"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/tools/utils"
	"gopkg.in/validator.v2"
)

type Artist struct {
	Id             uint64 `json:"id" example:"6" db:"id" validate:"min=0,nonnil"`
	Name           string `json:"name" example:"Imagine Dragons" db:"name" validate:"max=256,nonnil"`
	CountFollowers uint64 `json:"countFollowers" example:"1001" db:"count_followers" validate:"min=0,nonnil"`
	CountListening uint64 `json:"countListening" example:"7654" db:"count_listening" validate:"min=0,nonnil"`
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
	return validator.Validate(artist)

	//if artist.Id < 0 {
	//	return errors.New(constants.ErrorArtistPhotoIdIsNegative)
	//}
	//
	//if len(artist.Name) > constants.ArtistNameLen {
	//	return errors.New(constants.ErrorArtistMaxNameLen)
	//}
	//
	//if artist.CountFollowers < 0 {
	//	return errors.New(constants.ErrorArtistCountFollowersIsNegative)
	//}
	//
	//if artist.CountListening < 0 {
	//	return errors.New(constants.ErrorArtistCountListeningIsNegative)
	//}
	//
	//return nil

}

func (artist Artist) GetCountListening() uint64 {
	return artist.CountListening
}

func (artist Artist) CreatePath(fileFormat string) (string, error) {
	return constants.AssetsPrefix + constants.ArtistPreName + fmt.Sprint(artist.Id) + fileFormat, nil
}

func (artist Artist) CastDomainToDataTransferObject(dom utilsInterfaces.Domain, args ...interface{}) (utilsInterfaces.DataTransfer, error) {
	pathToPhoto, err := artist.CreatePath(constants.PngFormat)
	if err != nil {
		return nil, err
	}

	albums := args[0].([]AlbumDataTransfer)
	return ArtistDataTransfer{
		Id:     artist.Id,
		Name:   artist.Name,
		Cover:  pathToPhoto,
		Albums: albums,
	}, nil
}

type ArtistDataTransfer struct {
	Id     uint64              `json:"id" example:"1"`
	Name   string              `json:"name" example:"Mercury"`
	Cover  string              `json:"cover" example:"assets/artist_1.png"`
	Albums []AlbumDataTransfer `json:"albums"`
}

func (artist ArtistDataTransfer) CreateDataTransferFromInterface(data interface{}) (utilsInterfaces.DataTransfer, error) {
	temp := data.(map[string]interface{})

	id, err := utils.ToUint64(temp["id"])

	if err != nil {
		return nil, err
	}
	return ArtistDataTransfer{
		Id:     id,
		Name:   temp["name"].(string),
		Cover:  temp["cover"].(string),
		Albums: temp["albums"].([]AlbumDataTransfer),
	}, nil
}
