package AlbumPostgres

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
)

func GetFullAlbumByArtist(trackRepo domain.TrackRepo, album domain.Album, artist domain.Artist) (*domain.AlbumDataTransfer, error) {
	tracks, err := trackRepo.GetTracksFromAlbum(album.Id)
	if err != nil {
		return nil, err
	}

	tracksDto, err := domain.CastTracksByArtistToDto(tracks, artist)
	if err != nil {
		return nil, err
	}

	albumDto, err := album.CastToDtoWithoutArtistNameAndTracks()
	if err != nil {
		return nil, err
	}

	albumDto.Artist = artist.Name
	albumDto.Tracks = tracksDto
	return albumDto, nil
}
