package domain

import (
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	utilsInterfaces "github.com/go-park-mail-ru/2022_1_Wave/internal/app/interfaces"
	"gopkg.in/validator.v2"
)

type AlbumCover struct {
	Id uint64 `json:"id" example:"1" db:"id" validate:"min=0,nonnil"`
	//Title  string `json:"title" example:"Mercury" db:"title" validate:"max=256,nonnil"`
	Quote  string `json:"quote" example:"some phrases" db:"quote" validate:"max=512,nonnil"`
	IsDark bool   `json:"isDark" example:"true" db:"is_dark" validate:"nonnil"`
}

func (cover AlbumCover) GetId() uint64 {
	return cover.Id
}

func (cover AlbumCover) SetId(id uint64) (utilsInterfaces.Domain, error) {
	cover.Id = id
	return cover, nil
}

func (cover AlbumCover) Check() error {
	return validator.Validate(cover)
	//}
	//
	//if cover.Id < 0 {
	//	return errors.New(constants.ErrorAlbumIdIsNegative)
	//}
	//
	//if len(cover.Title) > constants.AlbumCoverTitleLen {
	//	return errors.New(constants.ErrorAlbumCoverMaxTitleLen)
	//}
	//
	//if len(cover.Quote) > constants.AlbumCoverQuoteLen {
	//	return errors.New(constants.ErrorAlbumCoverMaxQuoteLen)
	//}
	//
	//return nil
}

func (cover AlbumCover) GetCountListening() uint64 {
	// dummy func for look-like interface //
	return 0
}

func (cover AlbumCover) CastDomainToDataTransferObject(dom utilsInterfaces.Domain, args ...interface{}) (utilsInterfaces.DataTransfer, error) {
	return AlbumCoverDataTransfer{
		//Title:  cover.Title,
		Quote:  cover.Quote,
		IsDark: cover.IsDark,
	}, nil
}

type AlbumCoverDataTransfer struct {
	//Title  string `json:"title" example:"Mercury"`
	Quote  string `json:"quote" example:"some phrases"`
	IsDark bool   `json:"isDark" example:"true"`
}

func (cover AlbumCoverDataTransfer) CreateDataTransferFromInterface(data interface{}) (utilsInterfaces.DataTransfer, error) {
	temp := data.(map[string]interface{})
	return AlbumCoverDataTransfer{
		//Title:  temp[constants.FieldTitle].(string),
		Quote:  temp[constants.FieldQuote].(string),
		IsDark: temp[constants.FieldIsDark].(bool),
	}, nil
}
