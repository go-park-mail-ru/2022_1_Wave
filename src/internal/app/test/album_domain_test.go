package test

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDomainAlbumCreateDataTransferFromInterface(t *testing.T) {

	const title = "some title"
	const artist = "bla bla blashkin"
	const cover = "some path, that doesn't matter"

	data := map[string]interface{}{
		internal.FieldTitle:  title,
		internal.FieldArtist: artist,
		internal.FieldCover:  cover,
	}

	except := domain.AlbumDataTransfer{
		Title:  title,
		Artist: artist,
		Cover:  cover,
	}

	actual, err := domain.AlbumDataTransfer{}.CreateDataTransferFromInterface(data)
	require.NoError(t, err)
	require.Equal(t, except, actual)

}
