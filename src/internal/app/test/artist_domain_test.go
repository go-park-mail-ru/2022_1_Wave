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

	data := map[string]interface{}{
		internal.FieldName:  name,
		internal.FieldCover: cover,
	}

	except := domain.ArtistDataTransfer{
		Name:  name,
		Cover: cover,
	}

	actual, err := domain.ArtistDataTransfer{}.CreateDataTransferFromInterface(data)
	require.NoError(t, err)
	require.Equal(t, except, actual)

}
