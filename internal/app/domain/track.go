package domain

import (
	"fmt"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/interfaces"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/tools/utils"
	"gopkg.in/validator.v2"
)

//Mp4            string      `json:"mp4" example:"assets/track_1.mp4" db:"mp4"`
//CoverId        uint64      `json:"coverId" example:"4" db:"cover_id"`

type Track struct {
	Id             uint64 `json:"id" example:"4" db:"id" validate:"min=0"`
	AlbumId        uint64 `json:"albumId" example:"8" db:"album_id" validate:"min=0"`
	ArtistId       uint64 `json:"artistId" example:"8" db:"artist_id" validate:"min=0,nonnil"`
	Title          string `json:"title" example:"Rain" db:"title" validate:"max=256,nonnil"`
	Duration       uint64 `json:"duration" example:"180" db:"duration" validate:"min=0,nonnil"`
	CountLikes     uint64 `json:"countLikes" example:"54" db:"count_likes" validate:"min=0,nonnil"`
	CountListening uint64 `json:"countListening" example:"15632" db:"count_listening" validate:"min=0,nonnil"`
}

func (track Track) GetId() uint64 {
	return track.Id
}

func (track Track) SetId(id uint64) (utilsInterfaces.Domain, error) {
	track.Id = id
	return track, nil
}

func (track Track) Check() error {
	return validator.Validate(track)
	//if track.Id < 0 {
	//	return errors.New(constants.ErrorTrackIdIsNegative)
	//}
	//
	//if len(track.Title) > constants.TrackTitleLen {
	//	return errors.New(constants.ErrorTrackMaxTitleLen)
	//}
	//
	////if len(track.Mp4) > constants.TrackMp4LinkLen {
	////	return errors.New(constants.ErrorTrackMp4MaxLinkLen)
	////}
	//
	//if track.CountLikes < 0 {
	//	return errors.New(constants.ErrorTrackCountLikesIsNegative)
	//}
	//
	//if track.CountListening < 0 {
	//	return errors.New(constants.ErrorTrackCountListeningIsNegative)
	//}
	//
	//if track.ArtistId < 0 {
	//	return errors.New(constants.ErrorTrackArtistIdIsNegative)
	//}

	//if track.AlbumId.Valid && track.AlbumId.Int64 < 0 {
	//	return errors.New(constants.ErrorTrackAlbumIdIsNegative)
	//}

	//if track.CoverId < 0 {
	//	return errors.New(constants.ErrorTrackCoverIdIsNegative)
	//}

	//return nil
}

func (track Track) GetCountListening() uint64 {
	return track.CountListening
}

func (track Track) CreatePath(fileFormat string) (string, error) {
	return constants.AssetsPrefix + constants.TrackPreName + fmt.Sprint(track.Id) + fileFormat, nil
}

func (track Track) CreatePathById(fileFormat string, albumId uint64) (string, error) {
	return constants.AssetsPrefix + constants.AlbumPreName + fmt.Sprint(albumId) + fileFormat, nil
}

func (track Track) CastDomainToDataTransferObject(artist utilsInterfaces.Domain, args ...interface{}) (utilsInterfaces.DataTransfer, error) {

	pathToCover, err := track.CreatePathById(constants.PngFormat, track.AlbumId)
	if err != nil {
		return nil, nil
	}

	pathToSrc, err := track.CreatePath(constants.Mp3Format)
	if err != nil {
		return nil, nil
	}

	return TrackDataTransfer{
		Id:         track.Id,
		Title:      track.Title,
		Artist:     artist.(Artist).Name,
		Cover:      pathToCover,
		Src:        pathToSrc,
		Likes:      int(track.CountLikes),
		Listenings: int(track.CountListening),
		Duration:   int(track.Duration),
	}, nil
}

type TrackDataTransfer struct {
	Id         uint64 `json:"id" example:"1"`
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

	id, err := utils.ToUint64(temp[constants.FieldId])
	if err != nil {
		return nil, err
	}

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
		Id:         id,
		Title:      title,
		Artist:     artist,
		Cover:      cover,
		Src:        src,
		Likes:      likes,
		Listenings: listenings,
		Duration:   duration,
	}, nil
}
