package main

import (
	"fmt"
	InitDb "github.com/go-park-mail-ru/2022_1_Wave/init/db"
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/user/proto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/user/repository/postgresql"
	user_service "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/user/service"
	_ "github.com/jackc/pgx/stdlib"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main() {
	sqlxDb, err := InitDb.InitDatabase("DATABASE_CONNECTION")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	userRepo := postgresql.NewUserPostgresRepo(sqlxDb)

	defer func() {
		if sqlxDb != nil {
			_ = sqlxDb.Close()
		}
	}()

	port := ":8086"
	listen, err := net.Listen("tcp", port)
	if err != nil {
		logger.GlobalLogger.Logrus.Errorf("error listen on %s port: %s", port, err.Error())
	}

	server := grpc.NewServer()
	proto.RegisterProfileServer(server, user_service.NewUserService(userRepo))
	//logger.GlobalLogger.Logrus.Printf("started profile microservice on %s", port)
	err = server.Serve(listen)
	if err != nil {
		logger.GlobalLogger.Logrus.Errorf("cannot listen port %s: %s", port, err.Error())
	}
}
