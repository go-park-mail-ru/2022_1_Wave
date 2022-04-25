package domainCreator

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/album/albumProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/artist/artistProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/track/trackProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/tools/utils"
	"math/rand"
)

// -----------------------------------------
func ArtistConstructorRandom(id int64, maxNameLen int64, maxFollowers int64, maxListening int64) *artistProto.Artist {
	//rand.Seed(time.Now().Unix())
	return &artistProto.Artist{
		Id:              id,
		Name:            utils.RandomWord(maxNameLen),
		CountFollowers:  rand.Int63n(maxFollowers + 1),
		CountListenings: rand.Int63n(maxListening + 1),
	}
}

func AlbumConstructorRandom(id int64, authorsQuantity int64, maxAlbumTitleLen int64, maxLikes int64, maxListening int64) *albumProto.Album {
	//rand.Seed(time.Now().Unix())
	return &albumProto.Album{
		Id:              id,
		Title:           utils.RandomWord(maxAlbumTitleLen),
		ArtistId:        1 + rand.Int63n(authorsQuantity-1),
		CountLikes:      rand.Int63n(maxLikes + 1),
		CountListenings: rand.Int63n(maxListening + 1),
		Date:            0,
	}
}

func AlbumCoverConstructorRandom(id int64) *albumProto.AlbumCover {
	//rand.Seed(time.Now().Unix())
	return &albumProto.AlbumCover{
		Id: id,
		//Title:  utils.RandomWord(maxAlbumTitleLen),
		Quote:  utils.RandomWord(100),
		IsDark: true,
	}
}

func TrackConstructorRandom(id int64, albums []*albumProto.Album, maxTrackTitleLen int64, maxDuration int64, maxLikes int64, maxListening int64) *trackProto.Track {
	//rand.Seed(time.Now().Unix())
	album := albums[1+rand.Intn(len(albums)-1)]
	albumId := album.GetId()
	artistId := album.ArtistId

	return &trackProto.Track{
		Id:              id,
		AlbumId:         albumId,
		ArtistId:        artistId,
		Title:           utils.RandomWord(maxTrackTitleLen),
		Duration:        rand.Int63n(maxDuration + 1),
		CountLikes:      rand.Int63n(maxLikes + 1),
		CountListenings: rand.Int63n(maxListening + 1),
	}
}
