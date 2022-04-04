package domain

import "sync"

var AlbumMutex = &sync.RWMutex{}
var ArtistMutex = &sync.RWMutex{}
var TrackMutex = &sync.RWMutex{}
