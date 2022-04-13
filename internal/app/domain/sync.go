package domain

import "sync"

var AlbumMutex = &sync.RWMutex{}
var AlbumCoverMutex = &sync.RWMutex{}
var ArtistMutex = &sync.RWMutex{}
var TrackMutex = &sync.RWMutex{}
