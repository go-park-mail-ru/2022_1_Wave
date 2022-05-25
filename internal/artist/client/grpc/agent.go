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

func (agent GrpcAgent) GetAll() ([]*artistProto.Artist, error) {
	data, err := agent.ArtistGrpc.GetAll(context.Background(), &emptypb.Empty{})
	return data.Artists, err
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

func (agent GrpcAgent) GetById(id int64) (*artistProto.Artist, error) {
	data, err := agent.ArtistGrpc.GetById(context.Background(), &gatewayProto.IdArg{Id: id})
	return data, err
}

func (agent GrpcAgent) GetPopular() ([]*artistProto.Artist, error) {
	data, err := agent.ArtistGrpc.GetPopular(context.Background(), &emptypb.Empty{})
	return data.Artists, err
}

func (agent GrpcAgent) GetSize() (int64, error) {
	size, err := agent.ArtistGrpc.GetSize(context.Background(), &emptypb.Empty{})
	return size.Data, err
}

func (agent GrpcAgent) SearchByName(name string) ([]*artistProto.Artist, error) {
	data, err := agent.ArtistGrpc.SearchByName(context.Background(), &gatewayProto.StringArg{Str: name})
	return data.Artists, err
}

func (agent GrpcAgent) GetFavorites(userId int64) ([]*artistProto.Artist, error) {
	data, err := agent.ArtistGrpc.GetFavorites(context.Background(), &gatewayProto.IdArg{Id: userId})
	return data.Artists, err
}

func (agent GrpcAgent) AddToFavorites(userId int64, albumId int64) error {
	_, err := agent.ArtistGrpc.AddToFavorites(context.Background(), &gatewayProto.UserIdArtistIdArg{
		UserId:   userId,
		ArtistId: albumId,
	})
	return err
}

func (agent GrpcAgent) RemoveFromFavorites(userId int64, albumId int64) error {
	_, err := agent.ArtistGrpc.RemoveFromFavorites(context.Background(), &gatewayProto.UserIdArtistIdArg{
		UserId:   userId,
		ArtistId: albumId,
	})
	return err
}

func (agent GrpcAgent) Like(userId int64, id int64) error {
	_, err := agent.ArtistGrpc.Like(context.Background(), &gatewayProto.UserIdArtistIdArg{
		UserId:   userId,
		ArtistId: id,
	})
	return err
}

func (agent GrpcAgent) LikeCheckByUser(userId int64, id int64) (bool, error) {
	liked, err := agent.ArtistGrpc.LikeCheckByUser(context.Background(), &gatewayProto.UserIdArtistIdArg{
		UserId:   userId,
		ArtistId: id,
	})
	if err != nil {
		return false, err
	}
	return liked.Ok, nil
}
