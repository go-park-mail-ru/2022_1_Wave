package structStoragePostgresql

import (
	"errors"
	"github.com/go-park-mail-ru/2022_1_Wave/db"
	InitDb "github.com/go-park-mail-ru/2022_1_Wave/init/db"
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	AlbumPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/album/repository/postgres"
	AlbumCoverPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/albumCover/repository"
	ArtistPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/artist/repository/postgres"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/album/albumProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/artist/artistProto"
	auth_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/auth"
	auth_redis "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/auth/repository/redis"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/trackProto"
	user_microservice_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/user"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/user/repository/postgresql"
	PlaylistPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/playlist/repository"
	domainCreator "github.com/go-park-mail-ru/2022_1_Wave/internal/tools/domain"
	TrackPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/track/repository"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"os"
	"sync"
)

type Postgres struct {
	Sqlx           *sqlx.DB
	SessionRepo    auth_domain.AuthRepo
	UserRepo       user_microservice_domain.UserRepo
	AlbumRepo      domain.AlbumRepo
	AlbumCoverRepo domain.AlbumCoverRepo
	ArtistRepo     domain.ArtistRepo
	TrackRepo      domain.TrackRepo
	PlaylistRepo   domain.PlaylistRepo
}

func InitPostgres(quantity int64) error {
	storage := Postgres{}
	if quantity < 0 {
		return errors.New("quantity for db is negative")
	}

	sqlxDb, err := InitDb.InitDatabase("DATABASE_CONNECTION")
	logger.GlobalLogger.Logrus.Infoln("Success init db in init...")
	if err != nil {
		return err
	}

	storage.Sqlx = sqlxDb

	path := os.Getenv("DATABASE_MIGRATIONS")
	if path == "" {
		return nil
	}

	// migrate and generate
	logger.GlobalLogger.Logrus.Warnln("Finding migrations...")
	db.MigrateDB(storage.Sqlx.DB, path)

	storage.AlbumRepo = &AlbumPostgres.AlbumRepo{Sqlx: storage.Sqlx}
	storage.AlbumCoverRepo = &AlbumCoverPostgres.AlbumCoverRepo{Sqlx: storage.Sqlx}
	storage.ArtistRepo = &ArtistPostgres.ArtistRepo{Sqlx: storage.Sqlx}
	storage.TrackRepo = &TrackPostgres.TrackRepo{Sqlx: storage.Sqlx}
	storage.UserRepo = postgresql.NewUserPostgresRepo(storage.Sqlx)
	storage.SessionRepo = auth_redis.NewRedisAuthRepo("redis:6379")
	storage.PlaylistRepo = PlaylistPostgres.PlaylistRepo{Sqlx: storage.Sqlx}

	albums := make([]*albumProto.Album, quantity)
	albumsCover := make([]*albumProto.AlbumCover, quantity)
	tracks := make([]*trackProto.Track, quantity)
	artists := make([]*artistProto.Artist, quantity)

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
		for i := int64(0); i < quantity; i++ {
			id := i + 1
			albumsCover[i] = domainCreator.AlbumCoverConstructorRandom(id)
			if err := storage.AlbumCoverRepo.Create(albumsCover[i]); err != nil {
				ch <- err
				break
				//close(ch)
				//return
			}
		}
		close(ch)
		//return
	}(wg, coverError, mutex)

	// artists
	artistError := make(chan error, 1)
	wg.Add(1)
	go func(wg *sync.WaitGroup, ch chan error, mutex *sync.Mutex) {
		defer wg.Done()
		for i := int64(0); i < quantity; i++ {
			id := i + 1
			artists[i] = domainCreator.ArtistConstructorRandom(id, nameLen, maxFollowers, maxListening)
			if err := storage.ArtistRepo.Create(artists[i]); err != nil {
				ch <- err
				break
				//close(ch)
				//return
			}
		}
		close(ch)
		//return
	}(wg, artistError, mutex)

	wg.Wait()

	for err := range coverError {
		if err != nil {
			return err
		}
	}

	for err := range artistError {
		if err != nil {
			return err
		}
	}

	// albums
	for i := int64(0); i < quantity; i++ {
		id := i + 1
		albums[i] = domainCreator.AlbumConstructorRandom(id, quantity, albumLen, maxListening, maxLikes)
		if err := storage.AlbumRepo.Create(albums[i]); err != nil {
			return err
		}
	}

	const maxDuration = max / 10
	// tracks
	for i := int64(0); i < quantity; i++ {
		id := i + 1
		tracks[i] = domainCreator.TrackConstructorRandom(id, albums, songLen, maxDuration, maxLikes, maxListening)
		if err := storage.TrackRepo.Create(tracks[i]); err != nil {
			return err
		}
	}

	return nil
}

func (storage Postgres) Close() error {
	return nil
}

func (storage Postgres) GetAlbumRepo() domain.AlbumRepo {
	return storage.AlbumRepo
}

func (storage Postgres) GetAlbumCoverRepo() domain.AlbumCoverRepo {
	return storage.AlbumCoverRepo
}

func (storage Postgres) GetArtistRepo() domain.ArtistRepo {
	return storage.ArtistRepo
}

func (storage Postgres) GetTrackRepo() domain.TrackRepo {
	return storage.TrackRepo
}

func (storage Postgres) GetSessionRepo() auth_domain.AuthRepo {
	return storage.SessionRepo
}

func (storage Postgres) GetUserRepo() user_microservice_domain.UserRepo {
	return storage.UserRepo
}

func (storage Postgres) GetPlaylistRepo() domain.PlaylistRepo {
	return storage.PlaylistRepo
}
