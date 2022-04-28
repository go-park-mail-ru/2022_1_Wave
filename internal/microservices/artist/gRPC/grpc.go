package ArtistGrpc

import (
	"context"
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/album/albumProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/artist/artistProto"
	Gateway "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway"
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

func (useCase ArtistGrpc) CastToDTO(artist *artistProto.Artist) (*artistProto.ArtistDataTransfer, error) {
	coverPath, err := Gateway.PathToArtistCover(artist, internal.PngFormat)
	if err != nil {
		return nil, err
	}

	albums, err := (*useCase.AlbumRepo).GetAlbumsFromArtist(artist.Id)
	if err != nil {
		return nil, err
	}

	albumsDto := make([]*albumProto.AlbumDataTransfer, len(albums))

	for idx, album := range albums {
		albumDto, err := Gateway.GetFullAlbumByArtist(*useCase.TrackRepo, album, artist)
		if err != nil {
			return nil, err
		}
		albumsDto[idx] = albumDto
	}

	return &artistProto.ArtistDataTransfer{
		Id:     artist.Id,
		Name:   artist.Name,
		Cover:  coverPath,
		Likes:  artist.CountLikes,
		Albums: albumsDto,
	}, nil
}

func (useCase ArtistGrpc) GetAll(context.Context, *emptypb.Empty) (*artistProto.ArtistsResponse, error) {
	artists, err := (*useCase.ArtistRepo).GetAll()
	if err != nil {
		return nil, err
	}

	dto := make([]*artistProto.ArtistDataTransfer, len(artists))

	for i := 0; i < len(artists); i++ {
		data, err := useCase.CastToDTO(artists[i])
		if err != nil {
			return nil, err
		}
		dto[i] = data
	}

	return &artistProto.ArtistsResponse{Artists: dto}, nil
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

func (useCase ArtistGrpc) GetById(ctx context.Context, data *gatewayProto.IdArg) (*artistProto.ArtistDataTransfer, error) {
	artist, err := (*useCase.ArtistRepo).SelectByID(data.Id)
	if err != nil {
		return nil, err
	}
	dto, err := useCase.CastToDTO(artist)
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (useCase ArtistGrpc) GetPopular(context.Context, *emptypb.Empty) (*artistProto.ArtistsResponse, error) {
	artists, err := (*useCase.ArtistRepo).GetPopular()
	if err != nil {
		return nil, err
	}

	dto := make([]*artistProto.ArtistDataTransfer, len(artists))

	for i := 0; i < len(artists); i++ {
		data, err := useCase.CastToDTO(artists[i])
		if err != nil {
			return nil, err
		}
		dto[i] = data
	}

	return &artistProto.ArtistsResponse{Artists: dto}, nil
}

func (useCase ArtistGrpc) GetSize(context.Context, *emptypb.Empty) (*gatewayProto.IntResponse, error) {
	size, err := (*useCase.ArtistRepo).GetSize()
	return &gatewayProto.IntResponse{Data: size}, err
}

func (useCase ArtistGrpc) SearchByName(ctx context.Context, title *gatewayProto.StringArg) (*artistProto.ArtistsResponse, error) {
	artists, err := (*useCase.ArtistRepo).SearchByName(title.Str)

	dto := make([]*artistProto.ArtistDataTransfer, len(artists))

	for idx, artist := range artists {
		data, err := useCase.CastToDTO(artist)
		if err != nil {
			return nil, err
		}
		dto[idx] = data
	}

	return &artistProto.ArtistsResponse{Artists: dto}, err
}

func (useCase ArtistGrpc) GetFavorites(ctx context.Context, data *gatewayProto.IdArg) (*artistProto.ArtistsResponse, error) {
	artists, err := (*useCase.ArtistRepo).GetFavorites(data.Id)
	if err != nil {
		return nil, err
	}

	dto := make([]*artistProto.ArtistDataTransfer, len(artists))

	for i := 0; i < len(artists); i++ {
		data, err := useCase.CastToDTO(artists[i])
		if err != nil {
			return nil, err
		}
		dto[i] = data
	}

	return &artistProto.ArtistsResponse{Artists: dto}, nil
}

func (useCase ArtistGrpc) AddToFavorites(ctx context.Context, data *gatewayProto.UserIdArtistIdArg) (*emptypb.Empty, error) {
	if err := (*useCase.ArtistRepo).AddToFavorites(data.ArtistId, data.UserId); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (useCase ArtistGrpc) RemoveFromFavorites(ctx context.Context, data *gatewayProto.UserIdArtistIdArg) (*emptypb.Empty, error) {
	if err := (*useCase.TrackRepo).RemoveFromFavorites(data.ArtistId, data.UserId); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
