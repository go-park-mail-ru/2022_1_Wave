package structStorageLocal

import (
	"errors"
	"fmt"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/interfaces"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/repository/local"
	domainCreator "github.com/go-park-mail-ru/2022_1_Wave/internal/app/tools/domain"
)

type LocalStorage struct {
	AlbumRepo      utilsInterfaces.RepoInterface `json:"albumStorage"`
	AlbumCoverRepo utilsInterfaces.RepoInterface `json:"albumCoverStorage"`
	ArtistRepo     utilsInterfaces.RepoInterface `json:"artistStorage"`
	TrackRepo      utilsInterfaces.RepoInterface `json:"trackStorage"`
}

// ----------------------------------------------------------------------
func (storage LocalStorage) Open() (utilsInterfaces.GlobalStorageInterface, error) {
	return storage, nil
}

func (storage LocalStorage) Init(quantity int) (utilsInterfaces.GlobalStorageInterface, error) {
	if quantity < 0 {
		return nil, errors.New("quantity for db is negative")
	}

	albums := make([]utilsInterfaces.Domain, quantity)
	albumsCover := make([]utilsInterfaces.Domain, quantity)
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
	storage.AlbumCoverRepo = structRepoLocal.Repo{}
	storage.TrackRepo = structRepoLocal.Repo{}

	// albums and artists
	for i := 0; i < quantity; i++ {
		id := uint64(i)

		albumsCover[i] = domainCreator.AlbumCoverConstructorRandom(id, albumLen)
		artists[i] = domainCreator.ArtistConstructorRandom(id, nameLen, maxFollowers, maxListening)
		albums[i] = domainCreator.AlbumConstructorRandom(id, int64(quantity), albumLen, maxListening, maxLikes)

		storage.ArtistRepo, _ = storage.ArtistRepo.Insert(&artists[i], domain.ArtistMutex)
		storage.AlbumRepo, _ = storage.AlbumRepo.Insert(&albums[i], domain.AlbumMutex)
		storage.AlbumCoverRepo, _ = storage.AlbumCoverRepo.Insert(&albumsCover[i], domain.AlbumCoverMutex)

	}

	const maxDuration = max / 10

	// tracks
	for i := 0; i < quantity; i++ {
		id := uint64(i)

		tracks[i] = domainCreator.TrackConstructorRandom(id, albums, artists, songLen, maxDuration, maxLikes, maxListening)

		storage.TrackRepo, _ = storage.TrackRepo.Insert(&tracks[i], domain.TrackMutex)
	}

	artistsCreated, _ := storage.ArtistRepo.GetAll(domain.ArtistMutex)
	albumsCreated, _ := storage.AlbumRepo.GetAll(domain.AlbumMutex)
	albumsCoverCreated, _ := storage.AlbumCoverRepo.GetAll(domain.AlbumCoverMutex)
	tracksCreated, _ := storage.TrackRepo.GetAll(domain.TrackMutex)

	artistsSize := len(*artistsCreated)
	albumsSize := len(*albumsCreated)
	albumsCoverSize := len(*albumsCoverCreated)
	tracksSize := len(*tracksCreated)

	if artistsSize < quantity {
		return storage, errors.New(constants.ErrorLocalDbArtistsNotEnought + ": expected " + fmt.Sprint(quantity))
	}

	if albumsSize < quantity {
		return storage, errors.New(constants.ErrorLocalDbAlbumsNotEnought + ": expected " + fmt.Sprint(quantity))
	}

	if albumsCoverSize < quantity {
		return storage, errors.New(constants.ErrorLocalDbAlbumCoversNotEnought + ": expected " + fmt.Sprint(quantity))
	}

	if tracksSize < quantity {
		return storage, errors.New(constants.ErrorLocalDbTracksNotEnought + ": expected " + fmt.Sprint(quantity))
	}

	return storage, nil
}

func (storage LocalStorage) Close() error {
	return nil
}

func (storage LocalStorage) GetAlbumRepo() *utilsInterfaces.RepoInterface {
	return &storage.AlbumRepo
}

func (storage LocalStorage) GetAlbumCoverRepo() *utilsInterfaces.RepoInterface {
	return &storage.AlbumCoverRepo
}

func (storage LocalStorage) GetArtistRepo() *utilsInterfaces.RepoInterface {
	return &storage.ArtistRepo
}

func (storage LocalStorage) GetTrackRepo() *utilsInterfaces.RepoInterface {
	return &storage.TrackRepo
}

func (storage LocalStorage) GetAlbumCoverRepoLen() (int, error) {
	all, err := storage.AlbumCoverRepo.GetAll(domain.AlbumCoverMutex)
	if err != nil {
		return 0, err
	}
	return len(*all), nil
}

func (storage LocalStorage) GetAlbumRepoLen() (int, error) {
	all, err := storage.AlbumRepo.GetAll(domain.AlbumMutex)
	if err != nil {
		return 0, err
	}
	return len(*all), nil
}

func (storage LocalStorage) GetArtistRepoLen() (int, error) {
	all, err := storage.ArtistRepo.GetAll(domain.ArtistMutex)
	if err != nil {
		return 0, err
	}
	return len(*all), nil
}

func (storage LocalStorage) GetTrackRepoLen() (int, error) {
	all, err := storage.TrackRepo.GetAll(domain.TrackMutex)
	if err != nil {
		return 0, err
	}
	return len(*all), nil
}

//// -------------------------------------------------
//func artistConstructorRandom(id uint64, maxNameLen int, maxFollowers int64, maxListening int64) domain.Artist {
//	return domain.Artist{
//		Id:             id,
//		Name:           utils.RandomWord(maxNameLen),
//		CountFollowers: uint64(rand.Int63n(maxFollowers + 1)),
//		CountListening: uint64(rand.Int63n(maxListening + 1)),
//	}
//}
//
//func albumConstructorRandom(id uint64, authorsQuantity int64, maxAlbumTitleLen int, maxLikes int64, maxListening int64) domain.Album {
//	return domain.Album{
//		Id:             id,
//		Title:          utils.RandomWord(maxAlbumTitleLen),
//		ArtistId:       uint64(rand.Int63n(authorsQuantity)),
//		CountLikes:     uint64(rand.Int63n(maxLikes + 1)),
//		CountListening: uint64(rand.Int63n(maxListening + 1)),
//		Date:           0,
//	}
//}
//
//func albumCoverConstructorRandom(id uint64, maxAlbumTitleLen int) domain.AlbumCover {
//	return domain.AlbumCover{
//		Id:     id,
//		Title:  utils.RandomWord(maxAlbumTitleLen),
//		Quote:  utils.RandomWord(100),
//		IsDark: true,
//	}
//}
//
//func trackConstructorRandom(id uint64, albums []utilsInterfaces.Domain, artists []utilsInterfaces.Domain, maxTrackTitleLen int, maxDuration int64, maxLikes int64, maxListening int64) domain.Track {
//	album := albums[rand.Intn(len(albums))]
//	albumId := album.GetId()
//
//	artist := artists[rand.Intn(len(artists))]
//	artistId := artist.GetId()
//
//	return domain.Track{
//		Id:             id,
//		AlbumId:        albumId,
//		ArtistId:       artistId,
//		Title:          utils.RandomWord(maxTrackTitleLen),
//		Duration:       uint64(rand.Int63n(maxDuration + 1)),
//		CountLikes:     uint64(rand.Int63n(maxLikes + 1)),
//		CountListening: uint64(rand.Int63n(maxListening + 1)),
//	}
//}
