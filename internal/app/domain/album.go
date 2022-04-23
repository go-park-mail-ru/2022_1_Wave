package domain

import (
	"fmt"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"gopkg.in/validator.v2"
)

type Album struct {
	Id             int64  `json:"id" example:"8" db:"id" validate:"min=0,nonnil"`
	Title          string `json:"title" example:"Mercury" db:"title" validate:"max=256,nonnil"`
	ArtistId       int64  `json:"artistId" example:"4" db:"artist_id" validate:"min=0,nonnil"`
	CountLikes     int64  `json:"countLikes" example:"54" db:"count_likes" validate:"min=0,nonnil"`
	CountListening int64  `json:"countListening" example:"15632" db:"count_listening" validate:"min=0,nonnil"`
	Date           int64  `json:"date" example:"0" db:"date,nonnil"`
}

type AlbumRepo interface {
	Insert(Album) error
	Update(Album) error
	Delete(int64) error
	SelectByID(int64) (*Album, error)
	GetAll() ([]Album, error)
	GetPopular() ([]Album, error)
	GetLastId() (id int64, err error)
	//GetType() reflect.Type
	GetSize() (int64, error)
	GetAlbumsFromArtist(artist int64) ([]Album, error)
}

type AlbumDataTransfer struct {
	Id     int64               `json:"id" example:"1"`
	Title  string              `json:"title" example:"Mercury"`
	Artist string              `json:"artist" example:"Hexed"`
	Cover  string              `json:"cover" example:"assets/album_1.png"`
	Tracks []TrackDataTransfer `json:"tracks"`
}

func (album *Album) CastToDtoWithoutArtistNameAndTracks() (*AlbumDataTransfer, error) {
	cover, err := album.CreatePath(constants.PngFormat)
	if err != nil {
		return nil, err
	}

	return &AlbumDataTransfer{
		Id:     album.Id,
		Title:  album.Title,
		Artist: "",
		Cover:  cover,
		Tracks: nil,
	}, nil

}

func (album *Album) GetId() int64 {
	return album.Id
}

func (album *Album) SetId(id int64) error {
	album.Id = id
	return nil
}

func (album *Album) Check() error {
	return validator.Validate(album)
}

func (album *Album) GetCountListening() int64 {
	return album.CountListening
}

func (album *Album) CreatePath(fileFormat string) (string, error) {
	return constants.AssetsPrefix + constants.AlbumPreName + fmt.Sprint(album.Id) + fileFormat, nil
}
