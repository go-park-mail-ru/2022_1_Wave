package system

import (
	"errors"
	"github.com/go-park-mail-ru/2022_1_Wave/init/gRPC"
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	"github.com/go-park-mail-ru/2022_1_Wave/init/router"
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	AlbumGrpcAgent "github.com/go-park-mail-ru/2022_1_Wave/internal/album/client/grpc"
	AlbumUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/album/useCase"
	ArtistGrpcAgent "github.com/go-park-mail-ru/2022_1_Wave/internal/artist/client/grpc"
	ArtistUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/artist/useCase"
	auth_domain2 "github.com/go-park-mail-ru/2022_1_Wave/internal/auth"
	auth_grpc_agent "github.com/go-park-mail-ru/2022_1_Wave/internal/auth/client/grpc"
	AuthUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/auth/usecase"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	AlbumGrpc "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/album/gRPC"
	ArtistGrpc "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/artist/gRPC"
	auth_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/auth"
	auth_service "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/auth/service"
	PlaylistGrpc "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/playlist/gRPC"
	TrackGrpc "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/gRPC"
	user_microservice_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/user"
	user_service "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/user/service"
	PlaylistGrpcAgent "github.com/go-park-mail-ru/2022_1_Wave/internal/playlist/client/grpc"
	PlaylistUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/playlist/useCase"
	structStoragePostgresql "github.com/go-park-mail-ru/2022_1_Wave/internal/structs/storage/postgresql"
	TrackGrpcAgent "github.com/go-park-mail-ru/2022_1_Wave/internal/track/client/grpc"
	TrackUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/track/useCase"
	user_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/user"
	user_grpc_agent "github.com/go-park-mail-ru/2022_1_Wave/internal/user/client/grpc"
	UserUsecase "github.com/go-park-mail-ru/2022_1_Wave/internal/user/usecase"
	"github.com/labstack/echo/v4"
)

const local = "local"
const database = local

type repoContainer struct {
	Al   domain.AlbumRepo
	Alc  domain.AlbumCoverRepo
	Ar   domain.ArtistRepo
	Sess auth_domain.AuthRepo
	Us   user_microservice_domain.UserRepo
	Tr   domain.TrackRepo
	Play domain.PlaylistRepo
}

