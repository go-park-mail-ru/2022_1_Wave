package test

import (
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPostgresGetArtist(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Postgres)
	require.NoError(t, err)
	err = InitTestDb(internal.Artist, internal.Postgres)
	require.NoError(t, err)
	tester.Get(t)
}

func TestPostgresGetAllArtists(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Postgres)
	require.NoError(t, err)
	err = InitTestDb(internal.Artist, internal.Postgres)
	require.NoError(t, err)
	tester.GetAll(t)
}

func TestPostgresCreateArtist(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Postgres)
	require.NoError(t, err)
	err = InitTestDb(internal.Artist, internal.Postgres)
	require.NoError(t, err)
	creator := ArtistTestCreator{}
	tester.Create(t, creator)
}

func TestPostgresDeleteArtist(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Postgres)
	require.NoError(t, err)
	err = InitTestDb(internal.Artist, internal.Postgres)
	require.NoError(t, err)
	const idToDelete = uint64(1)
	tester.Delete(t, idToDelete)
}

func TestPostgresUpdateArtist(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Postgres)
	require.NoError(t, err)
	err = InitTestDb(internal.Artist, internal.Postgres)
	require.NoError(t, err)
	creator := ArtistTestCreator{}
	tester.Update(t, creator)
}

func TestPostgresPopularArtists(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Postgres)
	require.NoError(t, err)
	err = InitTestDb(internal.Artist, internal.Postgres)
	require.NoError(t, err)
	tester.GetPopular(t)
}
