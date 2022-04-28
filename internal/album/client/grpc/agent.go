package AlbumGrpcAgent

import (
	"context"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/album/albumProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway/gatewayProto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcAgent struct {
	AlbumGrpc albumProto.AlbumUseCaseClient
}

func MakeAgent(gRPC albumProto.AlbumUseCaseClient) GrpcAgent {
	return GrpcAgent{AlbumGrpc: gRPC}
}

func (agent GrpcAgent) GetAll() ([]*albumProto.Album, error) {
	data, err := agent.AlbumGrpc.GetAll(context.Background(), &emptypb.Empty{})
	return data.Albums, err
}

func (agent GrpcAgent) GetAllCovers() ([]*albumProto.AlbumCover, error) {
	data, err := agent.AlbumGrpc.GetAllCovers(context.Background(), &emptypb.Empty{})
	return data.Covers, err
}

func (agent GrpcAgent) GetLastId() (int64, error) {
	id, err := agent.AlbumGrpc.GetLastId(context.Background(), &emptypb.Empty{})
	if err != nil {
		return -1, err
	}
	return id.Data, nil
}

func (agent GrpcAgent) GetLastCoverId() (int64, error) {
	id, err := agent.AlbumGrpc.GetLastCoverId(context.Background(), &emptypb.Empty{})
	if err != nil {
		return -1, err
	}
	return id.Data, nil
}

func (agent GrpcAgent) Create(album *albumProto.Album) error {
	_, err := agent.AlbumGrpc.Create(context.Background(), album)
	return err
}

func (agent GrpcAgent) CreateCover(cover *albumProto.AlbumCover) error {
	_, err := agent.AlbumGrpc.CreateCover(context.Background(), cover)
	return err
}

func (agent GrpcAgent) Update(album *albumProto.Album) error {
	_, err := agent.AlbumGrpc.Update(context.Background(), album)
	return err
}

func (agent GrpcAgent) UpdateCover(cover *albumProto.AlbumCover) error {
	_, err := agent.AlbumGrpc.UpdateCover(context.Background(), cover)
	return err
}

func (agent GrpcAgent) Delete(id int64) error {
	_, err := agent.AlbumGrpc.Delete(context.Background(), &gatewayProto.IdArg{Id: id})
	return err
}

func (agent GrpcAgent) DeleteCover(id int64) error {
	_, err := agent.AlbumGrpc.DeleteCover(context.Background(), &gatewayProto.IdArg{Id: id})
	return err
}

func (agent GrpcAgent) GetById(id int64) (*albumProto.Album, error) {
	album, err := agent.AlbumGrpc.GetById(context.Background(), &gatewayProto.IdArg{Id: id})
	return album, err
}

func (agent GrpcAgent) GetCoverById(id int64) (*albumProto.AlbumCover, error) {
	return agent.AlbumGrpc.GetCoverById(context.Background(), &gatewayProto.IdArg{Id: id})
}

func (agent GrpcAgent) GetPopular() ([]*albumProto.Album, error) {
	data, err := agent.AlbumGrpc.GetPopular(context.Background(), &emptypb.Empty{})
	return data.Albums, err
}

func (agent GrpcAgent) GetAlbumsFromArtist(artistId int64) ([]*albumProto.Album, error) {
	albums, err := agent.AlbumGrpc.GetAlbumsFromArtist(context.Background(), &gatewayProto.IdArg{Id: artistId})
	return albums.Albums, err
}

func (agent GrpcAgent) GetSize() (int64, error) {
	size, err := agent.AlbumGrpc.GetSize(context.Background(), &emptypb.Empty{})
	return size.Data, err
}

func (agent GrpcAgent) SearchByTitle(title string) ([]*albumProto.Album, error) {
	data, err := agent.AlbumGrpc.SearchByTitle(context.Background(), &gatewayProto.StringArg{Str: title})
	return data.Albums, err
}

func (agent GrpcAgent) GetFavorites(id int64) ([]*albumProto.Album, error) {
	data, err := agent.AlbumGrpc.GetFavorites(context.Background(), &gatewayProto.IdArg{Id: id})
	return data.Albums, err
}

func (agent GrpcAgent) AddToFavorites(userId int64, albumId int64) error {
	_, err := agent.AlbumGrpc.AddToFavorites(context.Background(), &gatewayProto.UserIdAlbumIdArg{
		UserId:  userId,
		AlbumId: albumId,
	})
	return err
}

func (agent GrpcAgent) RemoveFromFavorites(userId int64, albumId int64) error {
	_, err := agent.AlbumGrpc.RemoveFromFavorites(context.Background(), &gatewayProto.UserIdAlbumIdArg{
		UserId:  userId,
		AlbumId: albumId,
	})
	return err
}
