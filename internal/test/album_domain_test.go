package test

import (
	"fmt"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAlbum(t *testing.T) {
	album := domain.Album{
		Id:             500,
		Title:          "asd",
		ArtistId:       10,
		CountLikes:     110,
		CountListening: 1110,
		Date:           10,
	}

	require.Equal(t, album.Id, album.GetId())
	require.NoError(t, album.Check())
	require.Equal(t, album.CountListening, album.GetCountListening())
	format := "aboba"
	expected := constants.AssetsPrefix + constants.AlbumPreName + fmt.Sprint(album.Id) + format
	actual, err := album.CreatePath(format)
	require.NoError(t, err)
	require.Equal(t, expected, actual)
	require.NoError(t, album.SetId(1000))
	require.Equal(t, 1000, album.Id)
}
