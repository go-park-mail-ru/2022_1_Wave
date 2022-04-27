package main

import (
	"database/sql"
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	auth_usecase "github.com/go-park-mail-ru/2022_1_Wave/internal/auth/usecase"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/auth/proto"
	auth_redis "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/auth/repository/redis"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/user/repository/postgresql"
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
	authRepo := auth_redis.NewRedisAuthRepo("")
	userRepo := postgresql.NewUserPostgresRepo(sqlxDb)

	defer func() {
		if sqlxDb != nil {
			_ = sqlxDb.Close()
		}
	}()

	port := os.Getenv("AUTH_PORT")
	listen, err := net.Listen("tcp", port)
	if err != nil {
		logger.GlobalLogger.Logrus.Errorf("error listen on %s port: %s", port, err.Error())
	}

	server := grpc.NewServer()
	proto.RegisterAuthorizationServer(server, auth_usecase.NewAuthService(authRepo, userRepo))
	logger.GlobalLogger.Logrus.Printf("started authorization microservice on %s", port)
	err = server.Serve(listen)
	if err != nil {
		logger.GlobalLogger.Logrus.Errorf("cannot listen port %s: %s", port, err.Error())
	}
}
