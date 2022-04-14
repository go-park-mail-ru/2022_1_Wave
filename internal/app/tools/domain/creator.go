package domainCreator

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/tools/utils"
	"math/rand"
)

// -----------------------------------------
func ArtistConstructorRandom(id int, maxNameLen int, maxFollowers int, maxListening int) domain.Artist {
	//rand.Seed(time.Now().Unix())
	return domain.Artist{
		Id:             id,
		Name:           utils.RandomWord(maxNameLen),
		CountFollowers: int(rand.Intn(maxFollowers + 1)),
		CountListening: int(rand.Intn(maxListening + 1)),
	}
}

func AlbumConstructorRandom(id int, authorsQuantity int, maxAlbumTitleLen int, maxLikes int, maxListening int) domain.Album {
	//rand.Seed(time.Now().Unix())
	return domain.Album{
		Id:             id,
		Title:          utils.RandomWord(maxAlbumTitleLen),
		ArtistId:       1 + int(rand.Intn(authorsQuantity-1)),
		CountLikes:     int(rand.Intn(maxLikes + 1)),
		CountListening: int(rand.Intn(maxListening + 1)),
		Date:           0,
	}
}

func AlbumCoverConstructorRandom(id int) domain.AlbumCover {
	//rand.Seed(time.Now().Unix())
	return domain.AlbumCover{
		Id: id,
		//Title:  utils.RandomWord(maxAlbumTitleLen),
		Quote:  utils.RandomWord(100),
		IsDark: true,
	}
}

func TrackConstructorRandom(id int, albums []domain.Album, maxTrackTitleLen int, maxDuration int, maxLikes int, maxListening int) domain.Track {
	//rand.Seed(time.Now().Unix())
	album := albums[1+rand.Intn(len(albums)-1)]
	albumId := album.GetId()
	artistId := album.ArtistId

	return domain.Track{
		Id:             id,
		AlbumId:        albumId,
		ArtistId:       artistId,
		Title:          utils.RandomWord(maxTrackTitleLen),
		Duration:       int(rand.Intn(maxDuration + 1)),
		CountLikes:     int(rand.Intn(maxLikes + 1)),
		CountListening: int(rand.Intn(maxListening + 1)),
	}
}
