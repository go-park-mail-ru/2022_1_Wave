package test

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDomainArtistCreateDataTransferFromInterface(t *testing.T) {

	const id = float64(1)
	const name = "bla bla blashkin"
	const cover = "some path, that doesn't matter"
	const likes = 500

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
		internal.FieldId:     id,
		internal.FieldName:   name,
		internal.FieldCover:  cover,
		internal.FieldAlbums: dataTransfer,
		internal.FieldLikes:  float64(likes),
	}

	except := domain.ArtistDataTransfer{
		Id:     uint64(id),
		Name:   name,
		Cover:  cover,
		Albums: dataTransfer,
		Likes:  uint64(likes),
	}

	actual, err := domain.ArtistDataTransfer{}.CreateDataTransferFromInterface(data)
	require.NoError(t, err)
	require.Equal(t, except, actual)

}
