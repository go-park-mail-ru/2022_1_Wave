package TrackGrpc

import (
	"context"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	Gateway "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway/gatewayProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/trackProto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcAgent struct {
	TrackGrpc trackProto.TrackUseCaseClient
}

func MakeAgent(gRPC trackProto.TrackUseCaseClient) GrpcAgent {
	return GrpcAgent{TrackGrpc: gRPC}
}

type TrackGrpc struct {
	TrackRepo  *domain.TrackRepo
	ArtistRepo *domain.ArtistRepo
	trackProto.UnimplementedTrackUseCaseServer
}

func MakeTrackGrpc(track domain.TrackRepo, artist domain.ArtistRepo) TrackGrpc {
	return TrackGrpc{
		TrackRepo:  &track,
		ArtistRepo: &artist,
	}
}

func (useCase TrackGrpc) CastToDTO(track *trackProto.Track) (*trackProto.TrackDataTransfer, error) {
	artist, err := (*useCase.ArtistRepo).SelectByID(track.ArtistId)
	if err != nil {
		return nil, err
	}

	trackDto, err := Gateway.CastTrackToDtoWithoutArtistName(track)
	if err != nil {
		return nil, err
	}

	trackDto.Artist = artist.Name
	return trackDto, nil
}

func (useCase TrackGrpc) GetAll(context.Context, *emptypb.Empty) (*trackProto.TracksResponse, error) {
	tracks, err := (*useCase.TrackRepo).GetAll()

	if err != nil {
		return nil, err
	}

	dto := make([]*trackProto.TrackDataTransfer, len(tracks))

	for idx, obj := range tracks {
		result, err := useCase.CastToDTO(obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = result
	}

	return &trackProto.TracksResponse{Tracks: dto}, nil
}

func (agent GrpcAgent) GetAll() (*trackProto.TracksResponse, error) {
	return agent.TrackGrpc.GetAll(context.Background(), &emptypb.Empty{})
}

func (useCase TrackGrpc) GetLastId(context.Context, *emptypb.Empty) (*gatewayProto.IntResponse, error) {
	id, err := (*useCase.TrackRepo).GetLastId()
	if err != nil {
		return nil, err
	}

	return &gatewayProto.IntResponse{Data: id}, nil
}

func (agent GrpcAgent) GetLastId() (*gatewayProto.IntResponse, error) {
	return agent.TrackGrpc.GetLastId(context.Background(), &emptypb.Empty{})
}

func (useCase TrackGrpc) Create(ctx context.Context, dom *trackProto.Track) (*emptypb.Empty, error) {
	err := (*useCase.TrackRepo).Create(dom)
	return &emptypb.Empty{}, err
}

func (agent GrpcAgent) Create(album *trackProto.Track) error {
	_, err := agent.TrackGrpc.Create(context.Background(), album)
	return err
}

func (useCase TrackGrpc) Update(ctx context.Context, dom *trackProto.Track) (*emptypb.Empty, error) {
	err := (*useCase.TrackRepo).Update(dom)
	return &emptypb.Empty{}, err
}

func (agent GrpcAgent) Update(album *trackProto.Track) error {
	_, err := agent.TrackGrpc.Update(context.Background(), album)
	return err
}

func (useCase TrackGrpc) Delete(ctx context.Context, data *gatewayProto.IdArg) (*emptypb.Empty, error) {
	err := (*useCase.TrackRepo).Delete(data.Id)
	return &emptypb.Empty{}, err
}

func (agent GrpcAgent) Delete(data *gatewayProto.IdArg) error {
	_, err := agent.TrackGrpc.Delete(context.Background(), data)
	return err
}

func (useCase TrackGrpc) GetById(ctx context.Context, data *gatewayProto.IdArg) (*trackProto.TrackDataTransfer, error) {
	track, err := (*useCase.TrackRepo).SelectByID(data.Id)
	if err != nil {
		return nil, err
	}
	dto, err := useCase.CastToDTO(track)
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (agent GrpcAgent) GetById(data *gatewayProto.IdArg) (*trackProto.TrackDataTransfer, error) {
	return agent.TrackGrpc.GetById(context.Background(), data)
}

func (useCase TrackGrpc) GetPopular(context.Context, *emptypb.Empty) (*trackProto.TracksResponse, error) {
	tracks, err := (*useCase.TrackRepo).GetPopular()
	if err != nil {
		return nil, err
	}

	dto := make([]*trackProto.TrackDataTransfer, len(tracks))

	for i := 0; i < len(tracks); i++ {
		data, err := useCase.CastToDTO(tracks[i])
		if err != nil {
			return nil, err
		}
		dto[i] = data
	}

	return &trackProto.TracksResponse{Tracks: dto}, nil
}

func (agent GrpcAgent) GetPopular() (*trackProto.TracksResponse, error) {
	return agent.TrackGrpc.GetPopular(context.Background(), &emptypb.Empty{})
}

func (useCase TrackGrpc) GetTracksFromAlbum(ctx context.Context, data *gatewayProto.IdArg) (*trackProto.TracksResponse, error) {
	tracks, err := (*useCase.TrackRepo).GetTracksFromAlbum(data.Id)
	if err != nil {
		return nil, err
	}

	dto := make([]*trackProto.TrackDataTransfer, len(tracks))

	for idx, obj := range tracks {
		result, err := useCase.CastToDTO(obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = result
	}

	return &trackProto.TracksResponse{Tracks: dto}, nil
}

func (agent GrpcAgent) GetTracksFromAlbum(data *gatewayProto.IdArg) (*trackProto.TracksResponse, error) {
	return agent.TrackGrpc.GetTracksFromAlbum(context.Background(), data)
}

func (useCase TrackGrpc) GetPopularTracksFromArtist(ctx context.Context, data *gatewayProto.IdArg) (*trackProto.TracksResponse, error) {
	tracks, err := (*useCase.TrackRepo).GetPopularTracksFromArtist(data.Id)
	if err != nil {
		return nil, err
	}

	dto := make([]*trackProto.TrackDataTransfer, len(tracks))

	for idx, obj := range tracks {
		result, err := useCase.CastToDTO(obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = result
	}

	return &trackProto.TracksResponse{Tracks: dto}, nil
}

func (agent GrpcAgent) GetPopularTracksFromArtist(data *gatewayProto.IdArg) (*trackProto.TracksResponse, error) {
	return agent.TrackGrpc.GetPopularTracksFromArtist(context.Background(), data)
}

func (useCase TrackGrpc) GetSize(context.Context, *emptypb.Empty) (*gatewayProto.IntResponse, error) {
	size, err := (*useCase.TrackRepo).GetSize()
	return &gatewayProto.IntResponse{Data: size}, err
}

func (agent GrpcAgent) GetSize() (*gatewayProto.IntResponse, error) {
	return agent.TrackGrpc.GetSize(context.Background(), &emptypb.Empty{})
}

func (useCase TrackGrpc) Like(ctx context.Context, data *gatewayProto.IdArg) (*emptypb.Empty, error) {
	err := (*useCase.TrackRepo).Like(data.Id)
	return &emptypb.Empty{}, err
}

func (agent GrpcAgent) Like(data *gatewayProto.IdArg) error {
	_, err := agent.TrackGrpc.Like(context.Background(), &gatewayProto.IdArg{Id: data.Id})
	return err
}

func (useCase TrackGrpc) Listen(ctx context.Context, data *gatewayProto.IdArg) (*emptypb.Empty, error) {
	err := (*useCase.TrackRepo).Listen(data.Id)
	return &emptypb.Empty{}, err
}

func (agent GrpcAgent) Listen(data *gatewayProto.IdArg) error {
	_, err := agent.TrackGrpc.Listen(context.Background(), &gatewayProto.IdArg{Id: data.Id})
	return err
}
