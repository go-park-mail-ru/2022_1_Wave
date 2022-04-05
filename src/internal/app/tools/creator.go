package tools

import (
	"errors"
	"fmt"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/interfaces"
	"reflect"
)

type domainCreator struct{}

var Creator = domainCreator{}

func (creator domainCreator) DomainConstructor(domainType reflect.Type) (utilsInterfaces.Domain, error) {
	switch domainType {
	case domain.AlbumDomainType:
		return domain.Album{}, nil
	case domain.ArtistDomainType:
		return domain.Artist{}, nil
	case domain.TrackDomainType:
		return domain.Track{}, nil
	default:
		return nil, errors.New(constants.BadType)
	}
}

// ---------------------------------------------------------
func (creator domainCreator) InitDomainAlbum(id uint64) (domain.Album, error) {
	return domain.Album{
		Id:      id,
		CoverId: id,
	}, nil
}

// ---------------------------------------------------------
func (creator domainCreator) InitDomainArtist(id uint64) (domain.Artist, error) {
	return domain.Artist{
		Id: id,
	}, nil
}

// ---------------------------------------------------------
func (creator domainCreator) InitDomainTrack(id uint64) (domain.Track, error) {
	return domain.Track{
		Id:      id,
		CoverId: id,
	}, nil
}

func (creator domainCreator) InitDomain(id uint64, domainType reflect.Type) (utilsInterfaces.Domain, error) {
	var initDomain utilsInterfaces.Domain
	var err error

	switch domainType {
	case domain.AlbumDomainType:
		initDomain, err = creator.InitDomainAlbum(id)
		if err == nil {
			return initDomain, nil
		}

	case domain.ArtistDomainType:
		initDomain, err = creator.InitDomainArtist(id)
		if err == nil {
			return initDomain, nil
		}

	case domain.TrackDomainType:
		initDomain, err = creator.InitDomainTrack(id)
		if err == nil {
			return initDomain, nil
		}
	}

	return nil, err
}

// ---------------------------------------------------------
func (creator domainCreator) CreateDomainAlbumFromInterface(data interface{}) (utilsInterfaces.Domain, error) {
	temp := data.(map[string]interface{})

	id, err := Converter.ToUint64(temp[constants.FieldId])
	if err != nil {
		return nil, err
	}

	title, err := Converter.ToString(temp[constants.FieldTitle])
	if err != nil {
		return nil, err
	}

	artistId, err := Converter.ToUint64(temp[constants.FieldAuthorId])
	if err != nil {
		return nil, err
	}

	countLikes, err := Converter.ToUint64(temp[constants.FieldCountLikes])
	if err != nil {
		return nil, err
	}

	countListening, err := Converter.ToUint64(temp[constants.FieldCountListening])
	if err != nil {
		return nil, err
	}

	date, err := Converter.ToInt(temp[constants.FieldDate])
	if err != nil {
		return nil, err
	}

	//coverId, err := Converter.ToUint64(temp[constants.FieldCoverId])
	//if err != nil {
	//	return nil, err
	//}

	fmt.Println("coverId = ", id)

	return domain.Album{
		Id:             id,
		Title:          title,
		ArtistId:       artistId,
		CountLikes:     countLikes,
		CountListening: countListening,
		Date:           date,
		CoverId:        id,
	}, nil
}

// ---------------------------------------------------------
func (creator domainCreator) CreateDomainArtistFromInterface(data interface{}) (utilsInterfaces.Domain, error) {
	temp := data.(map[string]interface{})

	id, err := Converter.ToUint64(temp[constants.FieldId])
	if err != nil {
		return nil, err
	}

	name, err := Converter.ToString(temp[constants.FieldName])
	if err != nil {
		return nil, err
	}

	//photo, err := Converter.ToString(temp[constants.FieldPhoto])
	//if err != nil {
	//	return nil, err
	//}

	countListening, err := Converter.ToUint64(temp[constants.FieldCountListening])
	if err != nil {
		return nil, err
	}

	countFollowers, err := Converter.ToUint64(temp[constants.FieldCountFollowers])
	if err != nil {
		return nil, err
	}

	return domain.Artist{
		Id:             id,
		Name:           name,
		PhotoId:        id,
		CountFollowers: countFollowers,
		CountListening: countListening,
	}, nil
}

