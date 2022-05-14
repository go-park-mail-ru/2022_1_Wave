package cmd

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
)

func MakeServers(reg *prometheus.Registry) (*grpc.Server, *http.Server, net.Listener, error) {
	port := os.Getenv("port")

	listen, err := net.Listen(internal.Tcp, port)
	if err != nil {
		return nil, nil, nil, err
	}
	metricsServer := &http.Server{Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{}), Addr: "0.0.0.0" + os.Getenv("exporterPort")}

	server := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)

	return server, metricsServer, listen, nil

}
