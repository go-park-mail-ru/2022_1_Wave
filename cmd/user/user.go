package main

import (
	"database/sql"
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/user/proto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/user/repository/postgresql"
	user_service "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/user/service"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"net"
	"os"
)

func InitDatabase() *sqlx.DB {
	dsn := os.Getenv("DATABASE_CONNECTION")
	if dsn == "" {
		dsn = "user=test dbname=test password=test host=localhost port=5500 sslmode=disable"
	}
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil
	}
	err = db.Ping() // вот тут будет первое подключение к базе
	if err != nil {
		return nil
	}

	db.SetMaxOpenConns(10)

	sqlxDb := sqlx.NewDb(db, "pgx")
	_ = sqlxDb.Ping()

	return sqlxDb
}

func main() {
	sqlxDb := InitDatabase()
	userRepo := postgresql.NewUserPostgresRepo(sqlxDb)

	defer func() {
		if sqlxDb != nil {
			_ = sqlxDb.Close()
		}
	}()

	port := os.Getenv("USER_PORT")
	listen, err := net.Listen("tcp", port)
	if err != nil {
		logger.GlobalLogger.Logrus.Errorf("error listen on %s port: %s", port, err.Error())
	}

	server := grpc.NewServer()
	proto.RegisterProfileServer(server, user_service.NewUserService(userRepo))
	logger.GlobalLogger.Logrus.Printf("started profile microservice on %s", port)
	err = server.Serve(listen)
	if err != nil {
		logger.GlobalLogger.Logrus.Errorf("cannot listen port %s: %s", port, err.Error())
	}
}
