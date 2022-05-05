package TrackGrpcAgent

import (
	"context"
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

func (agent GrpcAgent) GetAll() ([]*trackProto.Track, error) {
	tracks, err := agent.TrackGrpc.GetAll(context.Background(), &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return tracks.Tracks, err
}

func (agent GrpcAgent) GetLastId() (int64, error) {
	id, err := agent.TrackGrpc.GetLastId(context.Background(), &emptypb.Empty{})
	if err != nil {
		return -1, err
	}
	return id.Data, nil
}

func (agent GrpcAgent) Create(album *trackProto.Track) error {
	_, err := agent.TrackGrpc.Create(context.Background(), album)
	return err
}

func (agent GrpcAgent) Update(album *trackProto.Track) error {
	_, err := agent.TrackGrpc.Update(context.Background(), album)
	return err
}

func (agent GrpcAgent) Delete(id int64) error {
	_, err := agent.TrackGrpc.Delete(context.Background(), &gatewayProto.IdArg{Id: id})
	return err
}

func (agent GrpcAgent) GetById(id int64) (*trackProto.Track, error) {
	return agent.TrackGrpc.GetById(context.Background(), &gatewayProto.IdArg{Id: id})
}

func (agent GrpcAgent) GetPopular() ([]*trackProto.Track, error) {
	data, err := agent.TrackGrpc.GetPopular(context.Background(), &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return data.Tracks, err
}

func (agent GrpcAgent) GetTracksFromAlbum(id int64) ([]*trackProto.Track, error) {
	data, err := agent.TrackGrpc.GetTracksFromAlbum(context.Background(), &gatewayProto.IdArg{Id: id})
	if err != nil {
		return nil, err
	}
	return data.Tracks, err
}

func (agent GrpcAgent) GetPopularTracksFromArtist(artistId int64) ([]*trackProto.Track, error) {
	data, err := agent.TrackGrpc.GetPopularTracksFromArtist(context.Background(), &gatewayProto.IdArg{Id: artistId})
	if err != nil {
		return nil, err
	}
	return data.Tracks, err

}

func (agent GrpcAgent) GetSize() (int64, error) {
	size, err := agent.TrackGrpc.GetSize(context.Background(), &emptypb.Empty{})

	if err != nil {
		return -1, err
	}

	return size.Data, nil
}

func (agent GrpcAgent) Like(userId int64, id int64) error {
	_, err := agent.TrackGrpc.Like(context.Background(), &gatewayProto.UserIdTrackIdArg{
		UserId:  userId,
		TrackId: id,
	})
	return err
}

func (agent GrpcAgent) LikeCheckByUser(userId int64, id int64) (bool, error) {
	liked, err := agent.TrackGrpc.LikeCheckByUser(context.Background(), &gatewayProto.UserIdTrackIdArg{
		UserId:  userId,
		TrackId: id,
	})
	if err != nil {
		return false, err
	}
	return liked.Ok, nil
}

func (agent GrpcAgent) Listen(id int64) error {
	_, err := agent.TrackGrpc.Listen(context.Background(), &gatewayProto.IdArg{Id: id})
	return err
}

func (agent GrpcAgent) SearchByTitle(title string) ([]*trackProto.Track, error) {
	data, err := agent.TrackGrpc.SearchByTitle(context.Background(), &gatewayProto.StringArg{Str: title})
	if err != nil {
		return nil, err
	}
	return data.Tracks, nil
}

func (agent GrpcAgent) GetFavorites(userId int64) ([]*trackProto.Track, error) {
	data, err := agent.TrackGrpc.GetFavorites(context.Background(), &gatewayProto.IdArg{Id: userId})
	if err != nil {
		return nil, err
	}
	return data.Tracks, nil
}

func (agent GrpcAgent) AddToFavorites(userId int64, trackId int64) error {
	_, err := agent.TrackGrpc.AddToFavorites(context.Background(), &gatewayProto.UserIdTrackIdArg{
		UserId:  userId,
		TrackId: trackId,
	})
	return err

}

func (agent GrpcAgent) RemoveFromFavorites(userId int64, trackId int64) error {
	_, err := agent.TrackGrpc.RemoveFromFavorites(context.Background(), &gatewayProto.UserIdTrackIdArg{
		UserId:  userId,
		TrackId: trackId,
	})
	return err

}

func (agent GrpcAgent) GetTracksFromPlaylist(playlistId int64) ([]*trackProto.Track, error) {
	data, err := agent.TrackGrpc.GetTracksFromPlaylist(context.Background(), &gatewayProto.IdArg{Id: playlistId})
	if err != nil {
		return nil, err
	}
	return data.Tracks, nil

}
