package domainCreator

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/tools/utils"
	"math/rand"
)

// -----------------------------------------
func ArtistConstructorRandom(id int64, maxNameLen int64, maxFollowers int64, maxListening int64) domain.Artist {
	//rand.Seed(time.Now().Unix())
	return domain.Artist{
		Id:             id,
		Name:           utils.RandomWord(maxNameLen),
		CountFollowers: rand.Int63n(maxFollowers + 1),
		CountListening: rand.Int63n(maxListening + 1),
	}
}

func AlbumConstructorRandom(id int64, authorsQuantity int64, maxAlbumTitleLen int64, maxLikes int64, maxListening int64) domain.Album {
	//rand.Seed(time.Now().Unix())
	return domain.Album{
		Id:             id,
		Title:          utils.RandomWord(maxAlbumTitleLen),
		ArtistId:       1 + rand.Int63n(authorsQuantity-1),
		CountLikes:     rand.Int63n(maxLikes + 1),
		CountListening: rand.Int63n(maxListening + 1),
		Date:           0,
	}
}

func AlbumCoverConstructorRandom(id int64) domain.AlbumCover {
	//rand.Seed(time.Now().Unix())
	return domain.AlbumCover{
		Id: id,
		//Title:  utils.RandomWord(maxAlbumTitleLen),
		Quote:  utils.RandomWord(100),
		IsDark: true,
	}
}

func TrackConstructorRandom(id int64, albums []domain.Album, maxTrackTitleLen int64, maxDuration int64, maxLikes int64, maxListening int64) domain.Track {
	//rand.Seed(time.Now().Unix())
	album := albums[1+rand.Intn(len(albums)-1)]
	albumId := album.GetId()
	artistId := album.ArtistId

	return domain.Track{
		Id:             id,
		AlbumId:        albumId,
		ArtistId:       artistId,
		Title:          utils.RandomWord(maxTrackTitleLen),
		Duration:       rand.Int63n(maxDuration + 1),
		CountLikes:     rand.Int63n(maxLikes + 1),
		CountListening: rand.Int63n(maxListening + 1),
	}
}
