package test

import (
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	albumCoverUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/albumCover/usecase"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	structsUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/usecase"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPostgresGetAlbum(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Postgres)
	require.NoError(t, err)
	err = InitTestDb(internal.Album, internal.Postgres)
	require.NoError(t, err)
	tester.Get(t)
}

func TestPostgresGetAllAlbums(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Postgres)
	require.NoError(t, err)
	err = InitTestDb(internal.Album, internal.Postgres)
	require.NoError(t, err)
	tester.GetAll(t)
}

func TestPostgresCreateAlbum(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()

	_, err := logger.InitLogrus("0", internal.Postgres)
	require.NoError(t, err)
	err = InitTestDb(internal.Album, internal.Postgres)
	require.NoError(t, err)

	proxy, err := albumCoverUseCase.UseCase.Create(domain.AlbumCover{
		//Title:  "some new cover for new album",
		Quote:  "and quote for this",
		IsDark: true,
	})
	if err != nil {
		t.Fail()
	}

	albumCoverUseCase.UseCase = proxy.(structsUseCase.UseCase)

	creator := AlbumTestCreator{}
	tester.Create(t, creator)
}

func TestPostgresDeleteAlbum(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Postgres)
	require.NoError(t, err)
	err = InitTestDb(internal.Album, internal.Postgres)
	require.NoError(t, err)
	const idToDelete = uint64(1)
	tester.Delete(t, idToDelete)
}

func TestPostgresUpdateAlbum(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Postgres)
	require.NoError(t, err)
	err = InitTestDb(internal.Album, internal.Postgres)
	require.NoError(t, err)
	creator := AlbumTestCreator{}
	tester.Update(t, creator)
}

func TestPostgresPopularAlbums(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Postgres)
	require.NoError(t, err)
	err = InitTestDb(internal.Album, internal.Postgres)
	require.NoError(t, err)
	tester.GetPopular(t)
}
