package AlbumUseCase

import (
	"context"
	"fmt"
	AlbumPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/app/album/repository"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/album/albumProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/common/commonProto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AlbumUseCase struct {
	TrackRepo      *domain.TrackRepo
	ArtistRepo     *domain.ArtistRepo
	AlbumRepo      *domain.AlbumRepo
	AlbumCoverRepo *domain.AlbumCoverRepo
	albumProto.UnimplementedAlbumUseCaseServer
}

func MakeAlbumService(track domain.TrackRepo,
	artist domain.ArtistRepo,
	album domain.AlbumRepo,
	albumCover domain.AlbumCoverRepo) AlbumUseCase {
	return AlbumUseCase{
		TrackRepo:      &track,
		ArtistRepo:     &artist,
		AlbumRepo:      &album,
		AlbumCoverRepo: &albumCover,
	}
}

func (useCase AlbumUseCase) CastToDTO(album *albumProto.Album) (*albumProto.AlbumDataTransfer, error) {
	artist, err := (*useCase.ArtistRepo).SelectByID(album.ArtistId)
	if err != nil {
		return nil, err
	}
	return AlbumPostgres.GetFullAlbumByArtist(*useCase.TrackRepo, album, artist)
}

func (useCase AlbumUseCase) CastCoverToDTO(cover *albumProto.AlbumCover) (*albumProto.AlbumCoverDataTransfer, error) {
	return &albumProto.AlbumCoverDataTransfer{
		Quote:  cover.Quote,
		IsDark: cover.IsDark,
	}, nil
}

func (useCase AlbumUseCase) GetAll(context.Context, *emptypb.Empty) (*albumProto.AlbumsResponse, error) {
	albums, err := (*useCase.AlbumRepo).GetAll()
	if err != nil {
		return nil, err
	}

	dto := make([]*albumProto.AlbumDataTransfer, len(albums))

	for i := 0; i < len(albums); i++ {
		data, err := useCase.CastToDTO(albums[i])
		if err != nil {
			return nil, err
		}
		dto[i] = data
	}

	return &albumProto.AlbumsResponse{Albums: dto}, nil
}

func (useCase AlbumUseCase) GetAllCovers(context.Context, *emptypb.Empty) (*albumProto.AlbumsCoverResponse, error) {
	covers, err := (*useCase.AlbumCoverRepo).GetAll()
	if err != nil {
		return nil, err
	}

	dto := make([]*albumProto.AlbumCoverDataTransfer, len(covers))

	for idx, obj := range covers {
		data, err := useCase.CastCoverToDTO(obj)
		fmt.Println("casted=", data.IsDark)
		if err != nil {
			return nil, err
		}
		dto[idx] = data
	}

	return &albumProto.AlbumsCoverResponse{Albums: dto}, nil
}

func (useCase AlbumUseCase) GetLastId(context.Context, *emptypb.Empty) (*commonProto.IntResponse, error) {
	id, err := (*useCase.AlbumRepo).GetLastId()
	if err != nil {
		return nil, err
	}

	return &commonProto.IntResponse{Data: id}, nil
}

func (useCase AlbumUseCase) GetLastCoverId(context.Context, *emptypb.Empty) (*commonProto.IntResponse, error) {
	id, err := (*useCase.AlbumCoverRepo).GetLastId()
	if err != nil {
		return nil, err
	}

	return &commonProto.IntResponse{Data: id}, nil
}

func (useCase AlbumUseCase) Create(ctx context.Context, album *albumProto.Album) (*emptypb.Empty, error) {
	err := (*useCase.AlbumRepo).Create(album)
	return &emptypb.Empty{}, err
}

func (useCase AlbumUseCase) CreateCover(ctx context.Context, cover *albumProto.AlbumCover) (*emptypb.Empty, error) {
	err := (*useCase.AlbumCoverRepo).Create(cover)
	return &emptypb.Empty{}, err
}

func (useCase AlbumUseCase) Update(ctx context.Context, album *albumProto.Album) (*emptypb.Empty, error) {
	err := (*useCase.AlbumRepo).Update(album)
	return &emptypb.Empty{}, err
}

func (useCase AlbumUseCase) UpdateCover(ctx context.Context, cover *albumProto.AlbumCover) (*emptypb.Empty, error) {
	err := (*useCase.AlbumCoverRepo).Update(cover)
	return &emptypb.Empty{}, err
}

func (useCase AlbumUseCase) Delete(ctx context.Context, data *commonProto.IdArg) (*emptypb.Empty, error) {
	err := (*useCase.AlbumRepo).Delete(data.Id)
	return &emptypb.Empty{}, err
}

func (useCase AlbumUseCase) DeleteCover(ctx context.Context, data *commonProto.IdArg) (*emptypb.Empty, error) {
	err := (*useCase.AlbumCoverRepo).Delete(data.Id)
	return &emptypb.Empty{}, err
}

func (useCase AlbumUseCase) GetById(ctx context.Context, data *commonProto.IdArg) (*albumProto.AlbumDataTransfer, error) {
	album, err := (*useCase.AlbumRepo).SelectByID(data.Id)
	if err != nil {
		return nil, err
	}
	dto, err := useCase.CastToDTO(album)
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (useCase AlbumUseCase) GetCoverById(ctx context.Context, data *commonProto.IdArg) (*albumProto.AlbumCoverDataTransfer, error) {
	cover, err := (*useCase.AlbumCoverRepo).SelectByID(data.Id)
	if err != nil {
		return nil, err
	}
	dto, err := useCase.CastCoverToDTO(cover)
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (useCase AlbumUseCase) GetPopular(context.Context, *emptypb.Empty) (*albumProto.AlbumsResponse, error) {
	albums, err := (*useCase.AlbumRepo).GetPopular()
	if err != nil {
		return nil, err
	}

	dto := make([]*albumProto.AlbumDataTransfer, len(albums))

	for i := 0; i < len(albums); i++ {
		data, err := useCase.CastToDTO(albums[i])
		if err != nil {
			return nil, err
		}
		dto[i] = data
	}

	return &albumProto.AlbumsResponse{Albums: dto}, nil
}

func (useCase AlbumUseCase) GetAlbumsFromArtist(ctx context.Context, artistData *commonProto.IdArg) (*albumProto.AlbumsResponse, error) {
	albums, err := (*useCase.AlbumRepo).GetAlbumsFromArtist(artistData.Id)
	if err != nil {
		return nil, err
	}

	dto := make([]*albumProto.AlbumDataTransfer, len(albums))

	for idx, album := range albums {
		data, err := useCase.CastToDTO(album)
		if err != nil {
			return nil, err
		}
		dto[idx] = data
	}

	return &albumProto.AlbumsResponse{Albums: dto}, nil
}

func (useCase AlbumUseCase) GetSize(context.Context, *emptypb.Empty) (*commonProto.IntResponse, error) {
	size, err := (*useCase.AlbumRepo).GetSize()
	return &commonProto.IntResponse{Data: size}, err
}
