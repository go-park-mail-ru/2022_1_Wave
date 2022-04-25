package system

import (
	"errors"
	"github.com/go-park-mail-ru/2022_1_Wave/init/gRPC"
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	"github.com/go-park-mail-ru/2022_1_Wave/init/router"
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/auth/usecase"
	utilsInterfaces "github.com/go-park-mail-ru/2022_1_Wave/internal/app/interfaces"
	AlbumUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/album/useCase"
	ArtistUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/artist/useCase"
	TrackUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/track/useCase"
	structStoragePostgresql "github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/storage/postgresql"
	UserUsecase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/user/userUseCase"
	"github.com/labstack/echo/v4"
)

const local = "local"
const database = local

func Init(e *echo.Echo, quantity int64, dataBaseType string) error {
	var initedStorage utilsInterfaces.GlobalStorageInterface
	var err error
	switch dataBaseType {
	case internal.Postgres:
		initedStorage, err = structStoragePostgresql.Postgres{
			Sqlx:           nil,
			SessionRepo:    nil,
			UserRepo:       nil,
			AlbumRepo:      nil,
			AlbumCoverRepo: nil,
			ArtistRepo:     nil,
			TrackRepo:      nil,
		}.Init(quantity)
		if err != nil {
			return err
		}
	default:
		return errors.New(internal.BadType)
	}

	al := initedStorage.GetAlbumRepo()
	ar := initedStorage.GetArtistRepo()
	alc := initedStorage.GetAlbumCoverRepo()
	sess := initedStorage.GetSessionRepo()
	us := initedStorage.GetUserRepo()
	tr := initedStorage.GetTrackRepo()

	albumsQuant, err := al.GetSize()
	if err != nil {
		logger.GlobalLogger.Logrus.Fatal("Error:", err)
	}

	artistsQuant, err := ar.GetSize()
	if err != nil {
		logger.GlobalLogger.Logrus.Fatal("Error:", err)
	}

	albumCoversQuant, err := alc.GetSize()
	if err != nil {
		logger.GlobalLogger.Logrus.Fatal("Error:", err)
	}

	usersQuant, err := us.GetSize()
	if err != nil {
		logger.GlobalLogger.Logrus.Fatal("Error:", err)
	}

	tracksQuant, err := tr.GetSize()
	if err != nil {
		logger.GlobalLogger.Logrus.Fatal("Error:", err)
	}

	logger.GlobalLogger.Logrus.Info("Users:", usersQuant)
	logger.GlobalLogger.Logrus.Info("Artists:", artistsQuant)
	logger.GlobalLogger.Logrus.Info("Albums:", albumsQuant)
	logger.GlobalLogger.Logrus.Info("AlbumCovers:", albumCoversQuant)
	logger.GlobalLogger.Logrus.Info("Tracks:", tracksQuant)

	auth := AuthUseCase.NewAuthUseCase(sess, us)

	grpcLauncher := gRPC.Launcher{
		Network:      internal.LocalHost,
		AlbumServer:  AlbumUseCase.MakeAlbumService(tr, ar, al, alc),
		ArtistServer: ArtistUseCase.MakeArtistService(ar, al, tr),
		TrackServer:  TrackUseCase.MakeTrackService(tr, ar),
	}

	albumManager := grpcLauncher.LaunchAlbumService(":8081")
	artistManager := grpcLauncher.LaunchArtistService(":8082")
	trackManager := grpcLauncher.LaunchTrackService(":8083")

	user := UserUsecase.NewUserUseCase(us, sess)
	return router.Router(e, auth, albumManager, artistManager, trackManager, user)
}
