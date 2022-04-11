package test

import (
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLocalGetAlbumCover(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Local)
	require.NoError(t, err)
	err = InitTestDb(internal.AlbumCover, internal.Local)
	require.NoError(t, err)
	tester.Get(t, domain.AlbumCoverMutex)
}

func TestLocalGetAllAlbumCover(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Local)
	require.NoError(t, err)
	err = InitTestDb(internal.AlbumCover, internal.Local)
	require.NoError(t, err)
	tester.GetAll(t, domain.AlbumCoverMutex)
}

func TestLocalCreateAlbumCover(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Local)
	require.NoError(t, err)
	err = InitTestDb(internal.AlbumCover, internal.Local)
	require.NoError(t, err)
	creator := AlbumCoverTestCreator{}
	tester.Create(t, creator, domain.AlbumCoverMutex)
}

func TestLocalDeleteAlbumCover(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Local)
	require.NoError(t, err)
	err = InitTestDb(internal.AlbumCover, internal.Local)
	require.NoError(t, err)
	const idToDelete = uint64(1)
	tester.Delete(t, idToDelete, domain.AlbumCoverMutex)
}

func TestLocalUpdateAlbumCover(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Local)
	require.NoError(t, err)
	err = InitTestDb(internal.AlbumCover, internal.Local)
	require.NoError(t, err)
	creator := AlbumCoverTestCreator{}
	tester.Update(t, creator, domain.AlbumCoverMutex)
}
