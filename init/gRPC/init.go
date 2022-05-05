package gRPC

import (
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/album/albumProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/artist/artistProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/auth/proto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/playlist/playlistProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/trackProto"
	proto_user "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/user/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Launcher struct{}

func (launcher *Launcher) MakeAlbumGrpcClient(address string) albumProto.AlbumUseCaseClient {
	conn := launcher.createConnection("album", address)
	albumManager := albumProto.NewAlbumUseCaseClient(conn)
	return albumManager
}

func (launcher *Launcher) MakeArtistGrpcClient(address string) artistProto.ArtistUseCaseClient {
	conn := launcher.createConnection("artist", address)
	artistManager := artistProto.NewArtistUseCaseClient(conn)
	return artistManager
}

func (launcher *Launcher) MakeTrackGrpcClient(address string) trackProto.TrackUseCaseClient {
	conn := launcher.createConnection("track", address)
	trackManager := trackProto.NewTrackUseCaseClient(conn)
	return trackManager
}

func (launcher *Launcher) MakePlaylistGrpcClient(address string) playlistProto.PlaylistUseCaseClient {
	conn := launcher.createConnection("playlist", address)
	playlistManager := playlistProto.NewPlaylistUseCaseClient(conn)
	return playlistManager
}

func (launcher *Launcher) createConnection(host string, address string) *grpc.ClientConn {
	grpcConn, err := grpc.Dial(host+address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.GlobalLogger.Logrus.Fatalln("Error to launch grpc:", err)
	}
	return grpcConn
}

func (launcher *Launcher) MakeUserGrpcClient(address string) proto_user.ProfileClient {
	conn := launcher.createConnection("user", address)
	userClient := proto_user.NewProfileClient(conn)
	return userClient
}

func (launcher *Launcher) MakeAuthGrpcClient(address string) proto.AuthorizationClient {
	conn := launcher.createConnection("auth", address)
	authClient := proto.NewAuthorizationClient(conn)
	return authClient
}
