package AlbumUseCase

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/album/albumProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway/gatewayProto"
)

type AlbumAgent interface {
	GetAll() (*albumProto.AlbumsResponse, error)
	GetAllCovers() (*albumProto.AlbumsCoverResponse, error)
	GetLastId() (*gatewayProto.IntResponse, error)
	GetLastCoverId() (*gatewayProto.IntResponse, error)
	Create(*albumProto.Album) error
	CreateCover(*albumProto.AlbumCover) error
	Update(*albumProto.Album) error
	UpdateCover(*albumProto.AlbumCover) error
	Delete(*gatewayProto.IdArg) error
	DeleteCover(*gatewayProto.IdArg) error
	GetById(*gatewayProto.IdArg) (*albumProto.AlbumDataTransfer, error)
	GetCoverById(*gatewayProto.IdArg) (*albumProto.AlbumCoverDataTransfer, error)
	GetPopular() (*albumProto.AlbumsResponse, error)
	GetAlbumsFromArtist(*gatewayProto.IdArg) (*albumProto.AlbumsResponse, error)
	GetSize() (*gatewayProto.IntResponse, error)
}
