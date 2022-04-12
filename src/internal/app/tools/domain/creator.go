package domainCreator

import (
	"errors"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/interfaces"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/tools/utils"
	"math/rand"
	"reflect"
)

//type domainCreator struct{}
//
//var Creator = domainCreator{}

//func DomainConstructor(domainType reflect.Type) (utilsInterfaces.Domain, error) {
//	switch domainType {
//	case domain.AlbumDomainType:
//		return domain.Album{}, nil
//	case domain.AlbumCoverDomainType:
//		return domain.AlbumCover{}, nil
//	case domain.ArtistDomainType:
//		return domain.Artist{}, nil
//	case domain.TrackDomainType:
//		return domain.Track{}, nil
//	default:
//		return nil, errors.New(constants.BadType)
//	}
//}

// ---------------------------------------------------------
func InitDomainAlbum(id uint64) (domain.Album, error) {
	return domain.Album{
		Id: id,
		//CoverId: id,
	}, nil
}

// ---------------------------------------------------------
func InitDomainAlbumCover(id uint64) (domain.AlbumCover, error) {
	return domain.AlbumCover{
		Id: id,
	}, nil
}

// ---------------------------------------------------------
func InitDomainArtist(id uint64) (domain.Artist, error) {
	return domain.Artist{
		Id: id,
	}, nil
}

// ---------------------------------------------------------
func InitDomainTrack(id uint64) (domain.Track, error) {
	return domain.Track{
		Id: id,
		//CoverId: id,
	}, nil
}

func InitDomain(id uint64, domainType reflect.Type) (utilsInterfaces.Domain, error) {
	var initDomain utilsInterfaces.Domain
	var err error

	switch domainType {
	case domain.AlbumDomainType:
		initDomain, err = InitDomainAlbum(id)
		if err == nil {
			return initDomain, nil
		}

	case domain.AlbumCoverDomainType:
		initDomain, err = InitDomainAlbumCover(id)
		if err == nil {
			return initDomain, nil
		}

	case domain.ArtistDomainType:
		initDomain, err = InitDomainArtist(id)
		if err == nil {
			return initDomain, nil
		}

	case domain.TrackDomainType:
		initDomain, err = InitDomainTrack(id)
		if err == nil {
			return initDomain, nil
		}
	}

	return nil, err
}

// ---------------------------------------------------------
func CreateDomainAlbumFromInterface(data interface{}) (utilsInterfaces.Domain, error) {
	temp := data.(map[string]interface{})

	id, err := utils.ToUint64(temp[constants.FieldId])
	if err != nil {
		return nil, err
	}

	title, err := utils.ToString(temp[constants.FieldTitle])
	if err != nil {
		return nil, err
	}

	artistId, err := utils.ToUint64(temp[constants.FieldArtistId])
	if err != nil {
		return nil, err
	}

	countLikes, err := utils.ToUint64(temp[constants.FieldCountLikes])
	if err != nil {
		return nil, err
	}

	countListening, err := utils.ToUint64(temp[constants.FieldCountListening])
	if err != nil {
		return nil, err
	}

	date, err := utils.ToInt64(temp[constants.FieldDate])
	if err != nil {
		return nil, err
	}

	//coverId, err := utils.ToUint64(temp[constants.FieldCoverId])
	//if err != nil {
	//	return nil, err
	//}

	return domain.Album{
		Id:             id,
		Title:          title,
		ArtistId:       artistId,
		CountLikes:     countLikes,
		CountListening: countListening,
		Date:           date,
		//CoverId:        id,
	}, nil
}

// ---------------------------------------------------------
func CreateDomainAlbumCoverFromInterface(data interface{}) (utilsInterfaces.Domain, error) {
	temp := data.(map[string]interface{})

	id, err := utils.ToUint64(temp[constants.FieldId])
	if err != nil {
		return nil, err
	}

	//title, err := utils.ToString(temp[constants.FieldTitle])
	//if err != nil {
	//	return nil, err
	//}

	quote, err := utils.ToString(temp[constants.FieldQuote])
	if err != nil {
		return nil, err
	}

	isDark, err := utils.ToBool(temp[constants.FieldIsDark])
	if err != nil {
		return nil, err
	}

	return domain.AlbumCover{
		Id: id,
		//Title:  title,
		Quote:  quote,
		IsDark: isDark,
	}, nil
}

// ---------------------------------------------------------
func CreateDomainArtistFromInterface(data interface{}) (utilsInterfaces.Domain, error) {
	temp := data.(map[string]interface{})

	id, err := utils.ToUint64(temp[constants.FieldId])
	if err != nil {
		return nil, err
	}

	name, err := utils.ToString(temp[constants.FieldName])
	if err != nil {
		return nil, err
	}

	//photo, err := utils.ToString(temp[constants.FieldPhoto])
	//if err != nil {
	//	return nil, err
	//}

	countListening, err := utils.ToUint64(temp[constants.FieldCountListening])
	if err != nil {
		return nil, err
	}

	countFollowers, err := utils.ToUint64(temp[constants.FieldCountFollowers])
	if err != nil {
		return nil, err
	}

	return domain.Artist{
		Id:   id,
		Name: name,
		//PhotoId:        id,
		CountFollowers: countFollowers,
		CountListening: countListening,
	}, nil
}

