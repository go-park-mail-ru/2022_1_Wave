package utilsInterfaces

import (
	"reflect"
)

type CreatorInterface interface {
	InitDomainAlbum(id uint64) (Domain, error)
	InitDomainArtist(id uint64) (Domain, error)
	InitDomainTrack(id uint64) (Domain, error)
	InitDomain(id uint64, domainType reflect.Type) (Domain, error)
	CreateDomainAlbumFromInterface(data interface{}) (Domain, error)
	CreateDomainAlbumCoverFromInterface(data interface{}) (Domain, error)
	CreateDomainArtistFromInterface(data interface{}) (Domain, error)
	CreateDomainTrackFromInterface(data interface{}) (Domain, error)
	CreateDataTransferFromInterface(toType reflect.Type, data interface{}) (DataTransfer, error)
}
