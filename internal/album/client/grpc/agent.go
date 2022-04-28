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

func (agent GrpcAgent) GetAll() (*albumProto.AlbumsResponse, error) {
	return agent.AlbumGrpc.GetAll(context.Background(), &emptypb.Empty{})
}

func (agent GrpcAgent) GetAllCovers() (*albumProto.AlbumsCoverResponse, error) {
	return agent.AlbumGrpc.GetAllCovers(context.Background(), &emptypb.Empty{})
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

func (agent GrpcAgent) GetById(id int64) (*albumProto.AlbumDataTransfer, error) {
	return agent.AlbumGrpc.GetById(context.Background(), &gatewayProto.IdArg{Id: id})
}

func (agent GrpcAgent) GetCoverById(id int64) (*albumProto.AlbumCoverDataTransfer, error) {
	return agent.AlbumGrpc.GetCoverById(context.Background(), &gatewayProto.IdArg{Id: id})
}

func (agent GrpcAgent) GetPopular() (*albumProto.AlbumsResponse, error) {
	return agent.AlbumGrpc.GetPopular(context.Background(), &emptypb.Empty{})
}

func (agent GrpcAgent) GetAlbumsFromArtist(artistData *gatewayProto.IdArg) (*albumProto.AlbumsResponse, error) {
	return agent.AlbumGrpc.GetAlbumsFromArtist(context.Background(), artistData)
}

func (agent GrpcAgent) GetSize() (*gatewayProto.IntResponse, error) {
	return agent.AlbumGrpc.GetSize(context.Background(), &emptypb.Empty{})
}

func (agent GrpcAgent) SearchByTitle(title string) (*albumProto.AlbumsResponse, error) {
	return agent.AlbumGrpc.SearchByTitle(context.Background(), &gatewayProto.StringArg{Str: title})
}

func (agent GrpcAgent) GetFavorites(id int64) (*albumProto.AlbumsResponse, error) {
	return agent.AlbumGrpc.GetFavorites(context.Background(), &gatewayProto.IdArg{Id: id})
}

func (agent GrpcAgent) AddToFavorites(userId int64, albumId int64) (*emptypb.Empty, error) {
	return agent.AlbumGrpc.AddToFavorites(context.Background(), &gatewayProto.UserIdAlbumIdArg{
		UserId:  userId,
		AlbumId: albumId,
	})
}

func (agent GrpcAgent) RemoveFromFavorites(userId int64, albumId int64) (*emptypb.Empty, error) {
	return agent.AlbumGrpc.RemoveFromFavorites(context.Background(), &gatewayProto.UserIdAlbumIdArg{
		UserId:  userId,
		AlbumId: albumId,
	})
}
