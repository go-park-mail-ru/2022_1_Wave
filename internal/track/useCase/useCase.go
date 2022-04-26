package TrackUseCase

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway/gatewayProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/trackProto"
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
}
