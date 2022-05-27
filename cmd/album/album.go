package main

import (
	"github.com/go-park-mail-ru/2022_1_Wave/cmd"
	InitDb "github.com/go-park-mail-ru/2022_1_Wave/init/db"
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	AlbumPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/album/repository/postgres"
	AlbumCoverPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/albumCover/repository"
	ArtistPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/artist/repository/postgres"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/album/albumProto"
	AlbumGrpc "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/album/gRPC"
	TrackPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/track/repository"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"os"
	"time"
)

var (
	// Create a metrics registry.
	reg = prometheus.NewRegistry()

	// Create some standard server metrics.
	grpcMetrics = grpc_prometheus.NewServerMetrics()

	// Create a customized counter metric.
	customizedCounterMetric = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "demo_server_say_hello_method_handle_count",
		Help: "Total number of RPCs handled on the server.",
	}, []string{"name"})
)

func init() {
	// Register standard server metrics and customized metrics to registry.
	reg.MustRegister(grpcMetrics, customizedCounterMetric)
	customizedCounterMetric.WithLabelValues("Test")
}

func main() {
	logs, err := logger.InitLogrus(os.Getenv("port"), os.Getenv("dbType"))
	if err != nil {
		log.Fatalln("error to init logrus:", err)
	}

	sqlxDb, err := InitDb.InitPostgres("DATABASE_CONNECTION")
	if err != nil {
		logs.Logrus.Fatalln("error to init database: ", os.Getenv("dbType"), "err", err)
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

	server, httpServer, listen, err := cmd.MakeServers(reg)
	if err != nil {
		logs.Logrus.Fatalln("Error to launch album gRPC service:", err)
	}
	defer listen.Close()

	albumProto.RegisterAlbumUseCaseServer(server, AlbumGrpc.MakeAlbumGrpc(trackRepo, artistRepo, albumRepo, albumCoverRepo))

	grpcMetrics.InitializeMetrics(server)
	logs.Logrus.Info("success init metrics: album gRPC")

	// Start your http server for prometheus.
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			logs.Logrus.Fatal("Unable to start a http album metrics server:", err)
		}
	}()

	//day := (time.Now().Unix() - 345600) % 604800 / 86400

	//const week = 7 * 24 * time.Hour
	if _, err := albumRepo.CountPopularAlbumOfWeek(); err != nil {
		logs.Logrus.Fatal("Unable to count a inits popular albums of week, err:", err)
	}
	logs.Logrus.Info("Success init start popular albums of week")
	go func() {
		for now := range time.Tick(time.Hour) {
			if _, err := albumRepo.CountPopularAlbumOfWeek(); err != nil {
				logs.Logrus.Fatal("Unable to count a inits popular albums of week, time:", now, "err:", err)
			}
		}
	}()

	logs.Logrus.Info("Album gRPC ready to listen", os.Getenv("port"))
	err = server.Serve(listen)
	if err != nil {
		logs.Logrus.Errorf("cannot listen port %s: %s", os.Getenv("port"), err.Error())
	}
}
