package domain

import (
	"fmt"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"gopkg.in/validator.v2"
)

type Album struct {
	Id             int    `json:"id" example:"8" db:"id" validate:"min=0,nonnil"`
	Title          string `json:"title" example:"Mercury" db:"title" validate:"max=256,nonnil"`
	ArtistId       int    `json:"artistId" example:"4" db:"artist_id" validate:"min=0,nonnil"`
	CountLikes     int    `json:"countLikes" example:"54" db:"count_likes" validate:"min=0,nonnil"`
	CountListening int    `json:"countListening" example:"15632" db:"count_listening" validate:"min=0,nonnil"`
	Date           int64  `json:"date" example:"0" db:"date,nonnil"`
}

type AlbumRepo interface {
	Insert(Album) error
	Update(Album) error
	Delete(int) error
	SelectByID(int) (*Album, error)
	GetAll() ([]Album, error)
	GetPopular() ([]Album, error)
	GetLastId() (id int, err error)
	//GetType() reflect.Type
	GetSize() (int, error)

	//todo пока кастыль
	GetAlbumsFromArtist(artist int) ([]Album, error)
}

type AlbumDataTransfer struct {
	Id     int                 `json:"id" example:"1"`
	Title  string              `json:"title" example:"Mercury"`
	Artist string              `json:"artist" example:"Hexed"`
	Cover  string              `json:"cover" example:"assets/album_1.png"`
	Tracks []TrackDataTransfer `json:"tracks"`
}

func (album *Album) GetId() int {
	return album.Id
}

func (album *Album) SetId(id int) error {
	album.Id = id
	return nil
}

func (album *Album) Check() error {
	return validator.Validate(album)
}

func (album *Album) GetCountListening() int {
	return album.CountListening
}

func (album *Album) CreatePath(fileFormat string) (string, error) {
	return constants.AssetsPrefix + constants.AlbumPreName + fmt.Sprint(album.Id) + fileFormat, nil
}
