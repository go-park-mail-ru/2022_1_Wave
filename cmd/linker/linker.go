package main

import (
	"context"
	"github.com/go-park-mail-ru/2022_1_Wave/cmd"
	InitDb "github.com/go-park-mail-ru/2022_1_Wave/init/db"
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/linker/repository/mongo"
	LinkerGrpc "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/linker/gRPC"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/linker/linkerProto"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
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

	const mongoUrlEnv = "ME_CONFIG_MONGODB_URL"
	collection, err := InitDb.InitMongo(mongoUrlEnv, "wave", "links", context.TODO())
	if err != nil {
		logs.Logrus.Fatalln(err)
	}

	logs.Logrus.Info("success connect to ", os.Getenv("dbType"), "url:", os.Getenv(mongoUrlEnv))

	linkerRepo, err := LinkerMongo.NewLinkerMongoRepo(collection)
	if err != nil {
		logs.Logrus.Fatalln(err)
	}

	server, httpServer, listen, err := cmd.MakeServers(reg)
	if err != nil {
		logs.Logrus.Fatalln("Error to launch linker gRPC service:", err)
	}
	defer listen.Close()

	linkerProto.RegisterLinkerUseCaseServer(server, LinkerGrpc.MakeLinkerGrpc(linkerRepo))

	grpcMetrics.InitializeMetrics(server)
	logs.Logrus.Info("success init metrics: linker gRPC")

	// Start your http server for prometheus.
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			logs.Logrus.Fatal("Unable to start a http linker metrics server:", err)
		}
	}()

	logs.Logrus.Info("Linker gRPC ready to listen", os.Getenv("port"))
	err = server.Serve(listen)
	if err != nil {
		logs.Logrus.Errorf("cannot listen port %s: %s", os.Getenv("port"), err)
	}

}
