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

//func InitDatabase() *sqlx.DB {
//	dsn := os.Getenv("DATABASE_CONNECTION")
//	if dsn == "" {
//		dsn = "user=test dbname=test password=test host=localhost port=5500 sslmode=disable"
//	}
//	db, err := sql.Open("pgx", dsn)
//	if err != nil {
//		return nil
//	}
//	err = db.Ping() // вот тут будет первое подключение к базе
//	if err != nil {
//		return nil
//	}
//
//	db.SetMaxOpenConns(10)
//
//	sqlxDb := sqlx.NewDb(db, "pgx")
//	_ = sqlxDb.Ping()
//
//	return sqlxDb
//}

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
