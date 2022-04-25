package gRPC

import (
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/album/albumProto"
	AlbumUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/album/useCase"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/artist/artistProto"
	ArtistUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/artist/useCase"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/track/trackProto"
	TrackUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/track/useCase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
)

type Launcher struct {
	Network      string
	AlbumServer  AlbumUseCase.AlbumUseCase
	ArtistServer ArtistUseCase.ArtistUseCase
	TrackServer  TrackUseCase.TrackUseCase
}

func (launcher *Launcher) LaunchAlbumService(address string) albumProto.AlbumUseCaseClient {
	server := grpc.NewServer()
	albumProto.RegisterAlbumUseCaseServer(server, launcher.AlbumServer)

	conn := launcher.createConnection(address, server)

	albumManager := albumProto.NewAlbumUseCaseClient(conn)
	return albumManager
}

func (launcher *Launcher) LaunchArtistService(address string) artistProto.ArtistUseCaseClient {
	server := grpc.NewServer()
	artistProto.RegisterArtistUseCaseServer(server, launcher.ArtistServer)

	conn := launcher.createConnection(address, server)

	artistManager := artistProto.NewArtistUseCaseClient(conn)
	return artistManager
}

func (launcher *Launcher) LaunchTrackService(address string) trackProto.TrackUseCaseClient {
	server := grpc.NewServer()
	trackProto.RegisterTrackUseCaseServer(server, launcher.TrackServer)

	conn := launcher.createConnection(address, server)

	trackManager := trackProto.NewTrackUseCaseClient(conn)
	return trackManager
}

func (launcher *Launcher) createConnection(address string, server *grpc.Server) *grpc.ClientConn {
	listener, err := net.Listen(launcher.Network, address)
	go server.Serve(listener)
	grpcConn, err := grpc.Dial(internal.LocalHost+address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.GlobalLogger.Logrus.Fatalln("Error to launch grpc:", err)
	}
	return grpcConn
	//defer grpcConn.Close()
}
