package test

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDomainArtistCreateDataTransferFromInterface(t *testing.T) {

	const name = "bla bla blashkin"
	const cover = "some path, that doesn't matter"
	dataTransfer := []domain.AlbumDataTransfer{
		{
			Title:  "1",
			Artist: "1",
			Cover:  "1",
			Tracks: nil,
		},
		{
			Title:  "2",
			Artist: "2",
			Cover:  "2",
			Tracks: nil,
		},
	}

	data := map[string]interface{}{
		internal.FieldName:   name,
		internal.FieldCover:  cover,
		internal.FieldAlbums: dataTransfer,
	}

	except := domain.ArtistDataTransfer{
		Name:   name,
		Cover:  cover,
		Albums: dataTransfer,
	}

	actual, err := domain.ArtistDataTransfer{}.CreateDataTransferFromInterface(data)
	require.NoError(t, err)
	require.Equal(t, except, actual)

}
