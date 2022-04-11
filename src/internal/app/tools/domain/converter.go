package domainCreator

import (
	"errors"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/interfaces"
)

//func ToDomains(objects []interface{}) (*[]utilsInterfaces.Domain, error) {
//	it := make([]utilsInterfaces.Domain, len(objects))
//	for idx, object := range objects {
//		it[idx] = object.(utilsInterfaces.Domain)
//	}
//	return &it, nil
//}

func ToDomainsArrayPtr(holder *interface{}, repoName string) error {
	switch repoName {
	case constants.Album:
		*holder = &[]domain.Album{}
	case constants.AlbumCover:
		*holder = &[]domain.AlbumCover{}
	case constants.Artist:
		*holder = &[]domain.Artist{}
	case constants.Track:
		*holder = &[]domain.Track{}
	default:
		return errors.New(constants.BadType)
	}
	return nil
}

func GetValues(holder interface{}, repoName string) ([]utilsInterfaces.Domain, error) {
	var result []utilsInterfaces.Domain
	switch repoName {
	case constants.Album:
		values := holder.(*[]domain.Album)
		for _, obj := range *values {
			result = append(result, obj)
		}
	case constants.AlbumCover:
		values := holder.(*[]domain.AlbumCover)
		for _, obj := range *values {
			result = append(result, obj)
		}
	case constants.Artist:
		values := holder.(*[]domain.Artist)
		for _, obj := range *values {
			result = append(result, obj)
		}
	case constants.Track:
		values := holder.(*[]domain.Track)
		for _, obj := range *values {
			result = append(result, obj)
		}
	default:
		return nil, errors.New(constants.BadType)
	}
	return result, nil
}

//func ToDomain(holder interface{}, repoName string) error {
//	switch repoName {
//	case constants.Album:
//		return domain.Album{}, nil
//	case constants.AlbumCover:
//		return domain.AlbumCover{}, nil
//	case constants.Artist:
//		return domain.Artist{}, nil
//	case constants.Track:
//		return domain.Track{}, nil
//	default:
//		return nil, errors.New(constants.BadType)
//	}
//}

//func ToDomainPtrByDomainType(domainType reflect.Type) (utilsInterfaces.Domain, error) {
//	switch domainType {
//	case domain.AlbumDomainType:
//		return &domain.Album{}, nil
//	case domain.AlbumCoverDomainType:
//		return &domain.AlbumCover{}, nil
//	case domain.ArtistDomainType:
//		return &domain.Artist{}, nil
//	case domain.TrackDomainType:
//		return &domain.Track{}, nil
//	default:
//		return nil, errors.New(constants.BadType)
//	}
//}

//func ToDomainPtrByTableName(repoName string) (interface{}, error) {
//	switch repoName {
//	case constants.Album:
//		return &domain.Album{}, nil
//	case constants.AlbumCover:
//		return &domain.AlbumCover{}, nil
//	case constants.Artist:
//		return &domain.Artist{}, nil
//	case constants.Track:
//		return &domain.Track{}, nil
//	default:
//		return nil, errors.New(constants.BadType)
//	}
//}
