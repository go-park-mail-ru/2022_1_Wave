package domain

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/album/albumProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/artist/artistProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/playlist/playlistProto"
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
	SearchByTitle(title string) ([]*albumProto.Album, error)
	Like(id int64, userId int64) error
	LikeCheckByUser(id int64, userId int64) (bool, error)
	Listen(id int64) error
	GetFavorites(userId int64) ([]*albumProto.Album, error)
	AddToFavorites(id int64, userId int64) error
	RemoveFromFavorites(albumId int64, userId int64) error
	GetPopularAlbumOfWeekTop20() ([]*albumProto.Album, error)
	CountPopularAlbumOfWeek() (bool, error)
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
	SearchByName(string) ([]*artistProto.Artist, error)
	Like(id int64, userId int64) error
	LikeCheckByUser(id int64, userId int64) (bool, error)
	Listen(id int64) error
	GetFavorites(userId int64) ([]*artistProto.Artist, error)
	AddToFavorites(id int64, userId int64) error
	RemoveFromFavorites(artistId int64, userId int64) error
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
	GetTracksFromPlaylist(playlistId int64) ([]*trackProto.Track, error)
	GetPopularTracksFromArtist(artistId int64) ([]*trackProto.Track, error)
	Like(id int64, userId int64) error
	LikeCheckByUser(id int64, userId int64) (bool, error)
	Listen(id int64) error
	SearchByTitle(title string) ([]*trackProto.Track, error)
	GetFavorites(userId int64) ([]*trackProto.Track, error)
	AddToFavorites(id int64, userId int64) error
	RemoveFromFavorites(trackId int64, userId int64) error
	GetPopularTrackOfWeekTop20() ([]*trackProto.Track, error)
	CountPopularTrackOfWeek() (bool, error)
}

type PlaylistRepo interface {
	Create(userId int64, playlist *playlistProto.Playlist) error
	Update(userId int64, playlist *playlistProto.Playlist) error
	Delete(userId int64, playlistId int64) error
	SelectByIDOfCurrentUser(userId int64, playlistId int64) (*playlistProto.Playlist, error)
	SelectById(playlistId int64) (*playlistProto.Playlist, error)
	GetAllOfCurrentUser(userId int64) ([]*playlistProto.Playlist, error)
	GetAll() ([]*playlistProto.Playlist, error)
	GetLastIdOfCurrentUser(userId int64) (id int64, err error)
	GetLastId() (id int64, err error)
	GetSizeOfCurrentUser(userId int64) (int64, error)
	GetSize() (int64, error)
	AddToPlaylist(userId int64, playlistId int64, trackId int64) error
	RemoveFromPlaylist(userId int64, playlistId int64, trackId int64) error
}

type LinkerRepo interface {
	Get(hash string) (string, error)
	Create(url string) (string, error)
	Count(hash string) (int64, error)
}
