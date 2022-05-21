package main

import (
	"fmt"
	grpc_agent "github.com/go-park-mail-ru/2022_1_Wave/websocket-server/agents"
	"github.com/go-park-mail-ru/2022_1_Wave/websocket-server/auth/proto"
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
	userSyncPlayerRepo := redis.NewUserSyncElemsRepo(os.Getenv("REDIS_ADDR"))
	useCase := usecase.NewUserSyncPlayerUseCase(userSyncPlayerRepo)

	grpcConn, err := grpc.Dial(os.Getenv("AUTH_GRPC_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	fmt.Println("grpc auth con ", err)
	authClient := proto.NewAuthorizationClient(grpcConn)
	authAgent := grpc_agent.NewAuthGRPCAgent(authClient)

	handler := http.NewHandler(useCase, os.Getenv("REDIS_ADDR"), authAgent)

	middleware := middleware2.InitMiddleware(authAgent)

	e.GET("/player-sync", handler.PlayerStateLoop, middleware.Auth)

	e.Start("0.0.0.0:6789")
}
