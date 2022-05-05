package TrackGrpc

import (
	"context"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway/gatewayProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/trackProto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TrackGrpc struct {
	TrackRepo  *domain.TrackRepo
	ArtistRepo *domain.ArtistRepo
	AlbumRepo  *domain.AlbumRepo
	trackProto.UnimplementedTrackUseCaseServer
}

func MakeTrackGrpc(track domain.TrackRepo, artist domain.ArtistRepo, album domain.AlbumRepo) TrackGrpc {
	return TrackGrpc{
		TrackRepo:  &track,
		ArtistRepo: &artist,
		AlbumRepo:  &album,
	}
}

func (useCase TrackGrpc) GetAll(context.Context, *emptypb.Empty) (*trackProto.TracksResponse, error) {
	tracks, err := (*useCase.TrackRepo).GetAll()
	if err != nil {
		return nil, err
	}

	return &trackProto.TracksResponse{Tracks: tracks}, nil
}

func (useCase TrackGrpc) GetLastId(context.Context, *emptypb.Empty) (*gatewayProto.IntResponse, error) {
	id, err := (*useCase.TrackRepo).GetLastId()
	if err != nil {
		return nil, err
	}

	return &gatewayProto.IntResponse{Data: id}, nil
}

func (useCase TrackGrpc) Create(ctx context.Context, dom *trackProto.Track) (*emptypb.Empty, error) {
	err := (*useCase.TrackRepo).Create(dom)
	return &emptypb.Empty{}, err
}

func (useCase TrackGrpc) Update(ctx context.Context, dom *trackProto.Track) (*emptypb.Empty, error) {
	err := (*useCase.TrackRepo).Update(dom)
	return &emptypb.Empty{}, err
}

func (useCase TrackGrpc) Delete(ctx context.Context, data *gatewayProto.IdArg) (*emptypb.Empty, error) {
	err := (*useCase.TrackRepo).Delete(data.Id)
	return &emptypb.Empty{}, err
}

func (useCase TrackGrpc) GetById(ctx context.Context, data *gatewayProto.IdArg) (*trackProto.Track, error) {
	track, err := (*useCase.TrackRepo).SelectByID(data.Id)
	if err != nil {
		return nil, err
	}

	return track, nil
}

func (useCase TrackGrpc) GetPopular(context.Context, *emptypb.Empty) (*trackProto.TracksResponse, error) {
	tracks, err := (*useCase.TrackRepo).GetPopular()
	if err != nil {
		return nil, err
	}

	return &trackProto.TracksResponse{Tracks: tracks}, nil
}

func (useCase TrackGrpc) GetTracksFromAlbum(ctx context.Context, data *gatewayProto.IdArg) (*trackProto.TracksResponse, error) {
	tracks, err := (*useCase.TrackRepo).GetTracksFromAlbum(data.Id)
	if err != nil {
		return nil, err
	}

	return &trackProto.TracksResponse{Tracks: tracks}, nil
}

func (useCase TrackGrpc) GetPopularTracksFromArtist(ctx context.Context, data *gatewayProto.IdArg) (*trackProto.TracksResponse, error) {
	tracks, err := (*useCase.TrackRepo).GetPopularTracksFromArtist(data.Id)
	if err != nil {
		return nil, err
	}

	return &trackProto.TracksResponse{Tracks: tracks}, nil
}

func (useCase TrackGrpc) GetSize(context.Context, *emptypb.Empty) (*gatewayProto.IntResponse, error) {
	size, err := (*useCase.TrackRepo).GetSize()
	return &gatewayProto.IntResponse{Data: size}, err
}

func (useCase TrackGrpc) Like(ctx context.Context, data *gatewayProto.UserIdTrackIdArg) (*emptypb.Empty, error) {
	//track, err := (*useCase.TrackRepo).SelectByID(data.Id)
	//if err != nil {
	//	return nil, err
	//}

	if err := (*useCase.TrackRepo).Like(data.TrackId, data.UserId); err != nil {
		return nil, err
	}

	//if err := (*useCase.ArtistRepo).Like(track.ArtistId); err != nil {
	//	return nil, err
	//}
	//
	//if err := (*useCase.AlbumRepo).Like(track.AlbumId); err != nil {
	//	return nil, err
	//}

	return &emptypb.Empty{}, nil
}

func (useCase TrackGrpc) LikeCheckByUser(ctx context.Context, data *gatewayProto.UserIdTrackIdArg) (*trackProto.LikeCheckResponse, error) {
	liked, err := (*useCase.TrackRepo).LikeCheckByUser(data.TrackId, data.UserId)
	if err != nil {
		return nil, err
	}
	return &trackProto.LikeCheckResponse{Ok: liked}, nil
}

func (useCase TrackGrpc) Listen(ctx context.Context, data *gatewayProto.IdArg) (*emptypb.Empty, error) {
	track, err := (*useCase.TrackRepo).SelectByID(data.Id)
	if err != nil {
		return nil, err
	}

	if err := (*useCase.TrackRepo).Listen(track.Id); err != nil {
		return nil, err
	}

	if err := (*useCase.ArtistRepo).Listen(track.ArtistId); err != nil {
		return nil, err
	}

	if err := (*useCase.AlbumRepo).Listen(track.AlbumId); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (useCase TrackGrpc) SearchByTitle(ctx context.Context, title *gatewayProto.StringArg) (*trackProto.TracksResponse, error) {
	tracks, err := (*useCase.TrackRepo).SearchByTitle(title.Str)

	return &trackProto.TracksResponse{Tracks: tracks}, err
}

func (useCase TrackGrpc) GetFavorites(ctx context.Context, data *gatewayProto.IdArg) (*trackProto.TracksResponse, error) {
	tracks, err := (*useCase.TrackRepo).GetFavorites(data.Id)
	if err != nil {
		return nil, err
	}

	return &trackProto.TracksResponse{Tracks: tracks}, nil
}

func (useCase TrackGrpc) AddToFavorites(ctx context.Context, data *gatewayProto.UserIdTrackIdArg) (*emptypb.Empty, error) {
	if err := (*useCase.TrackRepo).AddToFavorites(data.TrackId, data.UserId); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (useCase TrackGrpc) RemoveFromFavorites(ctx context.Context, data *gatewayProto.UserIdTrackIdArg) (*emptypb.Empty, error) {
	if err := (*useCase.TrackRepo).RemoveFromFavorites(data.TrackId, data.UserId); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (useCase TrackGrpc) GetTracksFromPlaylist(ctx context.Context, data *gatewayProto.IdArg) (*trackProto.TracksResponse, error) {
	tracks, err := (*useCase.TrackRepo).GetTracksFromPlaylist(data.Id)
	if err != nil {
		return nil, err
	}

	return &trackProto.TracksResponse{Tracks: tracks}, nil
}
