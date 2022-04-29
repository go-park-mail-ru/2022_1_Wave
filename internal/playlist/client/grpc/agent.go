package PlaylistGrpcAgent

import (
	"context"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway/gatewayProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/playlist/playlistProto"
)

type GrpcAgent struct {
	PlaylistGrpc playlistProto.PlaylistUseCaseClient
}

func MakeAgent(gRPC playlistProto.PlaylistUseCaseClient) GrpcAgent {
	return GrpcAgent{
		PlaylistGrpc: gRPC,
	}
}

func (agent GrpcAgent) GetAll(userId int64) ([]*playlistProto.Playlist, error) {
	data, err := agent.PlaylistGrpc.GetAll(context.Background(), &gatewayProto.IdArg{Id: userId})
	return data.Playlists, err
}

func (agent GrpcAgent) GetLastId(userId int64) (int64, error) {
	id, err := agent.PlaylistGrpc.GetLastId(context.Background(), &gatewayProto.IdArg{Id: userId})
	return id.Data, err
}

func (agent GrpcAgent) Update(userId int64, playlist *playlistProto.Playlist) error {
	_, err := agent.PlaylistGrpc.Update(context.Background(), &playlistProto.UserIdPlaylistArg{
		UserId:   userId,
		Playlist: playlist,
	})
	return err
}
func (agent GrpcAgent) Delete(userId int64, playlistId int64) error {
	_, err := agent.PlaylistGrpc.Delete(context.Background(), &playlistProto.UserIdPlaylistIdArg{
		UserId:     userId,
		PlaylistId: playlistId,
	})
	return err
}
func (agent GrpcAgent) GetById(userId int64, playlistId int64) (*playlistProto.Playlist, error) {
	return agent.PlaylistGrpc.GetById(context.Background(), &playlistProto.UserIdPlaylistIdArg{
		UserId:     userId,
		PlaylistId: playlistId,
	})
}
func (agent GrpcAgent) GetSize(userId int64) (int64, error) {
	size, err := agent.PlaylistGrpc.GetSize(context.Background(), &gatewayProto.IdArg{Id: userId})
	return size.Data, err
}

func (agent GrpcAgent) Create(userId int64, playlist *playlistProto.Playlist) error {
	_, err := agent.PlaylistGrpc.Create(context.Background(), &playlistProto.UserIdPlaylistArg{
		UserId:   userId,
		Playlist: playlist,
	})
	return err
}
