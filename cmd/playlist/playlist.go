package main

import (
	InitDb "github.com/go-park-mail-ru/2022_1_Wave/init/db"
	AlbumPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/album/repository/postgres"
	ArtistPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/artist/repository/postgres"
	PlaylistGrpc "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/playlist/gRPC"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/playlist/playlistProto"
	PlaylistPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/playlist/repository"
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
	playlistRepo := PlaylistPostgres.NewPlaylistPostgresRepo(sqlxDb)

	defer func() {
		if sqlxDb != nil {
			_ = sqlxDb.Close()
		}
	}()

	port := ":8084"

	listen, _ := net.Listen("tcp", port)
	//if err != nil {
	//	logger.GlobalLogger.Logrus.Errorf("error listen on %s port: %s", port, err.Error())
	//}

	server := grpc.NewServer()
	playlistProto.RegisterPlaylistUseCaseServer(server, PlaylistGrpc.MakePlaylistGrpc(trackRepo, artistRepo, albumRepo, playlistRepo))

	server.Serve(listen)

	//if err != nil {
	//	logger.GlobalLogger.Logrus.Errorf("cannot listen port %s: %s", port, err.Error())
	//}
}
