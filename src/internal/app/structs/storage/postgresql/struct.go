package structStoragePostgresql

import (
	"database/sql"
	"errors"
	"github.com/go-park-mail-ru/2022_1_Wave/db"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	utilsInterfaces2 "github.com/go-park-mail-ru/2022_1_Wave/internal/app/interfaces"
	structRepoPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/repository/postgresql"
	domainCreator "github.com/go-park-mail-ru/2022_1_Wave/internal/app/tools/domain"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"os"
	"sync"
)

type Postgres struct {
	Sqlx           *sqlx.DB
	AlbumRepo      utilsInterfaces2.RepoInterface `json:"albumStorage"`
	AlbumCoverRepo utilsInterfaces2.RepoInterface `json:"albumCoverStorage"`
	ArtistRepo     utilsInterfaces2.RepoInterface `json:"artistStorage"`
	TrackRepo      utilsInterfaces2.RepoInterface `json:"trackStorage"`
}

func (storage Postgres) getPostgres() (*sql.DB, error) {
	dsn := os.Getenv("DATABASE_CONNECTION")
	//dsn := "user=postgres dbname=wave password=music host=0.0.0.0 port=5432 sslmode=disable"
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping() // вот тут будет первое подключение к базе
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(10)
	return db, nil
}

func (storage Postgres) Open() (utilsInterfaces2.GlobalStorageInterface, error) {
	var err error
	db, err := storage.getPostgres()
	if err != nil {
		return nil, err
	}

	storage.Sqlx = sqlx.NewDb(db, "pgx")
	err = storage.Sqlx.Ping() // вот тут будет первое подключение к базе
	if err != nil {
		return nil, err
	}
	return storage, nil
}

func (storage Postgres) Init(quantity int) (utilsInterfaces2.GlobalStorageInterface, error) {
	if quantity < 0 {
		return nil, errors.New("quantity for db is negative")
	}

	proxy, err := storage.Open()
	if err != nil {
		return storage, err
	}
	storage.Sqlx = proxy.(Postgres).Sqlx

	db.MigrateDB(storage.Sqlx.DB, "./db/migrations")

	storage.AlbumRepo = structRepoPostgres.Table{Sqlx: storage.Sqlx, Name: constants.Album}
	storage.AlbumCoverRepo = structRepoPostgres.Table{Sqlx: storage.Sqlx, Name: constants.AlbumCover}
	storage.ArtistRepo = structRepoPostgres.Table{Sqlx: storage.Sqlx, Name: constants.Artist}
	storage.TrackRepo = structRepoPostgres.Table{Sqlx: storage.Sqlx, Name: constants.Track}

	albums := make([]utilsInterfaces2.Domain, quantity)
	albumsCover := make([]utilsInterfaces2.Domain, quantity)
	tracks := make([]utilsInterfaces2.Domain, quantity)
	artists := make([]utilsInterfaces2.Domain, quantity)

	const max = 10000
	const nameLen = 10
	const albumLen = 10
	const songLen = 10

	const maxFollowers = max
	const maxListening = max
	const maxLikes = max

	wg := &sync.WaitGroup{}

	// albums cover
	coverError := make(chan error, 1)
	wg.Add(1)
	go func(wg *sync.WaitGroup, ch chan error) {
		defer wg.Done()
		for i := 0; i < quantity; i++ {
			id := uint64(i)
			albumsCover[i] = domainCreator.AlbumCoverConstructorRandom(id, albumLen)
			storage.AlbumCoverRepo, err = storage.AlbumCoverRepo.Insert(&albumsCover[i], domain.AlbumCoverMutex)
			if err != nil {
				ch <- err
				return
			}
		}
		return
	}(wg, coverError)

	// artists
	artistError := make(chan error, 1)
	wg.Add(1)
	go func(wg *sync.WaitGroup, ch chan error) {
		defer wg.Done()
		for i := 0; i < quantity; i++ {
			id := uint64(i)
			artists[i] = domainCreator.ArtistConstructorRandom(id, nameLen, maxFollowers, maxListening)
			storage.ArtistRepo, err = storage.ArtistRepo.Insert(&artists[i], domain.ArtistMutex)
			if err != nil {
				ch <- err
				return
			}
		}
		return
	}(wg, artistError)

	wg.Wait()

	select {
	case err := <-coverError:
		return storage, err
	case err := <-artistError:
		return storage, err
	default:
		// next
	}

	// albums
	for i := 0; i < quantity; i++ {
		id := uint64(i)

		albums[i] = domainCreator.AlbumConstructorRandom(id, int64(quantity), albumLen, maxListening, maxLikes)

		storage.AlbumRepo, err = storage.AlbumRepo.Insert(&albums[i], domain.AlbumMutex)
		if err != nil {
			return storage, err
		}
	}

	const maxDuration = max / 10
	// tracks
	for i := 0; i < quantity; i++ {
		id := uint64(i)

		tracks[i] = domainCreator.TrackConstructorRandom(id, albums, artists, songLen, maxDuration, maxLikes, maxListening)

		storage.TrackRepo, err = storage.TrackRepo.Insert(&tracks[i], domain.TrackMutex)
		if err != nil {
			return storage, err
		}
	}

	return storage, nil
}

func (storage Postgres) Close() error {
	return nil
}

func (storage Postgres) GetAlbumRepo() *utilsInterfaces2.RepoInterface {
	return &storage.AlbumRepo
}

func (storage Postgres) GetAlbumCoverRepo() *utilsInterfaces2.RepoInterface {
	return &storage.AlbumCoverRepo
}

func (storage Postgres) GetArtistRepo() *utilsInterfaces2.RepoInterface {
	return &storage.ArtistRepo
}

func (storage Postgres) GetTrackRepo() *utilsInterfaces2.RepoInterface {
	return &storage.TrackRepo
}

func (storage Postgres) GetAlbumRepoLen() (int, error) {
	size, err := storage.AlbumRepo.GetSize(domain.AlbumMutex)
	if err != nil {
		return 0, err
	}
	return int(size), nil
}

func (storage Postgres) GetAlbumCoverRepoLen() (int, error) {
	size, err := storage.AlbumCoverRepo.GetSize(domain.AlbumCoverMutex)
	if err != nil {
		return 0, err
	}
	return int(size), nil
}

func (storage Postgres) GetArtistRepoLen() (int, error) {
	size, err := storage.ArtistRepo.GetSize(domain.ArtistMutex)
	if err != nil {
		return 0, err
	}
	return int(size), nil
}

func (storage Postgres) GetTrackRepoLen() (int, error) {
	size, err := storage.TrackRepo.GetSize(domain.TrackMutex)
	if err != nil {
		return 0, err
	}
	return int(size), nil
}
