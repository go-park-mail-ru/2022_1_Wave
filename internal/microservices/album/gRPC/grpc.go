package AlbumGrpc

import (
	"context"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/album/albumProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway/gatewayProto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AlbumGrpc struct {
	TrackRepo      *domain.TrackRepo
	ArtistRepo     *domain.ArtistRepo
	AlbumRepo      *domain.AlbumRepo
	AlbumCoverRepo *domain.AlbumCoverRepo
	albumProto.UnimplementedAlbumUseCaseServer
}

func MakeAlbumGrpc(track domain.TrackRepo,
	artist domain.ArtistRepo,
	album domain.AlbumRepo,
	albumCover domain.AlbumCoverRepo) AlbumGrpc {
	return AlbumGrpc{
		TrackRepo:      &track,
		ArtistRepo:     &artist,
		AlbumRepo:      &album,
		AlbumCoverRepo: &albumCover,
	}
}

func (useCase AlbumGrpc) GetAll(context.Context, *emptypb.Empty) (*albumProto.AlbumsResponse, error) {
	albums, err := (*useCase.AlbumRepo).GetAll()
	if err != nil {
		return nil, err
	}
	return &albumProto.AlbumsResponse{Albums: albums}, nil
}

func (useCase AlbumGrpc) GetAllCovers(context.Context, *emptypb.Empty) (*albumProto.AlbumsCoverResponse, error) {
	covers, err := (*useCase.AlbumCoverRepo).GetAll()
	if err != nil {
		return nil, err
	}
	return &albumProto.AlbumsCoverResponse{Covers: covers}, nil
}

func (useCase AlbumGrpc) GetLastId(context.Context, *emptypb.Empty) (*gatewayProto.IntResponse, error) {
	id, err := (*useCase.AlbumRepo).GetLastId()
	if err != nil {
		return nil, err
	}

	return &gatewayProto.IntResponse{Data: id}, nil
}

func (useCase AlbumGrpc) GetLastCoverId(context.Context, *emptypb.Empty) (*gatewayProto.IntResponse, error) {
	id, err := (*useCase.AlbumCoverRepo).GetLastId()
	if err != nil {
		return nil, err
	}

	return &gatewayProto.IntResponse{Data: id}, nil
}

func (useCase AlbumGrpc) Create(ctx context.Context, album *albumProto.Album) (*emptypb.Empty, error) {
	err := (*useCase.AlbumRepo).Create(album)
	return &emptypb.Empty{}, err
}

func (useCase AlbumGrpc) CreateCover(ctx context.Context, cover *albumProto.AlbumCover) (*emptypb.Empty, error) {
	err := (*useCase.AlbumCoverRepo).Create(cover)
	return &emptypb.Empty{}, err
}

func (useCase AlbumGrpc) Update(ctx context.Context, album *albumProto.Album) (*emptypb.Empty, error) {
	err := (*useCase.AlbumRepo).Update(album)
	return &emptypb.Empty{}, err
}

func (useCase AlbumGrpc) UpdateCover(ctx context.Context, cover *albumProto.AlbumCover) (*emptypb.Empty, error) {
	err := (*useCase.AlbumCoverRepo).Update(cover)
	return &emptypb.Empty{}, err
}

func (useCase AlbumGrpc) Delete(ctx context.Context, data *gatewayProto.IdArg) (*emptypb.Empty, error) {
	err := (*useCase.AlbumRepo).Delete(data.Id)
	return &emptypb.Empty{}, err
}

func (useCase AlbumGrpc) DeleteCover(ctx context.Context, data *gatewayProto.IdArg) (*emptypb.Empty, error) {
	err := (*useCase.AlbumCoverRepo).Delete(data.Id)
	return &emptypb.Empty{}, err
}

func (useCase AlbumGrpc) GetById(ctx context.Context, data *gatewayProto.IdArg) (*albumProto.Album, error) {
	album, err := (*useCase.AlbumRepo).SelectByID(data.Id)
	if err != nil {
		return nil, err
	}
	return album, nil
}

func (useCase AlbumGrpc) GetCoverById(ctx context.Context, data *gatewayProto.IdArg) (*albumProto.AlbumCover, error) {
	cover, err := (*useCase.AlbumCoverRepo).SelectByID(data.Id)
	if err != nil {
		return nil, err
	}

	return cover, nil
}

func (useCase AlbumGrpc) GetPopular(context.Context, *emptypb.Empty) (*albumProto.AlbumsResponse, error) {
	albums, err := (*useCase.AlbumRepo).GetPopular()
	if err != nil {
		return nil, err
	}

	return &albumProto.AlbumsResponse{Albums: albums}, nil
}

func (useCase AlbumGrpc) GetAlbumsFromArtist(ctx context.Context, artistData *gatewayProto.IdArg) (*albumProto.AlbumsResponse, error) {
	albums, err := (*useCase.AlbumRepo).GetAlbumsFromArtist(artistData.Id)
	if err != nil {
		return nil, err
	}

	return &albumProto.AlbumsResponse{Albums: albums}, nil
}

func (useCase AlbumGrpc) GetSize(context.Context, *emptypb.Empty) (*gatewayProto.IntResponse, error) {
	size, err := (*useCase.AlbumRepo).GetSize()
	return &gatewayProto.IntResponse{Data: size}, err
}

func (useCase AlbumGrpc) SearchByTitle(ctx context.Context, title *gatewayProto.StringArg) (*albumProto.AlbumsResponse, error) {
	albums, err := (*useCase.AlbumRepo).SearchByTitle(title.Str)

	return &albumProto.AlbumsResponse{Albums: albums}, err
}

func (useCase AlbumGrpc) GetFavorites(ctx context.Context, data *gatewayProto.IdArg) (*albumProto.AlbumsResponse, error) {
	albums, err := (*useCase.AlbumRepo).GetFavorites(data.Id)
	if err != nil {
		return nil, err
	}

	return &albumProto.AlbumsResponse{Albums: albums}, nil
}

func (useCase AlbumGrpc) AddToFavorites(ctx context.Context, data *gatewayProto.UserIdAlbumIdArg) (*emptypb.Empty, error) {
	if err := (*useCase.AlbumRepo).AddToFavorites(data.AlbumId, data.UserId); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (useCase AlbumGrpc) RemoveFromFavorites(ctx context.Context, data *gatewayProto.UserIdAlbumIdArg) (*emptypb.Empty, error) {
	if err := (*useCase.AlbumRepo).RemoveFromFavorites(data.AlbumId, data.UserId); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (useCase AlbumGrpc) Like(ctx context.Context, data *gatewayProto.UserIdAlbumIdArg) (*emptypb.Empty, error) {
	if err := (*useCase.AlbumRepo).Like(data.AlbumId, data.UserId); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (useCase AlbumGrpc) LikeCheckByUser(ctx context.Context, data *gatewayProto.UserIdAlbumIdArg) (*gatewayProto.LikeCheckResponse, error) {
	liked, err := (*useCase.AlbumRepo).LikeCheckByUser(data.AlbumId, data.UserId)
	if err != nil {
		return nil, err
	}
	return &gatewayProto.LikeCheckResponse{Ok: liked}, nil
}

func (useCase AlbumGrpc) GetPopularAlbumOfWeekTop20(context.Context, *emptypb.Empty) (*albumProto.AlbumsResponse, error) {
	albums, err := (*useCase.AlbumRepo).GetPopularAlbumOfWeekTop20()
	if err != nil {
		return nil, err
	}

	return &albumProto.AlbumsResponse{Albums: albums}, nil
}
