package ArtistGrpcAgent

import (
	"context"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/artist/artistProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway/gatewayProto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcAgent struct {
	ArtistGrpc artistProto.ArtistUseCaseClient
}

func MakeAgent(gRPC artistProto.ArtistUseCaseClient) GrpcAgent {
	return GrpcAgent{ArtistGrpc: gRPC}
}

func (agent GrpcAgent) GetAll() (*artistProto.ArtistsResponse, error) {
	return agent.ArtistGrpc.GetAll(context.Background(), &emptypb.Empty{})
}

func (agent GrpcAgent) GetLastId() (int64, error) {
	id, err := agent.ArtistGrpc.GetLastId(context.Background(), &emptypb.Empty{})
	if err != nil {
		return -1, err
	}
	return id.Data, nil
}

func (agent GrpcAgent) Create(artist *artistProto.Artist) error {
	_, err := agent.ArtistGrpc.Create(context.Background(), artist)
	return err
}

func (agent GrpcAgent) Update(artist *artistProto.Artist) error {
	_, err := agent.ArtistGrpc.Update(context.Background(), artist)
	return err
}

func (agent GrpcAgent) Delete(id int64) error {
	_, err := agent.ArtistGrpc.Delete(context.Background(), &gatewayProto.IdArg{Id: id})
	return err
}

func (agent GrpcAgent) GetById(id int64) (*artistProto.ArtistDataTransfer, error) {
	return agent.ArtistGrpc.GetById(context.Background(), &gatewayProto.IdArg{Id: id})
}

func (agent GrpcAgent) GetPopular() (*artistProto.ArtistsResponse, error) {
	return agent.ArtistGrpc.GetPopular(context.Background(), &emptypb.Empty{})
}

func (agent GrpcAgent) GetSize() (*gatewayProto.IntResponse, error) {
	return agent.ArtistGrpc.GetSize(context.Background(), &emptypb.Empty{})
}

func (agent GrpcAgent) SearchByName(title *gatewayProto.StringArg) (*artistProto.ArtistsResponse, error) {
	return agent.ArtistGrpc.SearchByName(context.Background(), title)
}

func (agent GrpcAgent) GetFavorites(data *gatewayProto.IdArg) (*artistProto.ArtistsResponse, error) {
	return agent.ArtistGrpc.GetFavorites(context.Background(), data)
}

func (agent GrpcAgent) AddToFavorites(data *gatewayProto.UserIdArtistIdArg) (*emptypb.Empty, error) {
	return agent.ArtistGrpc.AddToFavorites(context.Background(), data)
}

func (agent GrpcAgent) RemoveFromFavorites(data *gatewayProto.UserIdArtistIdArg) (*emptypb.Empty, error) {
	return agent.ArtistGrpc.RemoveFromFavorites(context.Background(), data)
}
