package system

import (
	"errors"
	"github.com/go-park-mail-ru/2022_1_Wave/init/gRPC"
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	"github.com/go-park-mail-ru/2022_1_Wave/init/router"
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	AlbumUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/album/useCase"
	ArtistUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/artist/useCase"
	AuthUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/auth/usecase"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	AlbumGrpc "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/album/gRPC"
	ArtistGrpc "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/artist/gRPC"
	TrackGrpc "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/gRPC"
	structStoragePostgresql "github.com/go-park-mail-ru/2022_1_Wave/internal/structs/storage/postgresql"
	TrackUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/track/useCase"
	UserUsecase "github.com/go-park-mail-ru/2022_1_Wave/internal/user/userUseCase"
	"github.com/labstack/echo/v4"
)

const local = "local"
const database = local

type container struct {
	Al   domain.AlbumRepo
	Alc  domain.AlbumCoverRepo
	Ar   domain.ArtistRepo
	Sess domain.SessionRepo
	Us   domain.UserRepo
	Tr   domain.TrackRepo
}

func Init(e *echo.Echo, quantity int64, dataBaseType string) error {
	var initedStorage domain.GlobalStorageInterface
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

	repoContainer := container{
		Al:   initedStorage.GetAlbumRepo(),
		Alc:  initedStorage.GetAlbumCoverRepo(),
		Ar:   initedStorage.GetArtistRepo(),
		Sess: initedStorage.GetSessionRepo(),
		Us:   initedStorage.GetUserRepo(),
		Tr:   initedStorage.GetTrackRepo(),
	}

	albumsQuant, err := repoContainer.Al.GetSize()
	if err != nil {
		logger.GlobalLogger.Logrus.Fatal("Error:", err)
	}

	artistsQuant, err := repoContainer.Ar.GetSize()
	if err != nil {
		logger.GlobalLogger.Logrus.Fatal("Error:", err)
	}

	albumCoversQuant, err := repoContainer.Alc.GetSize()
	if err != nil {
		logger.GlobalLogger.Logrus.Fatal("Error:", err)
	}

	usersQuant, err := repoContainer.Us.GetSize()
	if err != nil {
		logger.GlobalLogger.Logrus.Fatal("Error:", err)
	}

	tracksQuant, err := repoContainer.Tr.GetSize()
	if err != nil {
		logger.GlobalLogger.Logrus.Fatal("Error:", err)
	}

	printDbQuants(usersQuant, artistsQuant, albumsQuant, albumCoversQuant, tracksQuant)

	auth := AuthUseCase.NewAuthUseCase(repoContainer.Sess, repoContainer.Us)

	albumClient, artistClient, trackClient := makeClients(internal.Grpc, repoContainer)

	user := UserUsecase.NewUserUseCase(repoContainer.Us, repoContainer.Sess)
	return router.Router(e, auth, albumClient, artistClient, trackClient, user)
}

func printDbQuants(usersQuant int, artistsQuant int64, albumsQuant int64, albumCoversQuant int64, tracksQuant int64) {
	logger.GlobalLogger.Logrus.Info("Users:", usersQuant)
	logger.GlobalLogger.Logrus.Info("Artists:", artistsQuant)
	logger.GlobalLogger.Logrus.Info("Albums:", albumsQuant)
	logger.GlobalLogger.Logrus.Info("AlbumCovers:", albumCoversQuant)
	logger.GlobalLogger.Logrus.Info("Tracks:", tracksQuant)
}

func makeGrpcClients(repoContainer container) (AlbumUseCase.AlbumAgent, ArtistUseCase.ArtistAgent, TrackUseCase.TrackAgent) {
	grpcLauncher := gRPC.Launcher{
		Network:      internal.Tcp,
		AlbumServer:  AlbumGrpc.MakeAlbumGrpc(repoContainer.Tr, repoContainer.Ar, repoContainer.Al, repoContainer.Alc),
		ArtistServer: ArtistGrpc.MakeArtistGrpc(repoContainer.Ar, repoContainer.Al, repoContainer.Tr),
		TrackServer:  TrackGrpc.MakeTrackGrpc(repoContainer.Tr, repoContainer.Ar),
	}

	albumClient := grpcLauncher.MakeAlbumGrpcClient(":8081")
	artistClient := grpcLauncher.MakeArtistGrpcClient(":8082")
	trackClient := grpcLauncher.MakeTrackGrpcClient(":8083")

	albumAgent := AlbumGrpc.MakeAgent(albumClient)
	artistAgent := ArtistGrpc.MakeAgent(artistClient)
	trackAgent := TrackGrpc.MakeAgent(trackClient)

	return albumAgent, artistAgent, trackAgent
}

func makeClients(clientsType string, repoContainer container) (AlbumUseCase.AlbumAgent, ArtistUseCase.ArtistAgent, TrackUseCase.TrackAgent) {
	switch clientsType {
	case internal.Grpc:
		return makeGrpcClients(repoContainer)
	default:
		return nil, nil, nil
	}

}
