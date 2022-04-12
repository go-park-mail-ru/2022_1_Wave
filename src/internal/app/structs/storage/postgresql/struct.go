package structStoragePostgresql

import (
	"database/sql"
	"errors"
	"github.com/go-park-mail-ru/2022_1_Wave/db"
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	authHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/app/auth/delivery/http"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/auth/delivery/http/http_middleware"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/auth/usecase"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	utilsInterfaces "github.com/go-park-mail-ru/2022_1_Wave/internal/app/interfaces"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/session/repository/redis"
	structRepoPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/repository/postgresql"
	domainCreator "github.com/go-park-mail-ru/2022_1_Wave/internal/app/tools/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/user/repository/postgresql"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"os"
	"sync"
)

type Postgres struct {
	Sqlx           *sqlx.DB
	AlbumRepo      utilsInterfaces.RepoInterface `json:"albumStorage"`
	AlbumCoverRepo utilsInterfaces.RepoInterface `json:"albumCoverStorage"`
	ArtistRepo     utilsInterfaces.RepoInterface `json:"artistStorage"`
	TrackRepo      utilsInterfaces.RepoInterface `json:"trackStorage"`
}

func (storage Postgres) getPostgres() (*sql.DB, error) {
	dsn := os.Getenv("DATABASE_CONNECTION")
	if dsn == "" {
		dsn = "user=test dbname=test password=test host=localhost port=5500 sslmode=disable"
	}
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
	if path == "" {
		path = "../../../db/migrations"
	}

	db.MigrateDB(storage.Sqlx.DB, path)

	albumTable := structRepoPostgres.Table{Sqlx: storage.Sqlx}
	albumTable, err = albumTable.SetTableName(constants.Album)
	if err != nil {
		return storage, err
	}

	albumCoverTable := structRepoPostgres.Table{Sqlx: storage.Sqlx}
	albumCoverTable, err = albumCoverTable.SetTableName(constants.AlbumCover)
	if err != nil {
		return storage, err
	}

	artistTable := structRepoPostgres.Table{Sqlx: storage.Sqlx}
	artistTable, err = artistTable.SetTableName(constants.Artist)
	if err != nil {
		return storage, err
	}

	trackTable := structRepoPostgres.Table{Sqlx: storage.Sqlx}
	trackTable, err = trackTable.SetTableName(constants.Track)
	if err != nil {
		return storage, err
	}

	storage.AlbumRepo = albumTable
	storage.AlbumCoverRepo = albumCoverTable
	storage.ArtistRepo = artistTable
	storage.TrackRepo = trackTable

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

	mutex := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	// albums cover
	coverError := make(chan error, 1)
	wg.Add(1)
	go func(wg *sync.WaitGroup, ch chan error, mutex *sync.Mutex) {
		defer wg.Done()
		for i := 0; i < quantity; i++ {
			id := uint64(i + 1)
			albumsCover[i] = domainCreator.AlbumCoverConstructorRandom(id)
			proxy, err := storage.AlbumCoverRepo.Insert(albumsCover[i], domain.AlbumMutex)
			if err != nil {
				ch <- err
				close(ch)
				return
			}
			mutex.Lock()
			storage.AlbumCoverRepo = proxy
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
			id := uint64(i + 1)
			artists[i] = domainCreator.ArtistConstructorRandom(id, nameLen, maxFollowers, maxListening)
			proxy, err := storage.ArtistRepo.Insert(artists[i], domain.ArtistMutex)
			if err != nil {
				ch <- err
				close(ch)
				return
			}
			mutex.Lock()
			storage.ArtistRepo = proxy
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
		id := uint64(i + 1)

		albums[i] = domainCreator.AlbumConstructorRandom(id, int64(quantity), albumLen, maxListening, maxLikes)

		storage.AlbumRepo, err = storage.AlbumRepo.Insert(albums[i], domain.AlbumMutex)
		if err != nil {
			return storage, err
		}
	}

	const maxDuration = max / 10
	// tracks
	for i := 0; i < quantity; i++ {
		id := uint64(i + 1)

		tracks[i] = domainCreator.TrackConstructorRandom(id, albums, songLen, maxDuration, maxLikes, maxListening)

		storage.TrackRepo, err = storage.TrackRepo.Insert(tracks[i], domain.TrackMutex)
		if err != nil {
			return storage, err
		}
	}

	sessionRepo := redis.NewRedisSessionRepo(":6379")
	userRepo := postgresql.NewUserPostgresRepo(storage.Sqlx)

	authUseCase := usecase.NewAuthUseCase(sessionRepo, userRepo)
	authHttp.M = http_middleware.InitMiddleware(authUseCase)
	authHttp.Handler = authHttp.AuthHandler{
		AuthUseCase: authUseCase,
	}

	return storage, nil
}

func (storage Postgres) Close() error {
	return nil
}

func (storage Postgres) GetAlbumRepo() *utilsInterfaces.RepoInterface {
	return &storage.AlbumRepo
}

func (storage Postgres) GetAlbumCoverRepo() *utilsInterfaces.RepoInterface {
	return &storage.AlbumCoverRepo
}

func (storage Postgres) GetArtistRepo() *utilsInterfaces.RepoInterface {
	return &storage.ArtistRepo
}

func (storage Postgres) GetTrackRepo() *utilsInterfaces.RepoInterface {
	return &storage.TrackRepo
}
