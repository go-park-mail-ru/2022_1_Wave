package test

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDomainAlbumCreateDataTransferFromInterface(t *testing.T) {

	const id = 1
	const title = "some title"
	const artist = "bla bla blashkin"
	const cover = "some path, that doesn't matter"
	dataTransfer := []domain.TrackDataTransfer{
		{
			Title:  "1",
			Artist: "1",
			//Cover:      "1",
			Src:        "1",
			Likes:      10,
			Listenings: 10,
			Duration:   10,
		},
		{
			Title:  "2",
			Artist: "2",
			//Cover:      "2",
			Src:        "2",
			Likes:      20,
			Listenings: 20,
			Duration:   20,
		},
	}

	data := map[string]interface{}{
		internal.FieldId:     float64(id),
		internal.FieldTitle:  title,
		internal.FieldArtist: artist,
		internal.FieldCover:  cover,
		internal.FieldTracks: dataTransfer,
	}

	except := domain.AlbumDataTransfer{
		Id:     id,
		Title:  title,
		Artist: artist,
		Cover:  cover,
		Tracks: dataTransfer,
	}

	actual, err := domain.AlbumDataTransfer{}.CreateDataTransferFromInterface(data)
	require.NoError(t, err)
	require.Equal(t, except, actual)

}
