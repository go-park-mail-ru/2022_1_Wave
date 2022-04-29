package TrackGrpc

import (
	"context"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway/gatewayProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/playlist/playlistProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/trackProto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PlaylistGrpc struct {
	PlaylistRepo *domain.PlaylistRepo
	TrackRepo    *domain.TrackRepo
	ArtistRepo   *domain.ArtistRepo
	AlbumRepo    *domain.AlbumRepo
	trackProto.UnimplementedTrackUseCaseServer
}

func MakePlaylistGrpc(track domain.TrackRepo, artist domain.ArtistRepo, album domain.AlbumRepo, playlist domain.PlaylistRepo) PlaylistGrpc {
	return PlaylistGrpc{
		TrackRepo:    &track,
		ArtistRepo:   &artist,
		AlbumRepo:    &album,
		PlaylistRepo: &playlist,
	}
}

func (useCase PlaylistGrpc) GetAll(ctx context.Context, arg *gatewayProto.IdArg) (*playlistProto.PlaylistsResponse, error) {
	playlists, err := (*useCase.PlaylistRepo).GetAll(arg.Id)

	if err != nil {
		return nil, err
	}

	return &playlistProto.PlaylistsResponse{Playlists: playlists}, nil
}

func (useCase PlaylistGrpc) GetLastId(ctx context.Context, arg *gatewayProto.IdArg) (*gatewayProto.IntResponse, error) {
	id, err := (*useCase.PlaylistRepo).GetLastId(arg.Id)
	if err != nil {
		return nil, err
	}

	return &gatewayProto.IntResponse{Data: id}, nil
}

func (useCase PlaylistGrpc) Create(ctx context.Context, input *playlistProto.UserIdPlaylistArg) (*emptypb.Empty, error) {
	err := (*useCase.PlaylistRepo).Create(input.UserId, input.Playlist)
	return &emptypb.Empty{}, err
}

func (useCase PlaylistGrpc) Update(ctx context.Context, input *playlistProto.UserIdPlaylistArg) (*emptypb.Empty, error) {
	err := (*useCase.PlaylistRepo).Update(input.UserId, input.Playlist)
	return &emptypb.Empty{}, err
}

func (useCase PlaylistGrpc) Delete(ctx context.Context, input *playlistProto.UserIdPlaylistIdArg) (*emptypb.Empty, error) {
	err := (*useCase.PlaylistRepo).Delete(input.UserId, input.PlaylistId)
	return &emptypb.Empty{}, err
}

func (useCase PlaylistGrpc) GetById(ctx context.Context, input *playlistProto.UserIdPlaylistIdArg) (*playlistProto.Playlist, error) {
	playlist, err := (*useCase.PlaylistRepo).SelectByID(input.UserId, input.PlaylistId)
	if err != nil {
		return nil, err
	}

	return playlist, nil
}

func (useCase PlaylistGrpc) GetSize(ctx context.Context, userId *gatewayProto.IdArg) (*gatewayProto.IntResponse, error) {
	size, err := (*useCase.PlaylistRepo).GetSize(userId.Id)
	return &gatewayProto.IntResponse{Data: size}, err
}
