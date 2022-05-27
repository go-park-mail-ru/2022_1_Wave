package main

import (
	"github.com/go-park-mail-ru/2022_1_Wave/cmd"
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/auth/proto"
	auth_redis "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/auth/repository/redis"
	auth_service "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/auth/service"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/prometheus/client_golang/prometheus"
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
		logs.Logrus.Fatalln("Error to launch auth gRPC service:", err)
	}
	//sqlxDb := InitDatabase()
	authRepo := auth_redis.NewRedisAuthRepo("redis:6379")
	//userRepo := postgresql.NewUserPostgresRepo(sqlxDb)

	/*defer func() {
		if sqlxDb != nil {
			_ = sqlxDb.Close()
		}
	}()*/

	server, httpServer, listen, err := cmd.MakeServers(reg)
	if err != nil {
		logs.Logrus.Fatalln("error to init database: ", os.Getenv("dbType"), err)
	}
	defer listen.Close()

	proto.RegisterAuthorizationServer(server, auth_service.NewAuthService(authRepo))
	grpcMetrics.InitializeMetrics(server)
	logs.Logrus.Info("success init metrics: auth gRPC")
	// Start your http server for prometheus.
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			logs.Logrus.Fatal("Unable to start a http auth metrics server:", err)
		}
	}()

	//logger.GlobalLogger.Logrus.Printf("started authorization microservice on %s", port)
	err = server.Serve(listen)
	if err != nil {
		logs.Logrus.Errorf("cannot listen port %s: %s", os.Getenv("port"), err.Error())
	}
}