// ---------------------------------------------------------
func CreateDomainTrackFromInterface(data interface{}) (utilsInterfaces.Domain, error) {
	temp := data.(map[string]interface{})

	id, err := utils.ToUint64(temp[constants.FieldId])
	if err != nil {
		return nil, err
	}

	//isNullAlbum, err := utils.CheckNullInt64(temp[constants.FieldAlbumId])
	//if err != nil {
	//	return nil, err
	//}

	//album := sql.NullInt64{}
	//if isNullAlbum {
	//	album.Valid = false
	//} else {
	//	albumId, err := utils.ToUint64(temp[constants.FieldAlbumId])
	//	if err != nil {
	//		return nil, err
	//	}
	//	album.Int64 = int64(albumId)
	//}

	albumIdField := temp[constants.FieldAlbumId]
	var albumId interface{}
	if albumIdField != nil {
		albumId, err = utils.ToUint64(albumIdField)
		if err != nil {
			return nil, err
		}
	}

	title, err := utils.ToString(temp[constants.FieldTitle])
	if err != nil {
		return nil, err
	}

	artistId, err := utils.ToUint64(temp[constants.FieldArtistId])
	if err != nil {
		return nil, err
	}

	countLikes, err := utils.ToUint64(temp[constants.FieldCountLikes])
	if err != nil {
		return nil, err
	}

	countListening, err := utils.ToUint64(temp[constants.FieldCountListening])
	if err != nil {
		return nil, err
	}

	duration, err := utils.ToUint64(temp[constants.FieldDuration])
	if err != nil {
		return nil, err
	}

	//mp4, err := utils.ToString(temp[constants.FieldMp4])
	//if err != nil {
	//	return nil, err
	//}

	//coverId, err := utils.ToUint64(temp[constants.FieldCoverId])
	//if err != nil {
	//	return nil, err
	//}

	return domain.Track{
		Id:       id,
		AlbumId:  albumId,
		ArtistId: artistId,
		Title:    title,
		Duration: duration,
		//Mp4:            mp4,
		//CoverId:        id,
		CountLikes:     countLikes,
		CountListening: countListening,
	}, nil
}

// ---------------------------------------------------------
func CreateDomainFromInterface(domainType reflect.Type, data interface{}) (utilsInterfaces.Domain, error) {
	var resultDomain utilsInterfaces.Domain
	var err error

	switch domainType {
	case domain.AlbumDomainType:
		resultDomain, err = CreateDomainAlbumFromInterface(data)
	case domain.AlbumCoverDomainType:
		resultDomain, err = CreateDomainAlbumCoverFromInterface(data)
	case domain.ArtistDomainType:
		resultDomain, err = CreateDomainArtistFromInterface(data)
	case domain.TrackDomainType:
		resultDomain, err = CreateDomainTrackFromInterface(data)
	default:
		resultDomain = nil
		err = errors.New(constants.BadType)
	}

	return resultDomain, err
}

// -----------------------------------------
func ArtistConstructorRandom(id uint64, maxNameLen int, maxFollowers int64, maxListening int64) domain.Artist {
	//rand.Seed(time.Now().Unix())
	return domain.Artist{
		Id:             id,
		Name:           utils.RandomWord(maxNameLen),
		CountFollowers: uint64(rand.Int63n(maxFollowers + 1)),
		CountListening: uint64(rand.Int63n(maxListening + 1)),
	}
}

func AlbumConstructorRandom(id uint64, authorsQuantity int64, maxAlbumTitleLen int, maxLikes int64, maxListening int64) domain.Album {
	//rand.Seed(time.Now().Unix())
	return domain.Album{
		Id:             id,
		Title:          utils.RandomWord(maxAlbumTitleLen),
		ArtistId:       1 + uint64(rand.Int63n(authorsQuantity-1)),
		CountLikes:     uint64(rand.Int63n(maxLikes + 1)),
		CountListening: uint64(rand.Int63n(maxListening + 1)),
		Date:           0,
	}
}

func AlbumCoverConstructorRandom(id uint64) domain.AlbumCover {
	//rand.Seed(time.Now().Unix())
	return domain.AlbumCover{
		Id: id,
		//Title:  utils.RandomWord(maxAlbumTitleLen),
		Quote:  utils.RandomWord(100),
		IsDark: true,
	}
}

func TrackConstructorRandom(id uint64, albums []utilsInterfaces.Domain, maxTrackTitleLen int, maxDuration int64, maxLikes int64, maxListening int64) domain.Track {
	//rand.Seed(time.Now().Unix())
	album := albums[1+rand.Intn(len(albums)-1)]
	albumId := album.GetId()
	artistId := album.(domain.Album).ArtistId

	return domain.Track{
		Id:             id,
		AlbumId:        albumId,
		ArtistId:       artistId,
		Title:          utils.RandomWord(maxTrackTitleLen),
		Duration:       uint64(rand.Int63n(maxDuration + 1)),
		CountLikes:     uint64(rand.Int63n(maxLikes + 1)),
		CountListening: uint64(rand.Int63n(maxListening + 1)),
	}
}
