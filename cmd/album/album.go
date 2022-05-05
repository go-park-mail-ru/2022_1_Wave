package main

import (
	"fmt"
	InitDb "github.com/go-park-mail-ru/2022_1_Wave/init/db"
	AlbumPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/album/repository/postgres"
	AlbumCoverPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/albumCover/repository"
	ArtistPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/artist/repository/postgres"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/album/albumProto"
	AlbumGrpc "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/album/gRPC"
	TrackPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/track/repository"
	_ "github.com/jackc/pgx/stdlib"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main() {
	sqlxDb, err := InitDb.InitDatabase("DATABASE_CONNECTION")
	if err != nil {
		os.Exit(1)
	}
	trackRepo := TrackPostgres.NewTrackPostgresRepo(sqlxDb)
	artistRepo := ArtistPostgres.NewArtistPostgresRepo(sqlxDb)
	albumRepo := AlbumPostgres.NewAlbumPostgresRepo(sqlxDb)
	albumCoverRepo := AlbumCoverPostgres.NewAlbumCoverPostgresRepo(sqlxDb)

	defer func() {
		if sqlxDb != nil {
			_ = sqlxDb.Close()
		}
	}()

	port := ":8081"

	listen, _ := net.Listen("tcp", port)
	//if err != nil {
	//	logger.GlobalLogger.Logrus.Errorf("error listen on %s port: %s", port, err.Error())
	//}

	server := grpc.NewServer()
	albumProto.RegisterAlbumUseCaseServer(server, AlbumGrpc.MakeAlbumGrpc(trackRepo, artistRepo, albumRepo, albumCoverRepo))

	err = server.Serve(listen)

	if err != nil {
		fmt.Println("here")
	}
}
