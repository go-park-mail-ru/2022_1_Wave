package TrackGrpc

import (
	"context"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	Gateway "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway/gatewayProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/playlist/playlistProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/trackProto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcAgent struct {
	PlaylistGrpc playlistProto.PlaylistUseCaseClient
}

func MakeAgent(gRPC playlistProto.PlaylistUseCaseClient) GrpcAgent {
	return GrpcAgent{PlaylistGrpc: gRPC}
}

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

func (useCase PlaylistGrpc) CastToDTO(playlist *playlistProto.Playlist) (*playlistProto.PlaylistDataTransfer, error) {
	tracks := make([]*trackProto.TrackDataTransfer, len(playlist.TracksId))
	for idx, trackId := range playlist.TracksId {
		track, err := (*useCase.TrackRepo).SelectByID(trackId)
		if err != nil {
			return nil, err
		}

		artist, err := (*useCase.ArtistRepo).SelectByID(track.ArtistId)
		if err != nil {
			return nil, err
		}

		trackDto, err := Gateway.CastTrackToDtoWithoutArtistName(track)
		if err != nil {
			return nil, err
		}
		trackDto.Artist = artist.Name

		tracks[idx] = trackDto
	}

	playlistDto := &playlistProto.PlaylistDataTransfer{
		Id:     0,
		Title:  "",
		Tracks: tracks,
	}

	return playlistDto, nil
}

func (useCase PlaylistGrpc) GetAll(ctx context.Context, arg *gatewayProto.IdArg) (*playlistProto.PlaylistsResponse, error) {
	playlists, err := (*useCase.PlaylistRepo).GetAll(arg.Id)

	if err != nil {
		return nil, err
	}

	dto := make([]*playlistProto.PlaylistDataTransfer, len(playlists))

	for idx, obj := range playlists {
		result, err := useCase.CastToDTO(obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = result
	}

	return &playlistProto.PlaylistsResponse{Playlists: dto}, nil
}

func (agent GrpcAgent) GetAll(arg *gatewayProto.IdArg) (*playlistProto.PlaylistsResponse, error) {
	return agent.PlaylistGrpc.GetAll(context.Background(), arg)
}

func (useCase PlaylistGrpc) GetLastId(ctx context.Context, arg *gatewayProto.IdArg) (*gatewayProto.IntResponse, error) {
	id, err := (*useCase.PlaylistRepo).GetLastId(arg.Id)
	if err != nil {
		return nil, err
	}

	return &gatewayProto.IntResponse{Data: id}, nil
}

func (agent GrpcAgent) GetLastId(arg *gatewayProto.IdArg) (*gatewayProto.IntResponse, error) {
	return agent.PlaylistGrpc.GetLastId(context.Background(), arg)
}

func (useCase PlaylistGrpc) Create(ctx context.Context, input *playlistProto.UserIdPlaylistArg) (*emptypb.Empty, error) {
	err := (*useCase.PlaylistRepo).Create(input.UserId, input.Playlist)
	return &emptypb.Empty{}, err
}

func (agent GrpcAgent) Create(input *playlistProto.UserIdPlaylistArg) error {
	_, err := agent.PlaylistGrpc.Create(context.Background(), input)
	return err
}

func (useCase PlaylistGrpc) Update(ctx context.Context, input *playlistProto.UserIdPlaylistArg) (*emptypb.Empty, error) {
	err := (*useCase.PlaylistRepo).Update(input.UserId, input.Playlist)
	return &emptypb.Empty{}, err
}

func (agent GrpcAgent) Update(input *playlistProto.UserIdPlaylistArg) error {
	_, err := agent.PlaylistGrpc.Update(context.Background(), input)
	return err
}

func (useCase PlaylistGrpc) Delete(ctx context.Context, input *playlistProto.UserIdPlaylistIdArg) (*emptypb.Empty, error) {
	err := (*useCase.PlaylistRepo).Delete(input.UserId, input.PlaylistId)
	return &emptypb.Empty{}, err
}

func (agent GrpcAgent) Delete(input *playlistProto.UserIdPlaylistIdArg) error {
	_, err := agent.PlaylistGrpc.Delete(context.Background(), input)
	return err
}

func (useCase PlaylistGrpc) GetById(ctx context.Context, input *playlistProto.UserIdPlaylistIdArg) (*playlistProto.PlaylistDataTransfer, error) {
	playlist, err := (*useCase.PlaylistRepo).SelectByID(input.UserId, input.PlaylistId)
	if err != nil {
		return nil, err
	}
	dto, err := useCase.CastToDTO(playlist)
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (agent GrpcAgent) GetById(input *playlistProto.UserIdPlaylistIdArg) (*playlistProto.PlaylistDataTransfer, error) {
	return agent.PlaylistGrpc.GetById(context.Background(), input)
}

func (useCase PlaylistGrpc) GetSize(ctx context.Context, userId *gatewayProto.IdArg) (*gatewayProto.IntResponse, error) {
	size, err := (*useCase.PlaylistRepo).GetSize(userId.Id)
	return &gatewayProto.IntResponse{Data: size}, err
}

func (agent GrpcAgent) GetSize(userId *gatewayProto.IdArg) (*gatewayProto.IntResponse, error) {
	return agent.PlaylistGrpc.GetSize(context.Background(), userId)
}
