package PlaylistGrpcAgent

import (
	"context"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway/gatewayProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/playlist/playlistProto"
	"github.com/golang/protobuf/ptypes/empty"
)

type GrpcAgent struct {
	PlaylistGrpc playlistProto.PlaylistUseCaseClient
}

func MakeAgent(gRPC playlistProto.PlaylistUseCaseClient) GrpcAgent {
	return GrpcAgent{
		PlaylistGrpc: gRPC,
	}
}

func (agent GrpcAgent) GetAllOfCurrentUser(userId int64) ([]*playlistProto.Playlist, error) {
	data, err := agent.PlaylistGrpc.GetAllOfCurrentUser(context.Background(), &gatewayProto.IdArg{Id: userId})
	return data.Playlists, err
}

func (agent GrpcAgent) GetAll() ([]*playlistProto.Playlist, error) {
	data, err := agent.PlaylistGrpc.GetAll(context.Background(), &empty.Empty{})
	return data.Playlists, err
}

func (agent GrpcAgent) GetLastIdOfCurrentUser(userId int64) (int64, error) {
	id, err := agent.PlaylistGrpc.GetLastIdOfCurrentUser(context.Background(), &gatewayProto.IdArg{Id: userId})
	return id.Data, err
}

func (agent GrpcAgent) GetLastId() (int64, error) {
	id, err := agent.PlaylistGrpc.GetLastId(context.Background(), &empty.Empty{})
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

func (agent GrpcAgent) GetByIdOfCurrentUser(userId int64, playlistId int64) (*playlistProto.Playlist, error) {
	return agent.PlaylistGrpc.GetByIdOfCurrentUser(context.Background(), &playlistProto.UserIdPlaylistIdArg{
		UserId:     userId,
		PlaylistId: playlistId,
	})
}

func (agent GrpcAgent) GetById(playlistId int64) (*playlistProto.Playlist, error) {
	return agent.PlaylistGrpc.GetById(context.Background(), &gatewayProto.IdArg{Id: playlistId})
}

func (agent GrpcAgent) GetSizeOfCurrentUser(userId int64) (int64, error) {
	size, err := agent.PlaylistGrpc.GetSizeOfCurrentUser(context.Background(), &gatewayProto.IdArg{Id: userId})
	return size.Data, err
}

func (agent GrpcAgent) GetSize() (int64, error) {
	size, err := agent.PlaylistGrpc.GetSize(context.Background(), &empty.Empty{})
	return size.Data, err
}

func (agent GrpcAgent) Create(userId int64, playlist *playlistProto.Playlist) error {
	_, err := agent.PlaylistGrpc.Create(context.Background(), &playlistProto.UserIdPlaylistArg{
		UserId:   userId,
		Playlist: playlist,
	})
	return err
}

func (agent GrpcAgent) AddToPlaylist(userId int64, playlistId int64, trackId int64) error {
	_, err := agent.PlaylistGrpc.AddToPlaylist(context.Background(), &playlistProto.UserIdPlaylistIdTracksArg{
		UserId:     userId,
		PlaylistId: playlistId,
		TrackId:    trackId,
	})
	return err
}

func (agent GrpcAgent) RemoveFromPlaylist(userId int64, playlistId int64, trackId int64) error {
	_, err := agent.PlaylistGrpc.RemoveFromPlaylist(context.Background(), &playlistProto.UserIdPlaylistIdTracksArg{
		UserId:     userId,
		PlaylistId: playlistId,
		TrackId:    trackId,
	})
	return err
}
