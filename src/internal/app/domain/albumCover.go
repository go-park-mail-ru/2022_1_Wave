package domain

import (
	"errors"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	utilsInterfaces "github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/interfaces"
)

type AlbumCover struct {
	Id     uint64 `json:"id" example:"1" db:"id"`
	Title  string `json:"title" example:"Mercury" db:"title"`
	Quote  string `json:"quote" example:"some phrases" db:"quote"`
	IsDark bool   `json:"isDark" example:"true" db:"is_dark"`
}

func (cover AlbumCover) GetId() uint64 {
	return cover.Id
}

func (cover AlbumCover) SetId(id uint64) (utilsInterfaces.Domain, error) {
	cover.Id = id
	//album.CoverId = album.Id
	return cover, nil
}

func (cover AlbumCover) Check() error {
	if cover.Id < 0 {
		return errors.New(constants.ErrorAlbumIdIsNegative)
	}

	if len(cover.Title) > constants.AlbumCoverTitleLen {
		return errors.New(constants.ErrorAlbumCoverMaxTitleLen)
	}

	if len(cover.Quote) > constants.AlbumCoverQuoteLen {
		return errors.New(constants.ErrorAlbumCoverMaxQuoteLen)
	}

	return nil
}

func (cover AlbumCover) GetCountListening() uint64 {
	// dummy func for look-like interface //
	return 0
}

func (cover AlbumCover) CastDomainToDataTransferObject(utilsInterfaces.Domain) (utilsInterfaces.DataTransfer, error) {
	return AlbumCoverDataTransfer{
		Title:  cover.Title,
		Quote:  cover.Quote,
		IsDark: cover.IsDark,
	}, nil
}

type AlbumCoverDataTransfer struct {
	Title  string `json:"title" example:"Mercury"`
	Quote  string `json:"quote" example:"some phrases"`
	IsDark bool   `json:"isDark" example:"true"`
}

func (cover AlbumCoverDataTransfer) CreateDataTransferFromInterface(data interface{}) (utilsInterfaces.DataTransfer, error) {
	temp := data.(map[string]interface{})
	return AlbumCoverDataTransfer{
		Title:  temp[constants.FieldTitle].(string),
		Quote:  temp[constants.FieldQuote].(string),
		IsDark: temp[constants.FieldIsDark].(bool),
	}, nil
}
