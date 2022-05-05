package main

import (
	"fmt"
	InitDb "github.com/go-park-mail-ru/2022_1_Wave/init/db"
	AlbumPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/album/repository/postgres"
	ArtistPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/artist/repository/postgres"
	PlaylistGrpc "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/playlist/gRPC"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/playlist/playlistProto"
	PlaylistPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/playlist/repository"
	TrackPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/track/repository"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
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

	httpServer := &http.Server{Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{}), Addr: fmt.Sprintf("0.0.0.0:%d", 9084)}
	server := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)

	playlistProto.RegisterPlaylistUseCaseServer(server, PlaylistGrpc.MakePlaylistGrpc(trackRepo, artistRepo, albumRepo, playlistRepo))
	grpcMetrics.InitializeMetrics(server)

	// Start your http server for prometheus.
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatal("Unable to start a http server.")
		}
	}()

	server.Serve(listen)

	//if err != nil {
	//	logger.GlobalLogger.Logrus.Errorf("cannot listen port %s: %s", port, err.Error())
	//}
}
