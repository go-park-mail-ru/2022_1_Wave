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

		albumsCover[i] = domainCreator.AlbumCoverConstructorRandom(id)
		artists[i] = domainCreator.ArtistConstructorRandom(id, nameLen, maxFollowers, maxListening)
		albums[i] = domainCreator.AlbumConstructorRandom(id, int64(quantity), albumLen, maxListening, maxLikes)

		storage.ArtistRepo, _ = storage.ArtistRepo.Insert(artists[i], domain.ArtistMutex)
		storage.AlbumRepo, _ = storage.AlbumRepo.Insert(albums[i], domain.AlbumMutex)
		storage.AlbumCoverRepo, _ = storage.AlbumCoverRepo.Insert(albumsCover[i], domain.AlbumCoverMutex)

	}

	const maxDuration = max / 10

	// tracks
	for i := 0; i < quantity; i++ {
		id := uint64(i)

		tracks[i] = domainCreator.TrackConstructorRandom(id, albums, songLen, maxDuration, maxLikes, maxListening)

		storage.TrackRepo, _ = storage.TrackRepo.Insert(tracks[i], domain.TrackMutex)
	}

	artistsCreated, _ := storage.ArtistRepo.GetAll(domain.ArtistMutex)
	albumsCreated, _ := storage.AlbumRepo.GetAll(domain.AlbumMutex)
	albumsCoverCreated, _ := storage.AlbumCoverRepo.GetAll(domain.AlbumCoverMutex)
	tracksCreated, _ := storage.TrackRepo.GetAll(domain.TrackMutex)

	artistsSize := len(artistsCreated)
	albumsSize := len(albumsCreated)
	albumsCoverSize := len(albumsCoverCreated)
	tracksSize := len(tracksCreated)

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