// ---------------------------------------------------------
func (creator domainCreator) CreateDomainTrackFromInterface(data interface{}) (utilsInterfaces.Domain, error) {
	temp := data.(map[string]interface{})

	id, err := Converter.ToUint64(temp[constants.FieldId])
	if err != nil {
		return nil, err
	}

	albumId, err := Converter.ToUint64(temp[constants.FieldAlbumId])
	if err != nil {
		return nil, err
	}

	title, err := Converter.ToString(temp[constants.FieldTitle])
	if err != nil {
		return nil, err
	}

	artistId, err := Converter.ToUint64(temp[constants.FieldAuthorId])
	if err != nil {
		return nil, err
	}

	countLikes, err := Converter.ToUint64(temp[constants.FieldCountLikes])
	if err != nil {
		return nil, err
	}

	countListening, err := Converter.ToUint64(temp[constants.FieldCountListening])
	if err != nil {
		return nil, err
	}

	duration, err := Converter.ToUint64(temp[constants.FieldDuration])
	if err != nil {
		return nil, err
	}

	mp4, err := Converter.ToString(temp[constants.FieldMp4])
	if err != nil {
		return nil, err
	}

	//coverId, err := Converter.ToUint64(temp[constants.FieldCoverId])
	//if err != nil {
	//	return nil, err
	//}

	return domain.Track{
		Id:             id,
		AlbumId:        albumId,
		ArtistId:       artistId,
		Title:          title,
		Duration:       duration,
		Mp4:            mp4,
		CoverId:        id,
		CountLikes:     countLikes,
		CountListening: countListening,
	}, nil
}

// ---------------------------------------------------------
func (creator domainCreator) CreateDomainFromInterface(domainType reflect.Type, data interface{}) (utilsInterfaces.Domain, error) {

	var resultDomain utilsInterfaces.Domain
	var err error

	if domainType == domain.AlbumDomainType {
		resultDomain, err = creator.CreateDomainAlbumFromInterface(data)
		if err == nil {
			return resultDomain, nil
		}
	}

	if domainType == domain.ArtistDomainType {
		resultDomain, err = creator.CreateDomainArtistFromInterface(data)
		if err == nil {
			return resultDomain, nil
		}
	}

	if domainType == domain.TrackDomainType {
		resultDomain, err = creator.CreateDomainTrackFromInterface(data)
		if err == nil {
			return resultDomain, nil
		}
	}

	return nil, err

}

// ---------------------------------------------------------
func (creator domainCreator) CreateAlbumDataTransferFromInterface(data interface{}) (utilsInterfaces.DataTransfer, error) {
	temp := data.(map[string]interface{})

	title, err := Converter.ToString(temp[constants.FieldTitle])
	if err != nil {
		return nil, err
	}

	artist, err := Converter.ToString(temp[constants.FieldArtist])
	if err != nil {
		return nil, err
	}

	cover, err := Converter.ToString(temp[constants.FieldCover])
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
func (creator domainCreator) CreateArtistDataTransferFromInterface(data interface{}) (utilsInterfaces.DataTransfer, error) {
	temp := data.(map[string]interface{})

	name, err := Converter.ToString(temp[constants.FieldName])
	if err != nil {
		return nil, err
	}

	cover, err := Converter.ToString(temp[constants.FieldCover])
	if err != nil {
		return nil, err
	}

	return domain.ArtistDataTransfer{
		Name:  name,
		Cover: cover,
	}, nil
}

// ---------------------------------------------------------
func (creator domainCreator) CreateTrackDataTransferFromInterface(data interface{}) (utilsInterfaces.DataTransfer, error) {
	temp := data.(map[string]interface{})

	title, err := Converter.ToString(temp[constants.FieldTitle])
	if err != nil {
		return nil, err
	}

	artist, err := Converter.ToString(temp[constants.FieldArtist])
	if err != nil {
		return nil, err
	}

	cover, err := Converter.ToString(temp[constants.FieldCover])
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
func (creator domainCreator) CreateDataTransferFromInterface(toType reflect.Type, data interface{}) (utilsInterfaces.DataTransfer, error) {
	var resultDataTransfer utilsInterfaces.DataTransfer
	var err error

	if toType == reflect.TypeOf(domain.AlbumDataTransfer{}) {
		resultDataTransfer, err = creator.CreateAlbumDataTransferFromInterface(data)
		if err == nil {
			return resultDataTransfer, nil
		}
	}

	if toType == reflect.TypeOf(domain.ArtistDataTransfer{}) {
		resultDataTransfer, err = creator.CreateArtistDataTransferFromInterface(data)
		if err == nil {
			return resultDataTransfer, nil
		}
	}

	if toType == reflect.TypeOf(domain.TrackDataTransfer{}) {
		resultDataTransfer, err = creator.CreateTrackDataTransferFromInterface(data)
		if err == nil {
			return resultDataTransfer, nil
		}
	}

	return nil, err

}
