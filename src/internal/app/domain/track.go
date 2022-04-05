package domain

import (
	"errors"
	"fmt"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/interfaces"
)

type Track struct {
	Id             uint64 `json:"id" example:"4"`
	AlbumId        uint64 `json:"albumId" example:"6"`
	ArtistId       uint64 `json:"authorId" example:"8"`
	Title          string `json:"title" example:"Rain"`
	Duration       uint64 `json:"duration" example:"180"`
	Mp4            string `json:"mp4" example:"assets/track_1.mp4"`
	CoverId        uint64 `json:"coverId" example:"4"`
	CountLikes     uint64 `json:"countLikes" example:"54"`
	CountListening uint64 `json:"countListening" example:"15632"`
}

func (track Track) GetId() uint64 {
	return track.Id
}

func (track Track) SetId(id uint64) (utilsInterfaces.Domain, error) {
	track.Id = id
	track.CoverId = track.Id
	return track, nil
}

func (track Track) Check() error {
	if track.Id < 0 {
		return errors.New(constants.ErrorTrackIdIsNegative)
	}

	if len(track.Title) > constants.TrackTitleLen {
		return errors.New(constants.ErrorTrackMaxTitleLen)
	}

	if len(track.Mp4) > constants.TrackMp4LinkLen {
		return errors.New(constants.ErrorTrackMp4MaxLinkLen)
	}

	if track.CountLikes < 0 {
		return errors.New(constants.ErrorTrackCountLikesIsNegative)
	}

	if track.CountListening < 0 {
		return errors.New(constants.ErrorTrackCountListeningIsNegative)
	}

	if track.ArtistId < 0 {
		return errors.New(constants.ErrorTrackArtistIdIsNegative)
	}

	if track.CoverId < 0 {
		return errors.New(constants.ErrorTrackCoverIdIsNegative)
	}

	return nil
}

func (track Track) GetCountListening() uint64 {
	return track.CountListening
}

func (track Track) CastDomainToDataTransferObject(artist utilsInterfaces.Domain) (utilsInterfaces.DataTransfer, error) {
	return TrackDataTransfer{
		Title:  track.Title,
		Artist: artist.(Artist).Name,
		Cover:  constants.AssetsPrefix + constants.TrackPreName + fmt.Sprint(track.CoverId) + constants.PngFormat,
	}, nil
}

type TrackDataTransfer struct {
	Title  string `json:"title" example:"Mercury"`
	Artist string `json:"artist" example:"Hexed"`
	Cover  string `json:"cover" example:"assets/album_1.png"`
}

func (track TrackDataTransfer) CreateDataTransferFromInterface(data interface{}) (utilsInterfaces.DataTransfer, error) {
	temp := data.(map[string]interface{})
	return TrackDataTransfer{
		Title:  temp["title"].(string),
		Artist: temp["artist"].(string),
		Cover:  temp["cover"].(string),
	}, nil
}
