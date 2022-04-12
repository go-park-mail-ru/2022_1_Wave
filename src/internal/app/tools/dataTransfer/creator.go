package dataTransferCreator

import (
	"errors"
	"fmt"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	albumUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/album/usecase"
	artistUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/artist/usecase"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	utilsInterfaces "github.com/go-park-mail-ru/2022_1_Wave/internal/app/interfaces"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/tools/utils"
	"reflect"
	"sync"
)

// ---------------------------------------------------------
func CreateAlbumDataTransferFromInterface(data interface{}) (utilsInterfaces.DataTransfer, error) {
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

	tracksArray := temp[constants.FieldTracks].([]interface{})
	tracks := make([]domain.TrackDataTransfer, len(tracksArray))
	for i, obj := range tracksArray {
		temp := obj.(map[string]interface{})
		track, err := CreateTrackDataTransferFromInterface(temp)
		if err != nil {
			return nil, err
		}
		tracks[i] = track.(domain.TrackDataTransfer)
	}

	return domain.AlbumDataTransfer{
		Title:  title,
		Artist: artist,
		Cover:  cover,
		Tracks: tracks,
	}, nil
}

// ---------------------------------------------------------
func CreateAlbumCoverDataTransferFromInterface(data interface{}) (utilsInterfaces.DataTransfer, error) {
	temp := data.(map[string]interface{})

	title, err := utils.ToString(temp[constants.FieldTitle])
	if err != nil {
		return nil, err
	}

	quote, err := utils.ToString(temp[constants.FieldQuote])
	if err != nil {
		return nil, err
	}

	isDark, err := utils.ToBool(temp[constants.FieldIsDark])
	if err != nil {
		return nil, err
	}

	return domain.AlbumCoverDataTransfer{
		Title:  title,
		Quote:  quote,
		IsDark: isDark,
	}, nil
}

// ---------------------------------------------------------
func CreateArtistDataTransferFromInterface(data interface{}) (utilsInterfaces.DataTransfer, error) {
	temp := data.(map[string]interface{})

	name, err := utils.ToString(temp[constants.FieldName])
	if err != nil {
		return nil, err
	}

	cover, err := utils.ToString(temp[constants.FieldCover])
	if err != nil {
		return nil, err
	}

	albumsArray := temp[constants.FieldAlbums].([]interface{})
	albums := make([]domain.AlbumDataTransfer, len(albumsArray))
	for i, obj := range albumsArray {
		temp := obj.(map[string]interface{})
		track, err := CreateAlbumDataTransferFromInterface(temp)
		if err != nil {
			return nil, err
		}
		albums[i] = track.(domain.AlbumDataTransfer)
	}

	return domain.ArtistDataTransfer{
		Name:   name,
		Cover:  cover,
		Albums: albums,
	}, nil
}

// ---------------------------------------------------------
func CreateTrackDataTransferFromInterface(data interface{}) (utilsInterfaces.DataTransfer, error) {
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

	return domain.TrackDataTransfer{
		Title:      title,
		Artist:     artist,
		Cover:      cover,
		Src:        src,
		Likes:      likes,
		Listenings: listenings,
		Duration:   duration,
	}, nil
}

// ---------------------------------------------------------
func CreateDataTransferFromInterface(dataTransferType reflect.Type, data interface{}) (utilsInterfaces.DataTransfer, error) {
	var resultDataTransfer utilsInterfaces.DataTransfer
	var err error

	switch dataTransferType {
	case domain.AlbumDataTransferType:
		resultDataTransfer, err = CreateAlbumDataTransferFromInterface(data)
	case domain.AlbumCoverDataTransferType:
		resultDataTransfer, err = CreateAlbumCoverDataTransferFromInterface(data)
	case domain.ArtistDataTransferType:
		resultDataTransfer, err = CreateArtistDataTransferFromInterface(data)
	case domain.TrackDataTransferType:
		resultDataTransfer, err = CreateTrackDataTransferFromInterface(data)
	default:
		resultDataTransfer = nil
		err = errors.New(constants.BadType)
	}

	return resultDataTransfer, err
}

func CreateDataTransfer(dom utilsInterfaces.Domain, mutex *sync.RWMutex) (utilsInterfaces.DataTransfer, error) {
	switch reflect.TypeOf(dom) {

	case domain.AlbumDomainType:
		artistId := dom.(domain.Album).ArtistId
		artistInCurrentAlbum, err := artistUseCase.UseCase.GetById(artistId, domain.ArtistMutex)
		if err != nil {
			return nil, err
		}

		tracks, err := albumUseCase.UseCase.GetTracksFromAlbum(dom.GetId(), domain.AlbumMutex)
		if err != nil {
			return nil, err
		}

		result := tracks.([]domain.Track)
		dataTransfers := make([]domain.TrackDataTransfer, len(result))

		for i, obj := range result {
			data, err := CreateDataTransfer(obj, domain.TrackMutex)
			if err != nil {
				return nil, err
			}
			dataTransfers[i] = data.(domain.TrackDataTransfer)
		}

		return dom.CastDomainToDataTransferObject(artistInCurrentAlbum, dataTransfers)

	case domain.AlbumCoverDomainType:
		return dom.CastDomainToDataTransferObject(nil)

	case domain.ArtistDomainType:
		albums, err := artistUseCase.UseCase.GetAlbumsFromArtist(dom.GetId(), domain.ArtistMutex)

		fmt.Println(albums)

		if err != nil {
			return nil, err
		}

		result := albums.([]domain.Album)
		dataTransfers := make([]domain.AlbumDataTransfer, len(result))

		for i, obj := range result {
			fmt.Println(reflect.TypeOf(obj))
			data, err := CreateDataTransfer(obj, domain.AlbumMutex)
			if err != nil {
				return nil, err
			}
			dataTransfers[i] = data.(domain.AlbumDataTransfer)
		}

		return dom.CastDomainToDataTransferObject(nil, dataTransfers)

	case domain.TrackDomainType:
		artistId := dom.(domain.Track).ArtistId
		artistInCurrentTrack, err := artistUseCase.UseCase.GetById(artistId, domain.ArtistMutex)
		if err != nil {
			return nil, err
		}
		return dom.CastDomainToDataTransferObject(artistInCurrentTrack)

	default:
		return nil, errors.New(constants.BadType)
	}
}
