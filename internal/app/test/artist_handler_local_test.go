package test

import (
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/stretchr/testify/require"
	"testing"
)

//func TestLocalGetArtist(t *testing.T) {
//	Mutex.Lock()
//	defer Mutex.Unlock()
//	_, err := logger.InitLogrus("0", internal.Local)
//	require.NoError(t, err)
//	err = InitTestDb(internal.Artist, internal.Local)
//	require.NoError(t, err)
//	tester.Get(t, domain.ArtistMutex)
//}
//
//func TestLocalGetAllArtist(t *testing.T) {
//	Mutex.Lock()
//	defer Mutex.Unlock()
//	_, err := logger.InitLogrus("0", internal.Local)
//	require.NoError(t, err)
//	err = InitTestDb(internal.Artist, internal.Local)
//	require.NoError(t, err)
//	tester.GetAll(t, domain.ArtistMutex)
//}

func TestLocalCreateArtist(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Local)
	require.NoError(t, err)
	err = InitTestDb(internal.Artist, internal.Local)
	require.NoError(t, err)
	creator := ArtistTestCreator{}
	tester.Create(t, creator, domain.ArtistMutex)
}

func TestLocalDeleteArtist(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Local)
	require.NoError(t, err)
	err = InitTestDb(internal.Artist, internal.Local)
	require.NoError(t, err)
	const idToDelete = uint64(1)
	tester.Delete(t, idToDelete, domain.ArtistMutex)
}

func TestLocalUpdateArtist(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Local)
	require.NoError(t, err)
	err = InitTestDb(internal.Artist, internal.Local)
	require.NoError(t, err)
	creator := ArtistTestCreator{}
	tester.Update(t, creator, domain.ArtistMutex)
}

//func TestLocalPopularArtist(t *testing.T) {
//	Mutex.Lock()
//	defer Mutex.Unlock()
//	_, err := logger.InitLogrus("0", internal.Local)
//	require.NoError(t, err)
//	err = InitTestDb(internal.Artist, internal.Local)
//	require.NoError(t, err)
//	tester.GetPopular(t, domain.ArtistMutex)
//}
