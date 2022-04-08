package domain

import (
	"reflect"
)

// domains
var AlbumDomainType = reflect.TypeOf(Album{})
var AlbumCoverDomainType = reflect.TypeOf(AlbumCover{})
var ArtistDomainType = reflect.TypeOf(Artist{})
var TrackDomainType = reflect.TypeOf(Track{})

// dataTransfers
var AlbumDataTransferType = reflect.TypeOf(AlbumDataTransfer{})
var AlbumCoverDataTransferType = reflect.TypeOf(AlbumCoverDataTransfer{})
var ArtistDataTransferType = reflect.TypeOf(ArtistDataTransfer{})
var TrackDataTransferType = reflect.TypeOf(TrackDataTransfer{})