func Init(e *echo.Echo, quantity int64, dataBaseType string) error {
	var initedStorage domain.GlobalStorageInterface
	var err error
	switch dataBaseType {
	case internal.Postgres:
		initedStorage = structStoragePostgresql.Postgres{
			Sqlx:           nil,
			SessionRepo:    nil,
			UserRepo:       nil,
			AlbumRepo:      nil,
			AlbumCoverRepo: nil,
			ArtistRepo:     nil,
			TrackRepo:      nil,
			PlaylistRepo:   nil,
		}

	default:
		return errors.New(internal.BadType)
	}

	initedStorage, err = initedStorage.Init(quantity)
	if err != nil {
		return err
	}

	repoContainer := repoContainer{
		Al:   initedStorage.GetAlbumRepo(),
		Alc:  initedStorage.GetAlbumCoverRepo(),
		Ar:   initedStorage.GetArtistRepo(),
		Sess: initedStorage.GetSessionRepo(),
		Us:   initedStorage.GetUserRepo(),
		Tr:   initedStorage.GetTrackRepo(),
		Play: initedStorage.GetPlaylistRepo(),
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

	playlistsQuant, err := repoContainer.Play.GetSize()
	if err != nil {
		logger.GlobalLogger.Logrus.Fatal("Error:", err)
	}

	printDbQuants(usersQuant, artistsQuant, albumsQuant, albumCoversQuant, tracksQuant, playlistsQuant)

	albumAgent, artistAgent, trackAgent, userAgent, authAgent, playlistAgent := makeAgents(internal.Grpc, repoContainer)

	auth := AuthUseCase.NewAuthService(authAgent, userAgent)
	user := UserUsecase.NewUserUseCase(userAgent, authAgent)

	album := AlbumUseCase.NewAlbumUseCase(albumAgent, artistAgent, trackAgent)
	artist := ArtistUseCase.NewArtistUseCase(albumAgent, artistAgent, trackAgent)
	track := TrackUseCase.NewTrackUseCase(albumAgent, artistAgent, trackAgent)
	playlist := PlaylistUseCase.NewPlaylistUseCase(playlistAgent, artistAgent, trackAgent)

	return router.Router(e, auth, album, artist, track, playlist, user)
}

func printDbQuants(usersQuant int, artistsQuant int64, albumsQuant int64, albumCoversQuant int64, tracksQuant int64, playlistsQuant int64) {
	logger.GlobalLogger.Logrus.Info("Users:", usersQuant)
	logger.GlobalLogger.Logrus.Info("Artists:", artistsQuant)
	logger.GlobalLogger.Logrus.Info("Albums:", albumsQuant)
	logger.GlobalLogger.Logrus.Info("AlbumCovers:", albumCoversQuant)
	logger.GlobalLogger.Logrus.Info("Tracks:", tracksQuant)
	logger.GlobalLogger.Logrus.Info("Playlists:", playlistsQuant)
}

func makeGrpcClients(repoContainer repoContainer) (AlbumGrpcAgent.GrpcAgent, ArtistGrpcAgent.GrpcAgent, TrackGrpcAgent.GrpcAgent, user_domain.UserAgent, auth_domain2.AuthAgent, domain.PlaylistAgent) {
	grpcLauncher := gRPC.Launcher{
		Network:        internal.Tcp,
		AlbumServer:    AlbumGrpc.MakeAlbumGrpc(repoContainer.Tr, repoContainer.Ar, repoContainer.Al, repoContainer.Alc),
		ArtistServer:   ArtistGrpc.MakeArtistGrpc(repoContainer.Ar, repoContainer.Al, repoContainer.Tr),
		TrackServer:    TrackGrpc.MakeTrackGrpc(repoContainer.Tr, repoContainer.Ar, repoContainer.Al),
		UserServer:     user_service.NewUserService(repoContainer.Us),
		AuthServer:     auth_service.NewAuthService(repoContainer.Sess),
		PlaylistServer: PlaylistGrpc.MakePlaylistGrpc(repoContainer.Tr, repoContainer.Ar, repoContainer.Al, repoContainer.Play),
	}

	albumClient := grpcLauncher.MakeAlbumGrpcClient(":8081")
	artistClient := grpcLauncher.MakeArtistGrpcClient(":8082")
	trackClient := grpcLauncher.MakeTrackGrpcClient(":8083")
	playlistClient := grpcLauncher.MakePlaylistGrpcClient(":8084")
	authClient := grpcLauncher.MakeAuthGrpcClient(":8085")
	userClient := grpcLauncher.MakeUserGrpcClient(":8086")

	albumAgent := AlbumGrpcAgent.MakeAgent(albumClient)
	artistAgent := ArtistGrpcAgent.MakeAgent(artistClient)
	trackAgent := TrackGrpcAgent.MakeAgent(trackClient)
	playlistAgent := PlaylistGrpcAgent.MakeAgent(playlistClient)
	authAgent := auth_grpc_agent.NewAuthGRPCAgent(authClient)
	userAgent := user_grpc_agent.NewUserGRPCAgent(userClient)

	return albumAgent, artistAgent, trackAgent, userAgent, authAgent, playlistAgent
}

func makeAgents(clientsType string, repoContainer repoContainer) (domain.AlbumAgent, domain.ArtistAgent, domain.TrackAgent, user_domain.UserAgent, auth_domain2.AuthAgent, domain.PlaylistAgent) {
	switch clientsType {
	case internal.Grpc:
		return makeGrpcClients(repoContainer)
	default:
		return nil, nil, nil, nil, nil, nil
	}

}
