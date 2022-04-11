package test

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDomainAlbumCoverCreateDataTransferFromInterface(t *testing.T) {

	const title = "some title"
	const quote = "bla bla blashkin and more phrases here"
	const isDark = false

	data := map[string]interface{}{
		internal.FieldTitle:  title,
		internal.FieldQuote:  quote,
		internal.FieldIsDark: isDark,
	}

	except := domain.AlbumCoverDataTransfer{
		Title:  title,
		Quote:  quote,
		IsDark: isDark,
	}

	actual, err := domain.AlbumCoverDataTransfer{}.CreateDataTransferFromInterface(data)
	require.NoError(t, err)
	require.Equal(t, except, actual)

}
