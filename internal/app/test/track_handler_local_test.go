package test

import (
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
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
//	tester.Get(t)
//}
//
//func TestLocalGetAllTrack(t *testing.T) {
//	Mutex.Lock()
//	defer Mutex.Unlock()
//	_, err := logger.InitLogrus("0", internal.Local)
//	require.NoError(t, err)
//	err = InitTestDb(internal.Track, internal.Local)
//	require.NoError(t, err)
//	tester.GetAll(t)
//}

func TestLocalCreateTrack(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Local)
	require.NoError(t, err)
	err = InitTestDb(internal.Track, internal.Local)
	require.NoError(t, err)
	creator := TrackTestCreator{}
	tester.Create(t, creator)
}

func TestLocalDeleteTrack(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Local)
	require.NoError(t, err)
	err = InitTestDb(internal.Track, internal.Local)
	require.NoError(t, err)
	const idToDelete = uint64(1)
	tester.Delete(t, idToDelete)
}

func TestLocalUpdateTrack(t *testing.T) {
	Mutex.Lock()
	defer Mutex.Unlock()
	_, err := logger.InitLogrus("0", internal.Local)
	require.NoError(t, err)
	err = InitTestDb(internal.Track, internal.Local)
	require.NoError(t, err)
	creator := TrackTestCreator{}
	tester.Update(t, creator)
}

//func TestLocalPopularTrack(t *testing.T) {
//	Mutex.Lock()
//	defer Mutex.Unlock()
//	_, err := logger.InitLogrus("0", internal.Local)
//	require.NoError(t, err)
//	err = InitTestDb(internal.Track, internal.Local)
//	require.NoError(t, err)
//	tester.GetPopular(t)
//}
