package ArtistUseCase

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/artist/artistProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway/gatewayProto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ArtistAgent interface {
	GetAll() (*artistProto.ArtistsResponse, error)
	GetLastId() (*gatewayProto.IntResponse, error)
	Create(*artistProto.Artist) error
	Update(*artistProto.Artist) error
	Delete(*gatewayProto.IdArg) error
	GetById(*gatewayProto.IdArg) (*artistProto.ArtistDataTransfer, error)
	GetPopular() (*artistProto.ArtistsResponse, error)
	GetSize() (*gatewayProto.IntResponse, error)
	SearchByName(arg *gatewayProto.StringArg) (*artistProto.ArtistsResponse, error)
	GetFavorites(*gatewayProto.IdArg) (*artistProto.ArtistsResponse, error)
	AddToFavorites(data *gatewayProto.UserIdArtistIdArg) (*emptypb.Empty, error)
	RemoveFromFavorites(data *gatewayProto.UserIdArtistIdArg) (*emptypb.Empty, error)
}
