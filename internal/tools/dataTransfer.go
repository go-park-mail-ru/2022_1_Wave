package tools

import (
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/album/albumProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/artist/artistProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/trackProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/tools/utils"
)

func CreateAlbumDataTransferFromInterface(data interface{}) (*albumProto.AlbumDataTransfer, error) {
	temp := data.(map[string]interface{})

	id, err := utils.ToInt64(temp[constants.FieldId])
	if err != nil {
		return nil, err
	}

	title, err := utils.ToString(temp[constants.FieldTitle])
	if err != nil {
		return nil, err
	}

	artist, err := utils.ToString(temp[constants.FieldArtist])
	if err != nil {
		return nil, err
	}

	cover, err := utils.ToString(temp[constants.FieldCover])
	if err != nil {
		return nil, err
	}

	tracksInterfaces := temp[constants.FieldTracks].([]interface{})
	tracks := make([]*trackProto.TrackDataTransfer, len(tracksInterfaces))
	for i, obj := range tracksInterfaces {
		temp := obj.(map[string]interface{})
		tracks[i], err = CreateTrackDataTransferFromInterface(temp)
		if err != nil {
			return nil, err
		}
	}

	return &albumProto.AlbumDataTransfer{
		Id:     id,
		Title:  title,
		Artist: artist,
		Cover:  cover,
		Tracks: tracks,
	}, nil
}

// ---------------------------------------------------------
func CreateAlbumCoverDataTransferFromInterface(data interface{}) (*albumProto.AlbumCoverDataTransfer, error) {
	temp := data.(map[string]interface{})

	quote, err := utils.ToString(temp[constants.FieldQuote])
	if err != nil {
		return nil, err
	}

	isDark, err := utils.ToBool(temp[constants.FieldIsDark])
	if err != nil {
		return nil, err
	}

	return &albumProto.AlbumCoverDataTransfer{
		//Title:  title,
		Quote:  quote,
		IsDark: isDark,
	}, nil
}

// ---------------------------------------------------------
func CreateArtistDataTransferFromInterface(data interface{}) (*artistProto.ArtistDataTransfer, error) {
	temp := data.(map[string]interface{})

	id, err := utils.ToInt64(temp[constants.FieldId])
	if err != nil {
		return nil, err
	}

	name, err := utils.ToString(temp[constants.FieldName])
	if err != nil {
		return nil, err
	}

	cover, err := utils.ToString(temp[constants.FieldCover])
	if err != nil {
		return nil, err
	}

	likes, err := utils.ToInt64(temp[constants.FieldLikes])
	if err != nil {
		return nil, err
	}

	albumsArray := temp[constants.FieldAlbums].([]interface{})
	albums := make([]*albumProto.AlbumDataTransfer, len(albumsArray))
	for i, obj := range albumsArray {
		temp := obj.(map[string]interface{})
		track, err := CreateAlbumDataTransferFromInterface(temp)
		if err != nil {
			return nil, err
		}
		albums[i] = track
	}

	return &artistProto.ArtistDataTransfer{
		Id:     id,
		Name:   name,
		Cover:  cover,
		Albums: albums,
		Likes:  likes,
	}, nil
}

// ---------------------------------------------------------
func CreateTrackDataTransferFromInterface(data interface{}) (*trackProto.TrackDataTransfer, error) {
	temp := data.(map[string]interface{})

	id, err := utils.ToInt64(temp[constants.FieldId])
	if err != nil {
		return nil, err
	}

	title, err := utils.ToString(temp[constants.FieldTitle])
	if err != nil {
		return nil, err
	}

	artist, err := utils.ToString(temp[constants.FieldArtist])
	if err != nil {
		return nil, err
	}

	cover, err := utils.ToString(temp[constants.FieldCover])
	if err != nil {
		return nil, err
	}

	src, err := utils.ToString(temp[constants.FieldSrc])
	if err != nil {
		return nil, err
	}

	likes, err := utils.ToInt64(temp[constants.FieldLikes])
	if err != nil {
		return nil, err
	}

	listenings, err := utils.ToInt64(temp[constants.FieldListenings])
	if err != nil {
		return nil, err
	}

	duration, err := utils.ToInt64(temp[constants.FieldDuration])
	if err != nil {
		return nil, err
	}

	return &trackProto.TrackDataTransfer{
		Id:         id,
		Title:      title,
		Artist:     artist,
		Cover:      cover,
		Src:        src,
		Likes:      likes,
		Listenings: listenings,
		Duration:   duration,
	}, nil
}
