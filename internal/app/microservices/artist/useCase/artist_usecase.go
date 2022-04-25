package ArtistUseCase

import (
	"context"
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	AlbumPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/app/album/repository"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/album/albumProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/artist/artistProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/common/commonProto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ArtistUseCase struct {
	ArtistRepo *domain.ArtistRepo
	AlbumRepo  *domain.AlbumRepo
	TrackRepo  *domain.TrackRepo
	artistProto.UnimplementedArtistUseCaseServer
}

func MakeArtistService(artistRepo domain.ArtistRepo, albumRepo domain.AlbumRepo, trackRepo domain.TrackRepo) ArtistUseCase {
	return ArtistUseCase{
		ArtistRepo: &artistRepo,
		AlbumRepo:  &albumRepo,
		TrackRepo:  &trackRepo}
}

func (useCase ArtistUseCase) CastToDTO(artist *artistProto.Artist) (*artistProto.ArtistDataTransfer, error) {
	coverPath, err := AlbumPostgres.PathToArtistCover(artist, internal.PngFormat)
	if err != nil {
		return nil, err
	}

	albums, err := (*useCase.AlbumRepo).GetAlbumsFromArtist(artist.Id)
	if err != nil {
		return nil, err
	}

	albumsDto := make([]*albumProto.AlbumDataTransfer, len(albums))

	for idx, album := range albums {
		albumDto, err := AlbumPostgres.GetFullAlbumByArtist(*useCase.TrackRepo, album, artist)
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

func (useCase ArtistUseCase) GetAll(context.Context, *emptypb.Empty) (*artistProto.ArtistsResponse, error) {
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

func (useCase ArtistUseCase) GetLastId(context.Context, *emptypb.Empty) (*commonProto.IntResponse, error) {
	id, err := (*useCase.ArtistRepo).GetLastId()
	if err != nil {
		return nil, err
	}

	return &commonProto.IntResponse{Data: id}, nil
}

func (useCase ArtistUseCase) Create(ctx context.Context, artist *artistProto.Artist) (*emptypb.Empty, error) {
	err := (*useCase.ArtistRepo).Create(artist)
	return &emptypb.Empty{}, err
}

func (useCase ArtistUseCase) Update(ctx context.Context, artist *artistProto.Artist) (*emptypb.Empty, error) {
	err := (*useCase.ArtistRepo).Update(artist)
	return &emptypb.Empty{}, err
}

func (useCase ArtistUseCase) Delete(ctx context.Context, data *commonProto.IdArg) (*emptypb.Empty, error) {
	err := (*useCase.ArtistRepo).Delete(data.Id)
	return &emptypb.Empty{}, err
}

func (useCase ArtistUseCase) GetById(ctx context.Context, data *commonProto.IdArg) (*artistProto.ArtistDataTransfer, error) {
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

func (useCase ArtistUseCase) GetPopular(context.Context, *emptypb.Empty) (*artistProto.ArtistsResponse, error) {
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

func (useCase ArtistUseCase) GetSize(context.Context, *emptypb.Empty) (*commonProto.IntResponse, error) {
	size, err := (*useCase.ArtistRepo).GetSize()
	return &commonProto.IntResponse{Data: size}, err
}
