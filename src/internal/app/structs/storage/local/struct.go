package structStorageLocal

import (
	"errors"
	"fmt"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/interfaces"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/repository/local"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/tools"
	"log"
	"math/rand"
)

type LocalStorage struct {
	AlbumRepo  utilsInterfaces.RepoInterface `json:"albumStorage"`
	ArtistRepo utilsInterfaces.RepoInterface `json:"artistStorage"`
	TrackRepo  utilsInterfaces.RepoInterface `json:"trackStorage"`
}

//var Storage = structsStorage.GlobalStorageWrapper{
//	Storage: LocalStorage{},
//	Mutex:   sync.RWMutex{},
//}

// ----------------------------------------------------------------------
func (storage LocalStorage) Open() error {
	return nil
}

func (storage LocalStorage) Init(quantity int) (utilsInterfaces.GlobalStorageInterface, error) {
	//Storage.Mutex.Lock()
	//defer Storage.Mutex.Unlock()

	albums := make([]utilsInterfaces.Domain, quantity)
	tracks := make([]utilsInterfaces.Domain, quantity)
	artists := make([]utilsInterfaces.Domain, quantity)

	const max = 10000
	const nameLen = 10
	const albumLen = 10
	const songLen = 10

	const maxFollowers = max
	const maxListening = max
	const maxLikes = max

	storage.ArtistRepo = structRepoLocal.Repo{}
	storage.AlbumRepo = structRepoLocal.Repo{}
	storage.TrackRepo = structRepoLocal.Repo{}

	// albums and artists
	for i := 0; i < quantity; i++ {
		id := uint64(i)

		artists[i] = artistConstructorRandom(id, nameLen, maxFollowers, maxListening)
		albums[i] = albumConstructorRandom(id, int64(quantity), albumLen, maxListening, maxLikes)

		storage.ArtistRepo, _ = storage.ArtistRepo.Insert(&artists[i], domain.ArtistMutex)
		storage.AlbumRepo, _ = storage.AlbumRepo.Insert(&albums[i], domain.AlbumMutex)

	}

	const maxDuration = max / 10

	// tracks
	for i := 0; i < quantity; i++ {
		id := uint64(i)

		tracks[i] = trackConstructorRandom(id, albums, artists, songLen, maxDuration, maxLikes, maxListening)

		storage.TrackRepo, _ = storage.TrackRepo.Insert(&tracks[i], domain.TrackMutex)
	}

	log.Println("Success init local storage.")

	artistsCreated, _ := storage.ArtistRepo.GetAll(domain.ArtistMutex)
	albumsCreated, _ := storage.AlbumRepo.GetAll(domain.AlbumMutex)
	tracksCreated, _ := storage.TrackRepo.GetAll(domain.TrackMutex)

	artistsSize := len(*artistsCreated)
	albumsSize := len(*albumsCreated)
	tracksSize := len(*tracksCreated)

	if artistsSize < quantity {
		return storage, errors.New(constants.ErrorLocalDbArtistsNotEnought + ": expected " + fmt.Sprint(quantity))
	}

	if albumsSize < quantity {
		return storage, errors.New(constants.ErrorLocalDbAlbumsNotEnought + ": expected " + fmt.Sprint(quantity))
	}

	if tracksSize < quantity {
		return storage, errors.New(constants.ErrorLocalDbTracksNotEnought + ": expected " + fmt.Sprint(quantity))
	}

	log.Println("Artists:", len(*artistsCreated))
	log.Println("Albums:", len(*albumsCreated))
	log.Println("Tracks:", len(*tracksCreated))

	return storage, nil
}

func (storage LocalStorage) Close() error {
	return nil
}

func (storage LocalStorage) GetAlbumRepo() *utilsInterfaces.RepoInterface {
	return &storage.AlbumRepo
}
func (storage LocalStorage) GetArtistRepo() *utilsInterfaces.RepoInterface {
	return &storage.ArtistRepo
}
func (storage LocalStorage) GetTrackRepo() *utilsInterfaces.RepoInterface {
	return &storage.TrackRepo
}

// -------------------------------------------------
func artistConstructorRandom(id uint64, maxNameLen int, maxFollowers int64, maxListening int64) domain.Artist {
	return domain.Artist{
		Id:             id,
		Name:           tools.RandomWord(maxNameLen),
		Photo:          "assets/artist_" + fmt.Sprint(id) + ".png",
		CountFollowers: uint64(rand.Int63n(maxFollowers + 1)),
		CountListening: uint64(rand.Int63n(maxListening + 1)),
	}
}

func albumConstructorRandom(id uint64, authorsQuantity int64, maxAlbumTitleLen int, maxLikes int64, maxListening int64) domain.Album {
	return domain.Album{
		Id:             id,
		Title:          tools.RandomWord(maxAlbumTitleLen),
		ArtistId:       uint64(rand.Int63n(authorsQuantity)),
		CountLikes:     uint64(rand.Int63n(maxLikes + 1)),
		CountListening: uint64(rand.Int63n(maxListening + 1)),
		Date:           0,
		CoverId:        id,
	}
}

func trackConstructorRandom(id uint64, albums []utilsInterfaces.Domain, artists []utilsInterfaces.Domain, maxTrackTitleLen int, maxDuration int64, maxLikes int64, maxListening int64) domain.Track {
	album := albums[rand.Intn(len(albums))]
	albumId := album.GetId()

	artist := artists[rand.Intn(len(artists))]
	artistId := artist.GetId()

	return domain.Track{
		Id:             id,
		AlbumId:        albumId,
		ArtistId:       artistId,
		Title:          tools.RandomWord(maxTrackTitleLen),
		Duration:       uint64(rand.Int63n(maxDuration + 1)),
		Mp4:            "assets/track_" + fmt.Sprint(id) + ".mp4",
		CoverId:        id,
		CountLikes:     uint64(rand.Int63n(maxLikes + 1)),
		CountListening: uint64(rand.Int63n(maxListening + 1)),
	}
}
