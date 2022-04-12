package test

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDomainTrackCreateDataTransferFromInterface(t *testing.T) {

	const title = "some title"
	const artist = "bla bla blashkin"
	const cover = "some path, that doesn't matter"
	const src = "some src"
	const likes = 500
	const listenings = 5000
	const duration = 500

	data := map[string]interface{}{
		internal.FieldTitle:      title,
		internal.FieldArtist:     artist,
		internal.FieldCover:      cover,
		internal.FieldSrc:        src,
		internal.FieldLikes:      float64(likes),
		internal.FieldListenings: float64(listenings),
		internal.FieldDuration:   float64(duration),
	}

	except := domain.TrackDataTransfer{
		Title:  title,
		Artist: artist,
		//Cover:      cover,
		Src:        src,
		Likes:      likes,
		Listenings: listenings,
		Duration:   duration,
	}

	actual, err := domain.TrackDataTransfer{}.CreateDataTransferFromInterface(data)
	require.NoError(t, err)
	require.Equal(t, except, actual)

}
