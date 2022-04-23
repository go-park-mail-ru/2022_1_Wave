package domain

import (
	"fmt"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"gopkg.in/validator.v2"
)

type Track struct {
	Id             int64  `json:"id" example:"4" db:"id" validate:"min=0"`
	AlbumId        int64  `json:"albumId" example:"8" db:"album_id" validate:"min=0"`
	ArtistId       int64  `json:"artistId" example:"8" db:"artist_id" validate:"min=0,nonnil"`
	Title          string `json:"title" example:"Rain" db:"title" validate:"max=256,nonnil"`
	Duration       int64  `json:"duration" example:"180" db:"duration" validate:"min=0,nonnil"`
	CountLikes     int64  `json:"countLikes" example:"54" db:"count_likes" validate:"min=0,nonnil"`
	CountListening int64  `json:"countListening" example:"15632" db:"count_listening" validate:"min=0,nonnil"`
}

type TrackRepo interface {
	Insert(Track) error
	Update(Track) error
	Delete(int64) error
	SelectByID(int64) (*Track, error)
	GetAll() ([]Track, error)
	GetPopular() ([]Track, error)
	GetLastId() (id int64, err error)
	//GetType() reflect.Type
	GetSize() (int64, error)
	GetTracksFromAlbum(albumId int64) ([]Track, error)
	GetPopularTracksFromArtist(artistId int64) ([]Track, error)
}

type TrackDataTransfer struct {
	Id         int64  `json:"id" example:"1"`
	Title      string `json:"title" example:"Mercury"`
	Artist     string `json:"artist" example:"Hexed"`
	Cover      string `json:"cover" example:"assets/track_1.png"`
	Src        string `json:"src" example:"assets/track_1.mp4"`
	Likes      int64  `json:"likes" example:"5"`
	Listenings int64  `json:"listenings" example:"500"`
	Duration   int64  `json:"duration" example:"531"`
}

func (track *Track) CastToDtoWithoutArtistName() (*TrackDataTransfer, error) {
	cover, err := track.CreatePathById(constants.PngFormat, track.AlbumId)
	if err != nil {
		return nil, err
	}

	src, err := track.CreatePath(constants.Mp3Format)
	if err != nil {
		return nil, err
	}

	return &TrackDataTransfer{
		Id:         track.Id,
		Title:      track.Title,
		Artist:     "",
		Cover:      cover,
		Src:        src,
		Likes:      track.CountLikes,
		Listenings: track.CountListening,
		Duration:   track.Duration,
	}, nil
}

func CastTracksByArtistToDto(tracks []Track, artist Artist) ([]TrackDataTransfer, error) {
	tracksDto := make([]TrackDataTransfer, len(tracks))
	for idx, track := range tracks {
		trackDto, err := track.CastToDtoWithoutArtistName()
		if err != nil {
			return nil, err
		}
		trackDto.Artist = artist.Name
		tracksDto[idx] = *trackDto
	}
	return tracksDto, nil
}

func (track *Track) GetId() int64 {
	return track.Id
}

func (track *Track) SetId(id int64) error {
	track.Id = id
	return nil
}

func (track *Track) Check() error {
	return validator.Validate(track)
}

func (track *Track) GetCountListening() int64 {
	return track.CountListening
}

func (track *Track) CreatePath(fileFormat string) (string, error) {
	return constants.AssetsPrefix + constants.TrackPreName + fmt.Sprint(track.Id) + fileFormat, nil
}

func (track *Track) CreatePathById(fileFormat string, albumId int64) (string, error) {
	return constants.AssetsPrefix + constants.AlbumPreName + fmt.Sprint(albumId) + fileFormat, nil
}
