package ArtistGrpc

import (
	"context"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/artist/artistProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway/gatewayProto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ArtistGrpc struct {
	ArtistRepo *domain.ArtistRepo
	AlbumRepo  *domain.AlbumRepo
	TrackRepo  *domain.TrackRepo
	artistProto.UnimplementedArtistUseCaseServer
}

func MakeArtistGrpc(artistRepo domain.ArtistRepo, albumRepo domain.AlbumRepo, trackRepo domain.TrackRepo) ArtistGrpc {
	return ArtistGrpc{
		ArtistRepo: &artistRepo,
		AlbumRepo:  &albumRepo,
		TrackRepo:  &trackRepo}
}

func (useCase ArtistGrpc) GetAll(context.Context, *emptypb.Empty) (*artistProto.ArtistsResponse, error) {
	artists, err := (*useCase.ArtistRepo).GetAll()
	if err != nil {
		return nil, err
	}

	return &artistProto.ArtistsResponse{Artists: artists}, nil
}

func (useCase ArtistGrpc) GetLastId(context.Context, *emptypb.Empty) (*gatewayProto.IntResponse, error) {
	id, err := (*useCase.ArtistRepo).GetLastId()
	if err != nil {
		return nil, err
	}

	return &gatewayProto.IntResponse{Data: id}, nil
}

func (useCase ArtistGrpc) Create(ctx context.Context, artist *artistProto.Artist) (*emptypb.Empty, error) {
	err := (*useCase.ArtistRepo).Create(artist)
	return &emptypb.Empty{}, err
}

func (useCase ArtistGrpc) Update(ctx context.Context, artist *artistProto.Artist) (*emptypb.Empty, error) {
	err := (*useCase.ArtistRepo).Update(artist)
	return &emptypb.Empty{}, err
}

func (useCase ArtistGrpc) Delete(ctx context.Context, data *gatewayProto.IdArg) (*emptypb.Empty, error) {
	err := (*useCase.ArtistRepo).Delete(data.Id)
	return &emptypb.Empty{}, err
}

func (useCase ArtistGrpc) GetById(ctx context.Context, data *gatewayProto.IdArg) (*artistProto.Artist, error) {
	artist, err := (*useCase.ArtistRepo).SelectByID(data.Id)
	if err != nil {
		return nil, err
	}
	return artist, nil
}

func (useCase ArtistGrpc) GetPopular(context.Context, *emptypb.Empty) (*artistProto.ArtistsResponse, error) {
	artists, err := (*useCase.ArtistRepo).GetPopular()
	if err != nil {
		return nil, err
	}

	return &artistProto.ArtistsResponse{Artists: artists}, nil
}

func (useCase ArtistGrpc) GetSize(context.Context, *emptypb.Empty) (*gatewayProto.IntResponse, error) {
	size, err := (*useCase.ArtistRepo).GetSize()
	return &gatewayProto.IntResponse{Data: size}, err
}

func (useCase ArtistGrpc) SearchByName(ctx context.Context, title *gatewayProto.StringArg) (*artistProto.ArtistsResponse, error) {
	artists, err := (*useCase.ArtistRepo).SearchByName(title.Str)

	return &artistProto.ArtistsResponse{Artists: artists}, err
}

func (useCase ArtistGrpc) GetFavorites(ctx context.Context, data *gatewayProto.IdArg) (*artistProto.ArtistsResponse, error) {
	artists, err := (*useCase.ArtistRepo).GetFavorites(data.Id)
	if err != nil {
		return nil, err
	}

	return &artistProto.ArtistsResponse{Artists: artists}, nil
}

func (useCase ArtistGrpc) AddToFavorites(ctx context.Context, data *gatewayProto.UserIdArtistIdArg) (*emptypb.Empty, error) {
	if err := (*useCase.ArtistRepo).AddToFavorites(data.ArtistId, data.UserId); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (useCase ArtistGrpc) RemoveFromFavorites(ctx context.Context, data *gatewayProto.UserIdArtistIdArg) (*emptypb.Empty, error) {
	if err := (*useCase.ArtistRepo).RemoveFromFavorites(data.ArtistId, data.UserId); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (useCase ArtistGrpc) Like(ctx context.Context, data *gatewayProto.UserIdArtistIdArg) (*emptypb.Empty, error) {
	if err := (*useCase.ArtistRepo).Like(data.ArtistId, data.UserId); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (useCase ArtistGrpc) LikeCheckByUser(ctx context.Context, data *gatewayProto.UserIdArtistIdArg) (*gatewayProto.LikeCheckResponse, error) {
	liked, err := (*useCase.ArtistRepo).LikeCheckByUser(data.ArtistId, data.UserId)
	if err != nil {
		return nil, err
	}
	return &gatewayProto.LikeCheckResponse{Ok: liked}, nil
}
