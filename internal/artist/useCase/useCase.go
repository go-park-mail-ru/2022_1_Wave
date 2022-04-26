package ArtistUseCase

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/artist/artistProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway/gatewayProto"
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
}
