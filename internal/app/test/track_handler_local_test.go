package test

import (
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/stretchr/testify/require"
	"testing"
)

//func TestLocalGetTrack(t *testing.T) {
//	Mutex.Lock()
//	defer Mutex.Unlock()
//	_, err := logger.InitLogrus("0", internal.Local)
//	require.NoError(t, err)
//	err = InitTestDb(internal.Track, internal.Local)
//	require.NoError(t, err)
//	tester.Get(t, domain.TrackMutex)
//}
//
//func TestLocalGetAllTrack(t *testing.T) {
//	Mutex.Lock()
//	defer Mutex.Unlock()
//	_, err := logger.InitLogrus("0", internal.Local)
//	require.NoError(t, err)
//	err = InitTestDb(internal.Track, internal.Local)
//	require.NoError(t, err)
//	tester.GetAll(t, domain.TrackMutex)
//}

func TestLocalCreateTrack(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Local)
	require.NoError(t, err)
	err = InitTestDb(internal.Track, internal.Local)
	require.NoError(t, err)
	creator := TrackTestCreator{}
	tester.Create(t, creator, domain.TrackMutex)
}

func TestLocalDeleteTrack(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Local)
	require.NoError(t, err)
	err = InitTestDb(internal.Track, internal.Local)
	require.NoError(t, err)
	const idToDelete = uint64(1)
	tester.Delete(t, idToDelete, domain.TrackMutex)
}

func TestLocalUpdateTrack(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Local)
	require.NoError(t, err)
	err = InitTestDb(internal.Track, internal.Local)
	require.NoError(t, err)
	creator := TrackTestCreator{}
	tester.Update(t, creator, domain.TrackMutex)
}

//func TestLocalPopularTrack(t *testing.T) {
//	Mutex.Lock()
//	defer Mutex.Unlock()
//	_, err := logger.InitLogrus("0", internal.Local)
//	require.NoError(t, err)
//	err = InitTestDb(internal.Track, internal.Local)
//	require.NoError(t, err)
//	tester.GetPopular(t, domain.TrackMutex)
//}
