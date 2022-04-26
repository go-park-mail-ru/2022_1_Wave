package domain

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/album/albumProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/artist/artistProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/trackProto"
)

type AlbumRepo interface {
	Create(*albumProto.Album) error
	Update(*albumProto.Album) error
	Delete(int64) error
	SelectByID(int64) (*albumProto.Album, error)
	GetAll() ([]*albumProto.Album, error)
	GetPopular() ([]*albumProto.Album, error)
	GetLastId() (id int64, err error)
	GetSize() (int64, error)
	GetAlbumsFromArtist(artist int64) ([]*albumProto.Album, error)
}

type AlbumCoverRepo interface {
	Create(*albumProto.AlbumCover) error
	Update(*albumProto.AlbumCover) error
	Delete(int64) error
	SelectByID(int64) (*albumProto.AlbumCover, error)
	GetAll() ([]*albumProto.AlbumCover, error)
	GetLastId() (id int64, err error)
	GetSize() (int64, error)
}

type ArtistRepo interface {
	Create(*artistProto.Artist) error
	Update(*artistProto.Artist) error
	Delete(int64) error
	SelectByID(int64) (*artistProto.Artist, error)
	GetAll() ([]*artistProto.Artist, error)
	GetPopular() ([]*artistProto.Artist, error)
	GetLastId() (id int64, err error)
	GetSize() (int64, error)
}

type TrackRepo interface {
	Create(*trackProto.Track) error
	Update(*trackProto.Track) error
	Delete(int64) error
	SelectByID(int64) (*trackProto.Track, error)
	GetAll() ([]*trackProto.Track, error)
	GetPopular() ([]*trackProto.Track, error)
	GetLastId() (id int64, err error)
	GetSize() (int64, error)
	GetTracksFromAlbum(albumId int64) ([]*trackProto.Track, error)
	GetPopularTracksFromArtist(artistId int64) ([]*trackProto.Track, error)
	Like(id int64) error
	Listen(id int64) error
}
