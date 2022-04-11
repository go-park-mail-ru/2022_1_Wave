package test

import (
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	albumUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/album/usecase"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	structsUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/usecase"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPostgresGetAlbumCover(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Postgres)
	require.NoError(t, err)
	err = InitTestDb(internal.AlbumCover, internal.Postgres)
	require.NoError(t, err)
	tester.Get(t, domain.AlbumCoverMutex)
}

func TestPostgresGetAllAlbumCovers(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Postgres)
	require.NoError(t, err)
	err = InitTestDb(internal.AlbumCover, internal.Postgres)
	require.NoError(t, err)
	tester.GetAll(t, domain.AlbumCoverMutex)
}

func TestPostgresCreateAlbumCover(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Postgres)
	require.NoError(t, err)
	err = InitTestDb(internal.AlbumCover, internal.Postgres)
	require.NoError(t, err)
	creator := AlbumCoverTestCreator{}
	tester.Create(t, creator, domain.AlbumCoverMutex)
}

func TestPostgresDeleteAlbumCover(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Postgres)
	require.NoError(t, err)
	err = InitTestDb(internal.AlbumCover, internal.Postgres)
	require.NoError(t, err)

	const idToDelete = uint64(1)
	proxy, err := albumUseCase.UseCase.Delete(idToDelete, domain.AlbumMutex)
	if err != nil {
		t.Fail()
	}

	albumUseCase.UseCase = proxy.(structsUseCase.UseCase)

	tester.Delete(t, idToDelete, domain.AlbumCoverMutex)
}

func TestPostgresUpdateAlbumCover(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Postgres)
	require.NoError(t, err)
	err = InitTestDb(internal.AlbumCover, internal.Postgres)
	require.NoError(t, err)
	creator := AlbumCoverTestCreator{}
	tester.Update(t, creator, domain.AlbumCoverMutex)
}
