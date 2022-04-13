package domain

import (
	"fmt"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"gopkg.in/validator.v2"
)

type Artist struct {
	Id             int    `json:"id" example:"6" db:"id" validate:"min=0,nonnil"`
	Name           string `json:"name" example:"Imagine Dragons" db:"name" validate:"max=256,nonnil"`
	CountLikes     int    `json:"countLikes" example:"54" db:"count_likes" validate:"min=0,nonnil"`
	CountFollowers int    `json:"countFollowers" example:"1001" db:"count_followers" validate:"min=0,nonnil"`
	CountListening int    `json:"countListening" example:"7654" db:"count_listening" validate:"min=0,nonnil"`
}

type ArtistDataTransfer struct {
	Id     int                 `json:"id" example:"1"`
	Name   string              `json:"name" example:"Mercury"`
	Cover  string              `json:"cover" example:"assets/artist_1.png"`
	Likes  int                 `json:"likes" example:"5"`
	Albums []AlbumDataTransfer `json:"albums"`
}

func (artist *Artist) GetId() int {
	return artist.Id
}

func (artist *Artist) GetIdRef() *int {
	return &artist.Id
}

func (artist *Artist) SetId(id int) error {
	artist.Id = id
	return nil
}

func (artist *Artist) Check() error {
	return validator.Validate(artist)
}

func (artist *Artist) GetCountListening() int {
	return artist.CountListening
}

func (artist *Artist) CreatePath(fileFormat string) (string, error) {
	return constants.AssetsPrefix + constants.ArtistPreName + fmt.Sprint(artist.Id) + fileFormat, nil
}
