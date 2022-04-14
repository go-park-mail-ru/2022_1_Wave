package domain

import (
	"fmt"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"gopkg.in/validator.v2"
)

type Track struct {
	Id             int    `json:"id" example:"4" db:"id" validate:"min=0"`
	AlbumId        int    `json:"albumId" example:"8" db:"album_id" validate:"min=0"`
	ArtistId       int    `json:"artistId" example:"8" db:"artist_id" validate:"min=0,nonnil"`
	Title          string `json:"title" example:"Rain" db:"title" validate:"max=256,nonnil"`
	Duration       int    `json:"duration" example:"180" db:"duration" validate:"min=0,nonnil"`
	CountLikes     int    `json:"countLikes" example:"54" db:"count_likes" validate:"min=0,nonnil"`
	CountListening int    `json:"countListening" example:"15632" db:"count_listening" validate:"min=0,nonnil"`
}

type TrackRepo interface {
	Insert(Track) error
	Update(Track) error
	Delete(int) error
	SelectByID(int) (*Track, error)
	GetAll() ([]Track, error)
	GetPopular() ([]Track, error)
	GetLastId() (id int, err error)
	//GetType() reflect.Type
	GetSize() (int, error)
	GetTracksFromAlbum(albumId int) ([]Track, error)
	GetPopularTracksFromArtist(artistId int) ([]Track, error)
}

type TrackDataTransfer struct {
	Id         int    `json:"id" example:"1"`
	Title      string `json:"title" example:"Mercury"`
	Artist     string `json:"artist" example:"Hexed"`
	Cover      string `json:"cover" example:"assets/track_1.png"`
	Src        string `json:"src" example:"assets/track_1.mp4"`
	Likes      int    `json:"likes" example:"5"`
	Listenings int    `json:"listenings" example:"500"`
	Duration   int    `json:"duration" example:"531"`
}

func (track *Track) GetId() int {
	return track.Id
}

func (track *Track) SetId(id int) error {
	track.Id = id
	return nil
}

func (track *Track) Check() error {
	return validator.Validate(track)
}

func (track *Track) GetCountListening() int {
	return track.CountListening
}

func (track *Track) CreatePath(fileFormat string) (string, error) {
	return constants.AssetsPrefix + constants.TrackPreName + fmt.Sprint(track.Id) + fileFormat, nil
}

func (track *Track) CreatePathById(fileFormat string, albumId int) (string, error) {
	return constants.AssetsPrefix + constants.AlbumPreName + fmt.Sprint(albumId) + fileFormat, nil
}
