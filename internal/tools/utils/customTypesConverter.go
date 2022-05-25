package utils

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/album/albumProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/artist/artistProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/trackProto"
)

func TracksToMap(tracks []*trackProto.TrackDataTransfer) map[int64]*trackProto.TrackDataTransfer {
	trackMap := map[int64]*trackProto.TrackDataTransfer{}
	for _, obj := range tracks {
		trackMap[obj.Id] = obj
	}
	return trackMap
}

func ArtistsToMap(artists []*artistProto.ArtistDataTransfer) map[int64]*artistProto.ArtistDataTransfer {
	artistMap := map[int64]*artistProto.ArtistDataTransfer{}
	for _, obj := range artists {
		artistMap[obj.Id] = obj
	}
	return artistMap
}

func AlbumsToMap(albums []*albumProto.AlbumDataTransfer) map[int64]*albumProto.AlbumDataTransfer {
	albumMap := map[int64]*albumProto.AlbumDataTransfer{}
	for _, obj := range albums {
		albumMap[obj.Id] = obj
	}
	return albumMap
}
