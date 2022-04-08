package domain

import (
	"errors"
	"fmt"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/interfaces"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/tools/utils"
)

//Mp4            string      `json:"mp4" example:"assets/track_1.mp4" db:"mp4"`
//CoverId        uint64      `json:"coverId" example:"4" db:"cover_id"`

type Track struct {
	Id uint64 `json:"id" example:"4" db:"id"`
	// AlbumId is uint64 but for null holder this is type interface
	AlbumId  interface{} `json:"albumId" db:"album_id"`
	ArtistId uint64      `json:"artistId" example:"8" db:"artist_id"`
	Title    string      `json:"title" example:"Rain" db:"title"`
	Duration uint64      `json:"duration" example:"180" db:"duration"`
	//Mp4            string      `db:"mp4"`
	//CoverId        uint64      `db:"cover_id"`
	CountLikes     uint64 `json:"countLikes" example:"54" db:"count_likes"`
	CountListening uint64 `json:"countListening" example:"15632" db:"count_listening"`
}

func (track Track) GetId() uint64 {
	return track.Id
}

func (track Track) SetId(id uint64) (utilsInterfaces.Domain, error) {
	track.Id = id
	//track.CoverId = track.Id
	return track, nil
}

func (track Track) Check() error {
	if track.Id < 0 {
		return errors.New(constants.ErrorTrackIdIsNegative)
	}

	if len(track.Title) > constants.TrackTitleLen {
		return errors.New(constants.ErrorTrackMaxTitleLen)
	}

	//if len(track.Mp4) > constants.TrackMp4LinkLen {
	//	return errors.New(constants.ErrorTrackMp4MaxLinkLen)
	//}

	if track.CountLikes < 0 {
		return errors.New(constants.ErrorTrackCountLikesIsNegative)
	}

	if track.CountListening < 0 {
		return errors.New(constants.ErrorTrackCountListeningIsNegative)
	}

	if track.ArtistId < 0 {
		return errors.New(constants.ErrorTrackArtistIdIsNegative)
	}

	//if track.AlbumId.Valid && track.AlbumId.Int64 < 0 {
	//	return errors.New(constants.ErrorTrackAlbumIdIsNegative)
	//}

	//if track.CoverId < 0 {
	//	return errors.New(constants.ErrorTrackCoverIdIsNegative)
	//}

	return nil
}

func (track Track) GetCountListening() uint64 {
	return track.CountListening
}

func (track Track) CreatePath(fileFormat string) (string, error) {
	return constants.AssetsPrefix + constants.TrackPreName + fmt.Sprint(track.Id) + fileFormat, nil
}

func (track Track) CastDomainToDataTransferObject(artist utilsInterfaces.Domain) (utilsInterfaces.DataTransfer, error) {
	ptr := artist.(*Artist)

	pathToCover, err := track.CreatePath(constants.PngFormat)
	if err != nil {
		return nil, nil
	}

	pathToSrc, err := track.CreatePath(constants.Mp4Format)
	if err != nil {
		return nil, nil
	}

	return TrackDataTransfer{
		Title:      track.Title,
		Artist:     (*ptr).Name,
		Cover:      pathToCover,
		Src:        pathToSrc,
		Likes:      int(track.CountLikes),
		Listenings: int(track.CountListening),
		Duration:   int(track.Duration),
	}, nil
}

type TrackDataTransfer struct {
	Title      string `json:"title" example:"Mercury"`
	Artist     string `json:"artist" example:"Hexed"`
	Cover      string `json:"cover" example:"assets/track_1.png"`
	Src        string `json:"src" example:"assets/track_1.mp4"`
	Likes      int    `json:"likes" example:"5"`
	Listenings int    `json:"listenings" example:"500"`
	Duration   int    `json:"duration" example:"531"`
}

func (track TrackDataTransfer) CreateDataTransferFromInterface(data interface{}) (utilsInterfaces.DataTransfer, error) {
	temp := data.(map[string]interface{})

	title, err := utils.ToString(temp[constants.FieldTitle])
	if err != nil {
		return nil, err
	}

	artist, err := utils.ToString(temp[constants.FieldArtist])
	if err != nil {
		return nil, err
	}

	cover, err := utils.ToString(temp[constants.FieldCover])
	if err != nil {
		return nil, err
	}

	src, err := utils.ToString(temp[constants.FieldSrc])
	if err != nil {
		return nil, err
	}

	likes, err := utils.ToInt(temp[constants.FieldLikes])
	if err != nil {
		return nil, err
	}

	listenings, err := utils.ToInt(temp[constants.FieldListenings])
	if err != nil {
		return nil, err
	}

	duration, err := utils.ToInt(temp[constants.FieldDuration])
	if err != nil {
		return nil, err
	}
	return TrackDataTransfer{
		Title:      title,
		Artist:     artist,
		Cover:      cover,
		Src:        src,
		Likes:      likes,
		Listenings: listenings,
		Duration:   duration,
	}, nil
}
