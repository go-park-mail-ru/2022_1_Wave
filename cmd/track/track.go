package main

import (
	"github.com/go-park-mail-ru/2022_1_Wave/cmd"
	InitDb "github.com/go-park-mail-ru/2022_1_Wave/init/db"
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	AlbumPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/album/repository/postgres"
	ArtistPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/artist/repository/postgres"
	TrackGrpc "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/gRPC"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/trackProto"
	TrackPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/track/repository"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"os"
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

	sqlxDb, err := InitDb.InitDatabase("DATABASE_CONNECTION")
	if err != nil {
		logs.Logrus.Fatalln("error to init database: ", os.Getenv("dbType"), err)
	}

	trackRepo := TrackPostgres.NewTrackPostgresRepo(sqlxDb)
	artistRepo := ArtistPostgres.NewArtistPostgresRepo(sqlxDb)
	albumRepo := AlbumPostgres.NewAlbumPostgresRepo(sqlxDb)

	defer func() {
		if sqlxDb != nil {
			_ = sqlxDb.Close()
		}
	}()
	server, httpServer, listen, err := cmd.MakeServers(reg)
	if err != nil {
		logs.Logrus.Fatalln("Error to launch track gRPC service")
	}
	defer listen.Close()

	trackProto.RegisterTrackUseCaseServer(server, TrackGrpc.MakeTrackGrpc(trackRepo, artistRepo, albumRepo))
	grpcMetrics.InitializeMetrics(server)
	grpc_prometheus.EnableHandlingTimeHistogram()
	// Start your http server for prometheus.
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			logs.Logrus.Fatal("Unable to start a http track metrics server.")
		}
	}()

	err = server.Serve(listen)
	if err != nil {
		logs.Logrus.Errorf("cannot listen port %s: %s", os.Getenv("port"), err.Error())
	}

	//if err != nil {
	//	logger.GlobalLogger.Logrus.Errorf("cannot listen port %s: %s", port, err.Error())
	//}
}
