package domain

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/album/albumProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/artist/artistProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/playlist/playlistProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/trackProto"
)

type AlbumAgent interface {
	GetAll() ([]*albumProto.Album, error)
	GetAllCovers() ([]*albumProto.AlbumCover, error)
	GetLastId() (int64, error)
	GetLastCoverId() (int64, error)
	Create(*albumProto.Album) error
	CreateCover(*albumProto.AlbumCover) error
	Update(*albumProto.Album) error
	UpdateCover(*albumProto.AlbumCover) error
	Delete(int64) error
	DeleteCover(int64) error
	GetById(int64) (*albumProto.Album, error)
	GetCoverById(int64) (*albumProto.AlbumCover, error)
	GetPopular() ([]*albumProto.Album, error)
	GetAlbumsFromArtist(int64) ([]*albumProto.Album, error)
	GetSize() (int64, error)
	SearchByTitle(title string) ([]*albumProto.Album, error)
	GetFavorites(userId int64) ([]*albumProto.Album, error)
	AddToFavorites(userId int64, albumId int64) error
	RemoveFromFavorites(userId int64, albumId int64) error
}

type ArtistAgent interface {
	GetAll() ([]*artistProto.Artist, error)
	GetLastId() (int64, error)
	Create(*artistProto.Artist) error
	Update(*artistProto.Artist) error
	Delete(int64) error
	GetById(int64) (*artistProto.Artist, error)
	GetPopular() ([]*artistProto.Artist, error)
	GetSize() (int64, error)
	SearchByName(name string) ([]*artistProto.Artist, error)
	GetFavorites(int64) ([]*artistProto.Artist, error)
	AddToFavorites(userId int64, artistId int64) error
	RemoveFromFavorites(userId int64, artistId int64) error
}

type TrackAgent interface {
	GetAll() ([]*trackProto.Track, error)
	GetLastId() (int64, error)
	Create(*trackProto.Track) error
	Update(*trackProto.Track) error
	Delete(int64) error
	GetById(int64) (*trackProto.Track, error)
	GetPopular() ([]*trackProto.Track, error)
	GetTracksFromAlbum(int64) ([]*trackProto.Track, error)
	GetPopularTracksFromArtist(int64) ([]*trackProto.Track, error)
	GetSize() (int64, error)
	Like(userId int64, trackId int64) error
	Listen(int64) error
	SearchByTitle(title string) ([]*trackProto.Track, error)
	GetFavorites(int64) ([]*trackProto.Track, error)
	AddToFavorites(userId int64, trackId int64) error
	RemoveFromFavorites(userId int64, artistId int64) error
}

type PlaylistAgent interface {
	GetAll(userId int64) ([]*playlistProto.Playlist, error)
	GetLastId(userId int64) (int64, error)
	Create(userId int64, playlist *playlistProto.Playlist) error
	Update(userId int64, playlist *playlistProto.Playlist) error
	Delete(userId int64, playlistId int64) error
	GetById(userId int64, playlistId int64) (*playlistProto.Playlist, error)
	GetSize(userId int64) (int64, error)
}
