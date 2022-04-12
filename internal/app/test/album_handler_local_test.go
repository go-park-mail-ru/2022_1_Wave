package test

import (
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/stretchr/testify/require"
	"testing"
)

//func TestLocalGetAlbum(t *testing.T) {
//	Mutex.Lock()
//	defer Mutex.Unlock()
//	_, err := logger.InitLogrus("0", internal.Local)
//	require.NoError(t, err)
//	err = InitTestDb(internal.Album, internal.Local)
//	require.NoError(t, err)
//	tester.Get(t, domain.AlbumMutex)
//}

//func TestLocalGetAllAlbum(t *testing.T) {
//	Mutex.Lock()
//	defer Mutex.Unlock()
//	_, err := logger.InitLogrus("0", internal.Local)
//	require.NoError(t, err)
//	err = InitTestDb(internal.Album, internal.Local)
//	require.NoError(t, err)
//	tester.GetAll(t, domain.AlbumMutex)
//}

func TestLocalCreateAlbum(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Local)
	require.NoError(t, err)
	err = InitTestDb(internal.Album, internal.Local)
	require.NoError(t, err)
	creator := AlbumTestCreator{}
	tester.Create(t, creator, domain.AlbumMutex)
}

func TestLocalDeleteAlbum(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Local)
	require.NoError(t, err)
	err = InitTestDb(internal.Album, internal.Local)
	require.NoError(t, err)
	const idToDelete = uint64(1)
	tester.Delete(t, idToDelete, domain.AlbumMutex)
}

func TestLocalUpdateAlbum(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Local)
	require.NoError(t, err)
	err = InitTestDb(internal.Album, internal.Local)
	require.NoError(t, err)
	creator := AlbumTestCreator{}
	tester.Update(t, creator, domain.AlbumMutex)
}

//func TestLocalPopularAlbum(t *testing.T) {
//	Mutex.Lock()
//	defer Mutex.Unlock()
//	_, err := logger.InitLogrus("0", internal.Local)
//	require.NoError(t, err)
//	err = InitTestDb(internal.Album, internal.Local)
//	require.NoError(t, err)
//	tester.GetPopular(t, domain.AlbumMutex)
//}
