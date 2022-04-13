package domain

import (
	"gopkg.in/validator.v2"
)

type AlbumCover struct {
	Id int `json:"id" example:"1" db:"id" validate:"min=0,nonnil"`
	//Title  string `json:"title" example:"Mercury" db:"title" validate:"max=256,nonnil"`
	Quote  string `json:"quote" example:"some phrases" db:"quote" validate:"max=512,nonnil"`
	IsDark bool   `json:"isDark" example:"true" db:"is_dark" validate:"nonnil"`
}

type AlbumCoverDataTransfer struct {
	//Title  string `json:"title" example:"Mercury"`
	Quote  string `json:"quote" example:"some phrases"`
	IsDark bool   `json:"isDark" example:"true"`
}

func (cover *AlbumCover) GetId() int {
	return cover.Id
}

func (cover *AlbumCover) SetId(id int) error {
	cover.Id = id
	return nil
}

func (cover *AlbumCover) Check() error {
	return validator.Validate(cover)
}
