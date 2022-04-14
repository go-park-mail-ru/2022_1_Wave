package structStoragePostgresql

import (
	"database/sql"
	"errors"
	"github.com/go-park-mail-ru/2022_1_Wave/db"
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	AlbumPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/app/album/repository"
	AlbumCoverPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/app/albumCover/repository"
	ArtistPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/app/artist/repository/postgres"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	utilsInterfaces "github.com/go-park-mail-ru/2022_1_Wave/internal/app/interfaces"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/session/repository/redis"
	domainCreator "github.com/go-park-mail-ru/2022_1_Wave/internal/app/tools/domain"
	TrackPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/app/track/repository"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/user/repository/postgresql"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"os"
	"sync"
)

type Postgres struct {
	Sqlx           *sqlx.DB
	SessionRepo    *domain.SessionRepo
	UserRepo       *domain.UserRepo
	AlbumRepo      *AlbumPostgres.AlbumRepo
	AlbumCoverRepo *AlbumCoverPostgres.AlbumCoverRepo
	ArtistRepo     *ArtistPostgres.ArtistRepo
	TrackRepo      *TrackPostgres.TrackRepo
}

func (storage Postgres) getPostgres() (*sql.DB, error) {
	dsn := os.Getenv("DATABASE_CONNECTION")
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

func (storage Postgres) Open() (utilsInterfaces.GlobalStorageInterface, error) {
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

func (storage Postgres) Init(quantity int) (utilsInterfaces.GlobalStorageInterface, error) {
	if quantity < 0 {
		return nil, errors.New("quantity for db is negative")
	}

	proxy, err := storage.Open()
	if err != nil {
		return storage, err
	}
	storage.Sqlx = proxy.(Postgres).Sqlx

	logger.GlobalLogger.Logrus.Warnln("Finding migrations...")
	path := os.Getenv("DATABASE_MIGRATIONS")

	db.MigrateDB(storage.Sqlx.DB, path)

	storage.AlbumRepo = &AlbumPostgres.AlbumRepo{Sqlx: storage.Sqlx}
	storage.AlbumCoverRepo = &AlbumCoverPostgres.AlbumCoverRepo{Sqlx: storage.Sqlx}
	storage.ArtistRepo = &ArtistPostgres.ArtistRepo{Sqlx: storage.Sqlx}
	storage.TrackRepo = &TrackPostgres.TrackRepo{Sqlx: storage.Sqlx}

	userRepo := postgresql.NewUserPostgresRepo(storage.Sqlx)
	storage.UserRepo = &userRepo

	sessionRepo := redis.NewRedisSessionRepo("redis:6379")
	storage.SessionRepo = &sessionRepo

	albums := make([]domain.Album, quantity)
	albumsCover := make([]domain.AlbumCover, quantity)
	tracks := make([]domain.Track, quantity)
	artists := make([]domain.Artist, quantity)

	const max = 10000
	const nameLen = 10
	const albumLen = 10
	const songLen = 10

	const maxFollowers = max
	const maxListening = max
	const maxLikes = max

	mutex := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	// albums cover
	coverError := make(chan error, 1)
	wg.Add(1)
	go func(wg *sync.WaitGroup, ch chan error, mutex *sync.Mutex) {
		defer wg.Done()
		for i := 0; i < quantity; i++ {
			id := i + 1
			albumsCover[i] = domainCreator.AlbumCoverConstructorRandom(id)
			if err := storage.AlbumCoverRepo.Insert(albumsCover[i]); err != nil {
				ch <- err
				close(ch)
				return
			}
			mutex.Lock()
			mutex.Unlock()
		}
		close(ch)
		return
	}(wg, coverError, mutex)

	// artists
	artistError := make(chan error, 1)
	wg.Add(1)
	go func(wg *sync.WaitGroup, ch chan error, mutex *sync.Mutex) {
		defer wg.Done()
		for i := 0; i < quantity; i++ {
			id := i + 1
			artists[i] = domainCreator.ArtistConstructorRandom(id, nameLen, maxFollowers, maxListening)
			if err := storage.ArtistRepo.Insert(artists[i]); err != nil {
				ch <- err
				close(ch)
				return
			}
			mutex.Lock()
			mutex.Unlock()
		}
		close(ch)
		return
	}(wg, artistError, mutex)

	wg.Wait()

	for err := range coverError {
		if err != nil {
			return nil, err
		}
	}

	for err := range artistError {
		if err != nil {
			return nil, err
		}
	}

	// albums
	for i := 0; i < quantity; i++ {
		id := i + 1
		albums[i] = domainCreator.AlbumConstructorRandom(id, quantity, albumLen, maxListening, maxLikes)
		if err := storage.AlbumRepo.Insert(albums[i]); err != nil {
			return storage, err
		}
	}

	const maxDuration = max / 10
	// tracks
	for i := 0; i < quantity; i++ {
		id := i + 1
		tracks[i] = domainCreator.TrackConstructorRandom(id, albums, songLen, maxDuration, maxLikes, maxListening)
		if err := storage.TrackRepo.Insert(tracks[i]); err != nil {
			return storage, err
		}
	}

	return storage, nil
}

func (storage Postgres) Close() error {
	return nil
}

//func (storage Postgres) GetAlbumRepo() *utilsInterfaces.AlbumRepo {
//	return storage.AlbumRepo
//}
//
//func (storage Postgres) GetAlbumCoverRepo() *utilsInterfaces.RepoInterface {
//	return &storage.AlbumCoverRepo
//}
//
//func (storage Postgres) GetArtistRepo() *utilsInterfaces.RepoInterface {
//	return &storage.ArtistRepo
//}
//
//func (storage Postgres) GetTrackRepo() *utilsInterfaces.RepoInterface {
//	return &storage.TrackRepo
//}
