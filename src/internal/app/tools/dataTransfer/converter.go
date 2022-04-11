package dataTransferCreator

import (
	"errors"
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/interfaces"
	"reflect"
)

func ToDataTransfers(toType reflect.Type, objects []interface{}) (*[]utilsInterfaces.DataTransfer, error) {
	it := make([]utilsInterfaces.DataTransfer, len(objects))
	var err error
	for idx, object := range objects {
		it[idx], err = CreateDataTransferFromInterface(toType, object)
		if err != nil {
			return nil, err
		}
	}
	return &it, nil
}

func GetDataTransferTypeByDomainType(domainType reflect.Type) (reflect.Type, error) {
	switch domainType {
	case domain.AlbumDomainType:
		return domain.AlbumDataTransferType, nil
	case domain.AlbumCoverDomainType:
		return domain.AlbumCoverDataTransferType, nil
	case domain.ArtistDomainType:
		return domain.ArtistDataTransferType, nil
	case domain.TrackDomainType:
		return domain.TrackDataTransferType, nil
	default:
		return nil, errors.New(internal.BadType)
	}
}
