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
	TrackGrpcAgent "github.com/go-park-mail-ru/2022_1_Wave/internal/track/client/grpc"
	TrackUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/track/useCase"
	user_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/user"
	user_grpc_agent "github.com/go-park-mail-ru/2022_1_Wave/internal/user/client/grpc"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/user/client/s3"
	UserUsecase "github.com/go-park-mail-ru/2022_1_Wave/internal/user/userUseCase"
	"github.com/labstack/echo/v4"
	"os"
	"strings"
)

func Init(e *echo.Echo) error {
	logger.GlobalLogger.Logrus.Infoln("in init system")

	albumAgent, artistAgent, trackAgent, userAgent, authAgent, playlistAgent := makeAgents(internal.Grpc)
	logger.GlobalLogger.Logrus.Infoln("inited agents...")
	auth := AuthUseCase.NewAuthService(authAgent, userAgent)
	user := UserUsecase.NewUserUseCase(userAgent, authAgent)

	album := AlbumUseCase.NewAlbumUseCase(albumAgent, artistAgent, trackAgent)
	artist := ArtistUseCase.NewArtistUseCase(albumAgent, artistAgent, trackAgent)
	track := TrackUseCase.NewTrackUseCase(albumAgent, artistAgent, trackAgent)
	playlist := PlaylistUseCase.NewPlaylistUseCase(playlistAgent, artistAgent, trackAgent)
	logger.GlobalLogger.Logrus.Infoln("inited services...")
	logger.GlobalLogger.Logrus.Infoln("routing...")

	AWS_ACCESS_KEY_ID := os.Getenv("AWS_ACCESS_KEY_ID")
	AWS_SECRET_ACCESS_KEY := os.Getenv("AWS_SECRET_ACCESS_KEY")
	AWS_REGION := os.Getenv("AWS_REGION")
	AWS_S3_URL := os.Getenv("AWS_S3_URL")

	AWS_S3_URL = strings.Split(AWS_S3_URL, "\n")[0]

	if AWS_ACCESS_KEY_ID == "" {
		return errors.New("invalid AWS_ACCESS_KEY_ID:" + AWS_ACCESS_KEY_ID)
	}

	if AWS_SECRET_ACCESS_KEY == "" {
		return errors.New("invalid AWS_SECRET_ACCESS_KEY:" + AWS_SECRET_ACCESS_KEY)
	}

	if AWS_REGION == "" {
		return errors.New("invalid AWS_REGION:" + AWS_REGION)
	}

	if AWS_S3_URL == "" {
		return errors.New("invalid AWS_S3_URL:" + AWS_S3_URL)
	}

	awsConfig := &s3.AWSConfig{
		AccessKeyID:     AWS_ACCESS_KEY_ID,
		AccessKeySecret: AWS_SECRET_ACCESS_KEY,
		Region:          AWS_REGION,
		BaseURL:         AWS_S3_URL,
	}

	s3Handler := s3.MakeHandler(awsConfig)

	return router.Router(e, auth, album, artist, track, playlist, user, s3Handler)
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
