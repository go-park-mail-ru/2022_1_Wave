package AlbumGrpc

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/album/albumProto"
	Gateway "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway/gatewayProto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcAgent struct {
	AlbumGrpc albumProto.AlbumUseCaseClient
}

func MakeAgent(gRPC albumProto.AlbumUseCaseClient) GrpcAgent {
	return GrpcAgent{AlbumGrpc: gRPC}
}

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

func (useCase AlbumGrpc) CastToDTO(album *albumProto.Album) (*albumProto.AlbumDataTransfer, error) {
	artist, err := (*useCase.ArtistRepo).SelectByID(album.ArtistId)
	if err != nil {
		return nil, err
	}
	return Gateway.GetFullAlbumByArtist(*useCase.TrackRepo, album, artist)
}

func (useCase AlbumGrpc) CastCoverToDTO(cover *albumProto.AlbumCover) (*albumProto.AlbumCoverDataTransfer, error) {
	return &albumProto.AlbumCoverDataTransfer{
		Quote:  cover.Quote,
		IsDark: cover.IsDark,
	}, nil
}

func (useCase AlbumGrpc) GetAll(context.Context, *emptypb.Empty) (*albumProto.AlbumsResponse, error) {
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

func (agent GrpcAgent) GetAll() (*albumProto.AlbumsResponse, error) {
	return agent.AlbumGrpc.GetAll(context.Background(), &emptypb.Empty{})
}

func (useCase AlbumGrpc) GetAllCovers(context.Context, *emptypb.Empty) (*albumProto.AlbumsCoverResponse, error) {
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

func (agent GrpcAgent) GetAllCovers() (*albumProto.AlbumsCoverResponse, error) {
	return agent.AlbumGrpc.GetAllCovers(context.Background(), &emptypb.Empty{})
}

func (useCase AlbumGrpc) GetLastId(context.Context, *emptypb.Empty) (*gatewayProto.IntResponse, error) {
	id, err := (*useCase.AlbumRepo).GetLastId()
	if err != nil {
		return nil, err
	}

	return &gatewayProto.IntResponse{Data: id}, nil
}

func (agent GrpcAgent) GetLastId() (*gatewayProto.IntResponse, error) {
	return agent.AlbumGrpc.GetLastId(context.Background(), &emptypb.Empty{})
}

func (useCase AlbumGrpc) GetLastCoverId(context.Context, *emptypb.Empty) (*gatewayProto.IntResponse, error) {
	id, err := (*useCase.AlbumCoverRepo).GetLastId()
	if err != nil {
		return nil, err
	}

	return &gatewayProto.IntResponse{Data: id}, nil
}

func (agent GrpcAgent) GetLastCoverId() (*gatewayProto.IntResponse, error) {
	return agent.AlbumGrpc.GetLastCoverId(context.Background(), &emptypb.Empty{})
}

func (useCase AlbumGrpc) Create(ctx context.Context, album *albumProto.Album) (*emptypb.Empty, error) {
	err := (*useCase.AlbumRepo).Create(album)
	return &emptypb.Empty{}, err
}

func (agent GrpcAgent) Create(album *albumProto.Album) error {
	_, err := agent.AlbumGrpc.Create(context.Background(), album)
	return err
}

func (useCase AlbumGrpc) CreateCover(ctx context.Context, cover *albumProto.AlbumCover) (*emptypb.Empty, error) {
	err := (*useCase.AlbumCoverRepo).Create(cover)
	return &emptypb.Empty{}, err
}

func (agent GrpcAgent) CreateCover(cover *albumProto.AlbumCover) error {
	_, err := agent.AlbumGrpc.CreateCover(context.Background(), cover)
	return err
}

func (useCase AlbumGrpc) Update(ctx context.Context, album *albumProto.Album) (*emptypb.Empty, error) {
	err := (*useCase.AlbumRepo).Update(album)
	return &emptypb.Empty{}, err
}

func (agent GrpcAgent) Update(album *albumProto.Album) error {
	_, err := agent.AlbumGrpc.Update(context.Background(), album)
	return err
}

func (useCase AlbumGrpc) UpdateCover(ctx context.Context, cover *albumProto.AlbumCover) (*emptypb.Empty, error) {
	err := (*useCase.AlbumCoverRepo).Update(cover)
	return &emptypb.Empty{}, err
}

func (agent GrpcAgent) UpdateCover(cover *albumProto.AlbumCover) error {
	_, err := agent.AlbumGrpc.UpdateCover(context.Background(), cover)
	return err
}

func (useCase AlbumGrpc) Delete(ctx context.Context, data *gatewayProto.IdArg) (*emptypb.Empty, error) {
	err := (*useCase.AlbumRepo).Delete(data.Id)
	return &emptypb.Empty{}, err
}

func (agent GrpcAgent) Delete(data *gatewayProto.IdArg) error {
	_, err := agent.AlbumGrpc.Delete(context.Background(), data)
	return err
}

func (useCase AlbumGrpc) DeleteCover(ctx context.Context, data *gatewayProto.IdArg) (*emptypb.Empty, error) {
	err := (*useCase.AlbumCoverRepo).Delete(data.Id)
	return &emptypb.Empty{}, err
}

func (agent GrpcAgent) DeleteCover(data *gatewayProto.IdArg) error {
	_, err := agent.AlbumGrpc.DeleteCover(context.Background(), data)
	return err
}

func (useCase AlbumGrpc) GetById(ctx context.Context, data *gatewayProto.IdArg) (*albumProto.AlbumDataTransfer, error) {
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

func (agent GrpcAgent) GetById(data *gatewayProto.IdArg) (*albumProto.AlbumDataTransfer, error) {
	return agent.AlbumGrpc.GetById(context.Background(), data)
}

func (useCase AlbumGrpc) GetCoverById(ctx context.Context, data *gatewayProto.IdArg) (*albumProto.AlbumCoverDataTransfer, error) {
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

func (agent GrpcAgent) GetCoverById(data *gatewayProto.IdArg) (*albumProto.AlbumCoverDataTransfer, error) {
	return agent.AlbumGrpc.GetCoverById(context.Background(), data)
}

func (useCase AlbumGrpc) GetPopular(context.Context, *emptypb.Empty) (*albumProto.AlbumsResponse, error) {
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

func (agent GrpcAgent) GetPopular() (*albumProto.AlbumsResponse, error) {
	return agent.AlbumGrpc.GetPopular(context.Background(), &emptypb.Empty{})
}

func (useCase AlbumGrpc) GetAlbumsFromArtist(ctx context.Context, artistData *gatewayProto.IdArg) (*albumProto.AlbumsResponse, error) {
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

func (agent GrpcAgent) GetAlbumsFromArtist(artistData *gatewayProto.IdArg) (*albumProto.AlbumsResponse, error) {
	return agent.AlbumGrpc.GetAlbumsFromArtist(context.Background(), artistData)
}

func (useCase AlbumGrpc) GetSize(context.Context, *emptypb.Empty) (*gatewayProto.IntResponse, error) {
	size, err := (*useCase.AlbumRepo).GetSize()
	return &gatewayProto.IntResponse{Data: size}, err
}

func (agent GrpcAgent) GetSize() (*gatewayProto.IntResponse, error) {
	return agent.AlbumGrpc.GetSize(context.Background(), &emptypb.Empty{})
}

func (useCase AlbumGrpc) SearchByTitle(ctx context.Context, title *gatewayProto.StringArg) (*albumProto.AlbumsResponse, error) {
	albums, err := (*useCase.AlbumRepo).SearchByTitle(title.Str)

	dto := make([]*albumProto.AlbumDataTransfer, len(albums))

	for idx, album := range albums {
		data, err := useCase.CastToDTO(album)
		if err != nil {
			return nil, err
		}
		dto[idx] = data
	}

	return &albumProto.AlbumsResponse{Albums: dto}, err
}

func (agent GrpcAgent) SearchByTitle(title *gatewayProto.StringArg) (*albumProto.AlbumsResponse, error) {
	return agent.AlbumGrpc.SearchByTitle(context.Background(), title)
}

func (useCase AlbumGrpc) GetFavorites(ctx context.Context, data *gatewayProto.IdArg) (*albumProto.AlbumsResponse, error) {
	albums, err := (*useCase.AlbumRepo).GetFavorites(data.Id)
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

func (agent GrpcAgent) GetFavorites(data *gatewayProto.IdArg) (*albumProto.AlbumsResponse, error) {
	return agent.AlbumGrpc.GetFavorites(context.Background(), data)
}

func (useCase AlbumGrpc) AddToFavorites(ctx context.Context, data *gatewayProto.UserIdAlbumIdArg) (*emptypb.Empty, error) {
	if err := (*useCase.AlbumRepo).AddToFavorites(data.AlbumId, data.UserId); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (agent GrpcAgent) AddToFavorites(data *gatewayProto.UserIdAlbumIdArg) (*emptypb.Empty, error) {
	return agent.AlbumGrpc.AddToFavorites(context.Background(), data)
}

func (useCase AlbumGrpc) RemoveFromFavorites(ctx context.Context, data *gatewayProto.UserIdAlbumIdArg) (*emptypb.Empty, error) {
	if err := (*useCase.TrackRepo).RemoveFromFavorites(data.AlbumId, data.UserId); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (agent GrpcAgent) RemoveFromFavorites(data *gatewayProto.UserIdAlbumIdArg) (*emptypb.Empty, error) {
	return agent.AlbumGrpc.RemoveFromFavorites(context.Background(), data)
}
