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
	PlaylistGrpcAgent "github.com/go-park-mail-ru/2022_1_Wave/internal/playlist/client/grpc"
	PlaylistUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/playlist/useCase"
	structStoragePostgresql "github.com/go-park-mail-ru/2022_1_Wave/internal/structs/storage/postgresql"
	TrackGrpcAgent "github.com/go-park-mail-ru/2022_1_Wave/internal/track/client/grpc"
	TrackUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/track/useCase"
	user_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/user"
	user_grpc_agent "github.com/go-park-mail-ru/2022_1_Wave/internal/user/client/grpc"
	UserUsecase "github.com/go-park-mail-ru/2022_1_Wave/internal/user/userUseCase"
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, quantity int64, dbType string) error {
	logger.GlobalLogger.Logrus.Infoln("in init system")
	var err error
	switch dbType {
	case internal.Postgres:
		err = structStoragePostgresql.InitPostgres(quantity)
	default:
		return errors.New(internal.BadType)
	}
	if err != nil {
		return err
	}

	logger.GlobalLogger.Logrus.Infoln("inited...")
	logger.GlobalLogger.Logrus.Infoln("success init")
	albumAgent, artistAgent, trackAgent, userAgent, authAgent, playlistAgent := makeAgents(internal.Grpc)

	auth := AuthUseCase.NewAuthService(authAgent, userAgent)
	user := UserUsecase.NewUserUseCase(userAgent, authAgent)

	album := AlbumUseCase.NewAlbumUseCase(albumAgent, artistAgent, trackAgent)
	artist := ArtistUseCase.NewArtistUseCase(albumAgent, artistAgent, trackAgent)
	track := TrackUseCase.NewTrackUseCase(albumAgent, artistAgent, trackAgent)
	playlist := PlaylistUseCase.NewPlaylistUseCase(playlistAgent, artistAgent, trackAgent)

	return router.Router(e, auth, album, artist, track, playlist, user)
}

func makeGrpcClients() (AlbumGrpcAgent.GrpcAgent, ArtistGrpcAgent.GrpcAgent, TrackGrpcAgent.GrpcAgent, user_domain.UserAgent, auth_domain2.AuthAgent, domain.PlaylistAgent) {
	grpcLauncher := gRPC.Launcher{}

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

func makeAgents(clientsType string) (domain.AlbumAgent, domain.ArtistAgent, domain.TrackAgent, user_domain.UserAgent, auth_domain2.AuthAgent, domain.PlaylistAgent) {
	switch clientsType {
	case internal.Grpc:
		return makeGrpcClients()
	default:
		return nil, nil, nil, nil, nil, nil
	}

}
