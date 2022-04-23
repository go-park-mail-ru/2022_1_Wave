package domain

import (
	"fmt"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"gopkg.in/validator.v2"
)

type Artist struct {
	Id             int64  `json:"id" example:"6" db:"id" validate:"min=0,nonnil"`
	Name           string `json:"name" example:"Imagine Dragons" db:"name" validate:"max=256,nonnil"`
	CountLikes     int64  `json:"countLikes" example:"54" db:"count_likes" validate:"min=0,nonnil"`
	CountFollowers int64  `json:"countFollowers" example:"1001" db:"count_followers" validate:"min=0,nonnil"`
	CountListening int64  `json:"countListening" example:"7654" db:"count_listening" validate:"min=0,nonnil"`
}

type ArtistRepo interface {
	Insert(Artist) error
	Update(Artist) error
	Delete(int64) error
	SelectByID(int64) (*Artist, error)
	GetAll() ([]Artist, error)
	GetPopular() ([]Artist, error)
	GetLastId() (id int64, err error)
	//GetType() reflect.Type
	GetSize() (int64, error)
}

type ArtistDataTransfer struct {
	Id     int64               `json:"id" example:"1"`
	Name   string              `json:"name" example:"Mercury"`
	Cover  string              `json:"cover" example:"assets/artist_1.png"`
	Likes  int64               `json:"likes" example:"5"`
	Albums []AlbumDataTransfer `json:"albums"`
}

func (artist *Artist) GetId() int64 {
	return artist.Id
}

func (artist *Artist) GetIdRef() *int64 {
	return &artist.Id
}

func (artist *Artist) SetId(id int64) error {
	artist.Id = id
	return nil
}

func (artist *Artist) Check() error {
	return validator.Validate(artist)
}

func (artist *Artist) GetCountListening() int64 {
	return artist.CountListening
}

func (artist *Artist) CreatePath(fileFormat string) (string, error) {
	return constants.AssetsPrefix + constants.ArtistPreName + fmt.Sprint(artist.Id) + fileFormat, nil
}
