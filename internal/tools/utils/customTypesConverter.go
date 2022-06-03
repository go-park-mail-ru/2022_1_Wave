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
	for idx, obj := range artists {
		artistMap[int64(idx+1)] = obj
	}
	return artistMap
}

func AlbumsToMap(albums []*albumProto.AlbumDataTransfer) map[int64]*albumProto.AlbumDataTransfer {
	albumMap := map[int64]*albumProto.AlbumDataTransfer{}
	for idx, obj := range albums {
		albumMap[int64(idx+1)] = obj
	}
	return albumMap
}
