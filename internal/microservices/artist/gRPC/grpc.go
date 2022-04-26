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

type GrpcAgent struct {
	ArtistGrpc artistProto.ArtistUseCaseClient
}

func MakeAgent(gRPC artistProto.ArtistUseCaseClient) GrpcAgent {
	return GrpcAgent{ArtistGrpc: gRPC}
}

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

func (agent GrpcAgent) GetAll() (*artistProto.ArtistsResponse, error) {
	return agent.ArtistGrpc.GetAll(context.Background(), &emptypb.Empty{})
}

func (useCase ArtistGrpc) GetLastId(context.Context, *emptypb.Empty) (*gatewayProto.IntResponse, error) {
	id, err := (*useCase.ArtistRepo).GetLastId()
	if err != nil {
		return nil, err
	}

	return &gatewayProto.IntResponse{Data: id}, nil
}

func (agent GrpcAgent) GetLastId() (*gatewayProto.IntResponse, error) {
	return agent.ArtistGrpc.GetLastId(context.Background(), &emptypb.Empty{})
}

func (useCase ArtistGrpc) Create(ctx context.Context, artist *artistProto.Artist) (*emptypb.Empty, error) {
	err := (*useCase.ArtistRepo).Create(artist)
	return &emptypb.Empty{}, err
}

func (agent GrpcAgent) Create(artist *artistProto.Artist) error {
	_, err := agent.ArtistGrpc.Create(context.Background(), artist)
	return err
}

func (useCase ArtistGrpc) Update(ctx context.Context, artist *artistProto.Artist) (*emptypb.Empty, error) {
	err := (*useCase.ArtistRepo).Update(artist)
	return &emptypb.Empty{}, err
}

func (agent GrpcAgent) Update(artist *artistProto.Artist) error {
	_, err := agent.ArtistGrpc.Update(context.Background(), artist)
	return err
}

func (useCase ArtistGrpc) Delete(ctx context.Context, data *gatewayProto.IdArg) (*emptypb.Empty, error) {
	err := (*useCase.ArtistRepo).Delete(data.Id)
	return &emptypb.Empty{}, err
}

func (agent GrpcAgent) Delete(data *gatewayProto.IdArg) error {
	_, err := agent.ArtistGrpc.Delete(context.Background(), data)
	return err
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

func (agent GrpcAgent) GetById(data *gatewayProto.IdArg) (*artistProto.ArtistDataTransfer, error) {
	return agent.ArtistGrpc.GetById(context.Background(), data)
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

func (agent GrpcAgent) GetPopular() (*artistProto.ArtistsResponse, error) {
	return agent.ArtistGrpc.GetPopular(context.Background(), &emptypb.Empty{})
}

func (useCase ArtistGrpc) GetSize(context.Context, *emptypb.Empty) (*gatewayProto.IntResponse, error) {
	size, err := (*useCase.ArtistRepo).GetSize()
	return &gatewayProto.IntResponse{Data: size}, err
}

func (agent GrpcAgent) GetSize() (*gatewayProto.IntResponse, error) {
	return agent.ArtistGrpc.GetSize(context.Background(), &emptypb.Empty{})
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

func (agent GrpcAgent) SearchByName(title *gatewayProto.StringArg) (*artistProto.ArtistsResponse, error) {
	return agent.ArtistGrpc.SearchByName(context.Background(), title)
}
