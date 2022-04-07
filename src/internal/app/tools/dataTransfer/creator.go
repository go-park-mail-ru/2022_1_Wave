package dataTransferCreator

import (
	"errors"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	artistUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/artist/usecase"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	utilsInterfaces "github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/interfaces"
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

	return domain.AlbumDataTransfer{
		Title:  title,
		Artist: artist,
		Cover:  cover,
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

	return domain.ArtistDataTransfer{
		Name:  name,
		Cover: cover,
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

	return domain.TrackDataTransfer{
		Title:  title,
		Artist: artist,
		Cover:  cover,
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
	case domain.TrackDomainType:
		resultDataTransfer, err = CreateTrackDataTransferFromInterface(data)
	default:
		resultDataTransfer = nil
		err = errors.New(constants.BadType)
	}

	return resultDataTransfer, err
}

func CreateDataTransfer(domainType reflect.Type, dom utilsInterfaces.Domain, mutex *sync.RWMutex) (utilsInterfaces.DataTransfer, error) {
	switch domainType {
	case domain.AlbumDomainType:
		artistId := dom.(*domain.Album).ArtistId
		artistInCurrentAlbum, err := artistUseCase.UseCase.GetById(artistId, mutex)
		if err != nil {
			return nil, err
		}
		return dom.CastDomainToDataTransferObject(*artistInCurrentAlbum)
	case domain.AlbumCoverDomainType:
		return dom.CastDomainToDataTransferObject(nil)
	case domain.ArtistDomainType:
		return dom.CastDomainToDataTransferObject(nil)
	case domain.TrackDomainType:
		artistId := dom.(*domain.Track).ArtistId
		artistInCurrentTrack, err := artistUseCase.UseCase.GetById(artistId, mutex)
		if err != nil {
			return nil, err
		}
		return dom.CastDomainToDataTransferObject(*artistInCurrentTrack)
	default:
		return nil, errors.New(constants.BadType)
	}
}
