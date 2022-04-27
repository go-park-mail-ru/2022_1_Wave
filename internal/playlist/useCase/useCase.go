package PlaylistUseCase

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway/gatewayProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/playlist/playlistProto"
)

type PlaylistAgent interface {
	GetAll(*gatewayProto.IdArg) (*playlistProto.PlaylistsResponse, error)
	GetLastId(*gatewayProto.IdArg) (*gatewayProto.IntResponse, error)
	Create(*playlistProto.UserIdPlaylistArg) error
	Update(*playlistProto.UserIdPlaylistArg) error
	Delete(*playlistProto.UserIdPlaylistIdArg) error
	GetById(*playlistProto.UserIdPlaylistIdArg) (*playlistProto.PlaylistDataTransfer, error)
	GetSize(userId *gatewayProto.IdArg) (*gatewayProto.IntResponse, error)
}
