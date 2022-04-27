package TrackUseCase

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway/gatewayProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/trackProto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TrackAgent interface {
	GetAll() (*trackProto.TracksResponse, error)
	GetLastId() (*gatewayProto.IntResponse, error)
	Create(*trackProto.Track) error
	Update(*trackProto.Track) error
	Delete(*gatewayProto.IdArg) error
	GetById(*gatewayProto.IdArg) (*trackProto.TrackDataTransfer, error)
	GetPopular() (*trackProto.TracksResponse, error)
	GetTracksFromAlbum(*gatewayProto.IdArg) (*trackProto.TracksResponse, error)
	GetPopularTracksFromArtist(*gatewayProto.IdArg) (*trackProto.TracksResponse, error)
	GetSize() (*gatewayProto.IntResponse, error)
	Like(arg *gatewayProto.IdArg) error
	Listen(arg *gatewayProto.IdArg) error
	SearchByTitle(arg *gatewayProto.StringArg) (*trackProto.TracksResponse, error)
	GetFavorites(*gatewayProto.IdArg) (*trackProto.TracksResponse, error)
	AddToFavorites(data *gatewayProto.UserIdTrackIdArg) (*emptypb.Empty, error)
	RemoveFromFavorites(data *gatewayProto.UserIdTrackIdArg) (*emptypb.Empty, error)
}
