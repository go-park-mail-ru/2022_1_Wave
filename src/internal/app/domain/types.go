package domain

import (
	"reflect"
)

// domains
var AlbumDomainType = reflect.TypeOf(Album{})
var ArtistDomainType = reflect.TypeOf(Artist{})
var TrackDomainType = reflect.TypeOf(Track{})

// dataTransfers
var AlbumDataTransferType = reflect.TypeOf(AlbumDataTransfer{})
var ArtistDataTransferType = reflect.TypeOf(ArtistDataTransfer{})
var TrackDataTransferType = reflect.TypeOf(TrackDataTransfer{})
