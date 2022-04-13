package test

import (
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPostgresGetTrack(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Postgres)
	require.NoError(t, err)
	err = InitTestDb(internal.Track, internal.Postgres)
	require.NoError(t, err)
	tester.Get(t)
}

func TestPostgresGetAllTracks(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Postgres)
	require.NoError(t, err)
	err = InitTestDb(internal.Track, internal.Postgres)
	require.NoError(t, err)
	tester.GetAll(t)
}

func TestPostgresCreateTrack(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Postgres)
	require.NoError(t, err)
	err = InitTestDb(internal.Track, internal.Postgres)
	require.NoError(t, err)
	creator := TrackTestCreator{}
	tester.Create(t, creator)
}

func TestPostgresDeleteTrack(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Postgres)
	require.NoError(t, err)
	err = InitTestDb(internal.Track, internal.Postgres)
	require.NoError(t, err)
	const idToDelete = uint64(1)
	tester.Delete(t, idToDelete)
}

func TestPostgresUpdateTrack(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Postgres)
	require.NoError(t, err)
	err = InitTestDb(internal.Track, internal.Postgres)
	require.NoError(t, err)
	creator := TrackTestCreator{}
	tester.Update(t, creator)
}

func TestPostgresPopularTracks(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Postgres)
	require.NoError(t, err)
	err = InitTestDb(internal.Track, internal.Postgres)
	require.NoError(t, err)
	tester.GetPopular(t)
}
