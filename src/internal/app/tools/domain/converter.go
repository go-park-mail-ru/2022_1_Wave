package domainCreator

import (
	"errors"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/interfaces"
)

func ToDomains(objects []interface{}) (*[]utilsInterfaces.Domain, error) {
	it := make([]utilsInterfaces.Domain, len(objects))
	for idx, object := range objects {
		it[idx] = object.(utilsInterfaces.Domain)
	}
	return &it, nil
}

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

func ToDomainPtr(repoName string) (utilsInterfaces.Domain, error) {
	var holder utilsInterfaces.Domain
	switch repoName {
	case constants.Album:
		holder = &domain.Album{}
	case constants.AlbumCover:
		holder = &domain.AlbumCover{}
	case constants.Artist:
		holder = &domain.Artist{}
	case constants.Track:
		holder = &domain.Track{}
	default:
		return nil, errors.New(constants.BadType)
	}
	return holder, nil
}
