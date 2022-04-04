package tools

import (
	"errors"
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/interfaces"
	"reflect"
)

type domainConverter struct{}

var Converter = domainConverter{}

func (c domainConverter) ToDomains(objects []interface{}) (*[]utilsInterfaces.Domain, error) {
	it := make([]utilsInterfaces.Domain, len(objects))
	for idx, object := range objects {
		it[idx] = object.(utilsInterfaces.Domain)
	}
	return &it, nil
}

func (c domainConverter) ToDataTransfers(toType reflect.Type, objects []interface{}) (*[]utilsInterfaces.DataTransfer, error) {
	it := make([]utilsInterfaces.DataTransfer, len(objects))
	var err error
	for idx, object := range objects {
		it[idx], err = Creator.CreateDataTransferFromInterface(toType, object)
		if err != nil {
			return nil, err
		}
	}
	return &it, nil
}

func (c domainConverter) ToString(value interface{}) (string, error) {
	str, ok := value.(string)

	if !ok {
		return "", errors.New("error to casting to string")
	}

	return str, nil
}

func (c domainConverter) ToUint64(value interface{}) (uint64, error) {
	float, ok := value.(float64)

	if !ok {
		return 0, errors.New("error to casting to uint64")
	}

	unsigned := uint64(float)

	return unsigned, nil
}

func (c domainConverter) ToInt(value interface{}) (int, error) {
	float, ok := value.(float64)

	if !ok {
		return 0, errors.New("error to casting to uint64")
	}

	integer := int(float)

	return integer, nil
}

func (c domainConverter) GetDataTransferTypeByDomainType(domainType reflect.Type) (reflect.Type, error) {
	switch domainType {
	case domain.AlbumDomainType:
		return domain.AlbumDataTransferType, nil
	case domain.ArtistDomainType:
		return domain.ArtistDataTransferType, nil
	case domain.TrackDomainType:
		return domain.TrackDataTransferType, nil
	default:
		return nil, errors.New(internal.BadType)
	}
}
