package main

import (
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	auth_grpc_agent "github.com/go-park-mail-ru/2022_1_Wave/internal/auth/client/grpc"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/auth/proto"
	"github.com/go-park-mail-ru/2022_1_Wave/websocket-server/delivery/http"
	middleware2 "github.com/go-park-mail-ru/2022_1_Wave/websocket-server/delivery/middleware"
	"github.com/go-park-mail-ru/2022_1_Wave/websocket-server/repository/redis"
	"github.com/go-park-mail-ru/2022_1_Wave/websocket-server/usecase"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

func main() {
	e := echo.New()

	logs, err := logger.InitLogrus("6789", "redis-sync-player")
	if err != nil {
		e.Logger.Fatalf("error to init logrus:", err)
	}

	e.Use(logs.ColoredLogMiddleware)
	e.Use(logs.JsonLogMiddleware)
	e.Logger.SetOutput(logs.Logrus.Writer())

	userSyncPlayerRepo := redis.NewUserSyncElemsRepo(os.Getenv("REDIS_ADDR"))
	useCase := usecase.NewUserSyncPlayerUseCase(userSyncPlayerRepo)

	grpcAuthConn, err := grpc.Dial(os.Getenv("AUTH_GRPC_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))

	authClient := proto.NewAuthorizationClient(grpcAuthConn)
	authAgent := auth_grpc_agent.NewAuthGRPCAgent(authClient)

	/*grpcTrackConn, err := grpc.Dial(os.Getenv("TRACK_GRPC_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))

	trackClient := trackProto.NewTrackUseCaseClient(grpcTrackConn)
	trackAgent := TrackGrpcAgent.MakeAgent(trackClient)*/

	handler := http.NewHandler(useCase, os.Getenv("REDIS_ADDR"), authAgent)

	middleware := middleware2.InitMiddleware(authAgent)

	e.GET("/api/v1/player-sync", handler.PlayerStateLoop, middleware.Auth)

	e.Start("0.0.0.0:6789")
}
