package gRPC

import (
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/album/albumProto"
	AlbumGrpc "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/album/gRPC"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/artist/artistProto"
	ArtistGrpc "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/artist/gRPC"
	TrackGrpc "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/gRPC"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/trackProto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
)

type Launcher struct {
	Network      string
	AlbumServer  AlbumGrpc.AlbumGrpc
	ArtistServer ArtistGrpc.ArtistGrpc
	TrackServer  TrackGrpc.TrackGrpc
}

func (launcher *Launcher) MakeAlbumGrpcClient(address string) albumProto.AlbumUseCaseClient {
	server := grpc.NewServer()
	albumProto.RegisterAlbumUseCaseServer(server, launcher.AlbumServer)

	conn := launcher.createConnection(address, server)

	albumManager := albumProto.NewAlbumUseCaseClient(conn)
	return albumManager
}

func (launcher *Launcher) MakeArtistGrpcClient(address string) artistProto.ArtistUseCaseClient {
	server := grpc.NewServer()
	artistProto.RegisterArtistUseCaseServer(server, launcher.ArtistServer)

	conn := launcher.createConnection(address, server)

	artistManager := artistProto.NewArtistUseCaseClient(conn)
	return artistManager
}

func (launcher *Launcher) MakeTrackGrpcClient(address string) trackProto.TrackUseCaseClient {
	server := grpc.NewServer()
	trackProto.RegisterTrackUseCaseServer(server, launcher.TrackServer)

	conn := launcher.createConnection(address, server)

	trackManager := trackProto.NewTrackUseCaseClient(conn)
	return trackManager
}

func (launcher *Launcher) createConnection(address string, server *grpc.Server) *grpc.ClientConn {
	listener, err := net.Listen(launcher.Network, address)
	if err != nil {
		logger.GlobalLogger.Logrus.Fatalln("Error to launch grpc:", err)
	}
	go server.Serve(listener)
	grpcConn, err := grpc.Dial(internal.LocalHost+address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.GlobalLogger.Logrus.Fatalln("Error to launch grpc:", err)
	}
	return grpcConn
	//defer grpcConn.Close()
}
