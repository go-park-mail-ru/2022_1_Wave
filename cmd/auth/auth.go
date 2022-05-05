package main

import (
	"fmt"
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/auth/proto"
	auth_redis "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/auth/repository/redis"
	auth_service "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/auth/service"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
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
	//sqlxDb := InitDatabase()
	authRepo := auth_redis.NewRedisAuthRepo("redis:6379")
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

	httpServer := &http.Server{Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{}), Addr: fmt.Sprintf("0.0.0.0:%d", 9085)}
	server := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)

	proto.RegisterAuthorizationServer(server, auth_service.NewAuthService(authRepo))
	grpcMetrics.InitializeMetrics(server)

	// Start your http server for prometheus.
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatal("Unable to start a http server.")
		}
	}()

	//logger.GlobalLogger.Logrus.Printf("started authorization microservice on %s", port)
	err = server.Serve(listen)
	if err != nil {
		logger.GlobalLogger.Logrus.Errorf("cannot listen port %s: %s", port, err.Error())
	}
}
