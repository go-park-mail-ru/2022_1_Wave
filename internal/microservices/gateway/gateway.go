package Gateway

import (
	"fmt"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/album/albumProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/artist/artistProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/playlist/playlistProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/trackProto"
	"gopkg.in/validator.v2"
)

type CheckConstraint interface {
	*albumProto.Album | *albumProto.AlbumCover | *artistProto.Artist | *trackProto.Track | *playlistProto.Playlist
}

// --------------------------------------
func Check[V CheckConstraint](object V) error {
	return validator.Validate(object)
}

// --------------------------------------
func PathToAlbumCover(album *albumProto.Album, fileFormat string) (string, error) {
	return constants.AssetsPrefix + constants.AlbumPreName + fmt.Sprint(album.Id) + fileFormat, nil
}

// --------------------------------------
func CastAlbumToDtoWithoutArtistNameAndTracks(album *albumProto.Album, albumAgent domain.AlbumAgent, userId int64) (*albumProto.AlbumDataTransfer, error) {
	cover, err := PathToAlbumCover(album, constants.PngFormat)
	if err != nil {
		return nil, err
	}

	liked, err := albumAgent.LikeCheckByUser(userId, album.Id)
	if err != nil {
		liked = false
	}

	return &albumProto.AlbumDataTransfer{
		Id:       album.Id,
		Title:    album.Title,
		Artist:   "",
		ArtistId: album.ArtistId,
		Cover:    cover,
		Tracks:   nil,
		IsLiked:  liked,
	}, nil
}

// --------------------------------------
func PathToArtistCover(artist *artistProto.Artist, fileFormat string) (string, error) {
	return constants.AssetsPrefix + constants.ArtistPreName + fmt.Sprint(artist.Id) + fileFormat, nil
}

// --------------------------------------
func PathToTrackFile(track *trackProto.Track, fileFormat string) (string, error) {
	return constants.AssetsPrefix + constants.TrackPreName + fmt.Sprint(track.Id) + fileFormat, nil
}

// --------------------------------------
func PathToTrackFileByAlbumId(fileFormat string, albumId int64) (string, error) {
	return constants.AssetsPrefix + constants.AlbumPreName + fmt.Sprint(albumId) + fileFormat, nil
}

// --------------------------------------
func CastTrackToDtoWithoutArtistName(track *trackProto.Track, trackAgent domain.TrackAgent, userId int64) (*trackProto.TrackDataTransfer, error) {
	cover, err := PathToTrackFileByAlbumId(constants.PngFormat, track.AlbumId)
	if err != nil {
		return nil, err
	}

	src, err := PathToTrackFile(track, constants.Mp3Format)
	if err != nil {
		return nil, err
	}

	liked, err := trackAgent.LikeCheckByUser(userId, track.Id)
	if err != nil {
		liked = false
	}

	return &trackProto.TrackDataTransfer{
		Id:         track.Id,
		Title:      track.Title,
		Artist:     "",
		ArtistId:   track.ArtistId,
		AlbumId:    track.AlbumId,
		Cover:      cover,
		Src:        src,
		Likes:      track.CountLikes,
		Listenings: track.CountListenings,
		Duration:   track.Duration,
		IsLiked:    liked,
	}, nil
}

// --------------------------------------
func CastTrackToDto(track *trackProto.Track, artist *artistProto.Artist, trackAgent domain.TrackAgent, userId int64) (*trackProto.TrackDataTransfer, error) {
	trackDto, err := CastTrackToDtoWithoutArtistName(track, trackAgent, userId)
	if err != nil {
		return nil, err
	}
	trackDto.Artist = artist.Name
	return trackDto, nil
}

// --------------------------------------
func CastTracksByArtistToDto(userId int64, trackAgent domain.TrackAgent, tracks []*trackProto.Track, artist *artistProto.Artist) ([]*trackProto.TrackDataTransfer, error) {
	var err error
	tracksDto := make([]*trackProto.TrackDataTransfer, len(tracks))
	for idx, track := range tracks {
		tracksDto[idx], err = CastTrackToDto(track, artist, trackAgent, userId)
		if err != nil {
			return nil, err
		}
	}
	return tracksDto, nil
}

// --------------------------------------
func GetFullAlbumByArtist(userId int64, trackAgent domain.TrackAgent, albumAgent domain.AlbumAgent, album *albumProto.Album, artist *artistProto.Artist) (*albumProto.AlbumDataTransfer, error) {
	tracks, err := trackAgent.GetTracksFromAlbum(album.Id)
	if err != nil {
		return nil, err
	}

	tracksDto, err := CastTracksByArtistToDto(userId, trackAgent, tracks, artist)
	if err != nil {
		return nil, err
	}

	albumDto, err := CastAlbumToDtoWithoutArtistNameAndTracks(album, albumAgent, userId)
	if err != nil {
		return nil, err
	}

	albumDto.Artist = artist.Name
	albumDto.Tracks = tracksDto
	return albumDto, nil
}
