package main

import (
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/auth/proto"
	auth_redis "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/auth/repository/redis"
	auth_service "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/auth/service"
	_ "github.com/jackc/pgx/stdlib"
	"google.golang.org/grpc"
	"net"
)

func main() {
	authRepo := auth_redis.NewRedisAuthRepo("")
	//userRepo := postgresql.NewUserPostgresRepo(sqlxDb)

	/*defer func() {
		if sqlxDb != nil {
			_ = sqlxDb.Close()
		}
	}()*/

	port := ":8085"
	listen, err := net.Listen("tcp", port)
	if err != nil {
		logger.GlobalLogger.Logrus.Errorf("error listen on %s port: %s", port, err.Error())
	}

	server := grpc.NewServer()
	proto.RegisterAuthorizationServer(server, auth_service.NewAuthService(authRepo))
	//logger.GlobalLogger.Logrus.Printf("started authorization microservice on %s", port)
	err = server.Serve(listen)
	if err != nil {
		logger.GlobalLogger.Logrus.Errorf("cannot listen port %s: %s", port, err.Error())
	}
}
