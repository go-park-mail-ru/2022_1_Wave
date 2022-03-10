package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-park-mail-ru/2022_1_Wave/db"
	"github.com/go-park-mail-ru/2022_1_Wave/db/models"
	"github.com/go-park-mail-ru/2022_1_Wave/db/views"
	"github.com/go-park-mail-ru/2022_1_Wave/pkg/status"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"
)

// ----------------------------------------------------------------------
type getTrackTestCase struct {
	id     uint64
	track  views.Track
	status int
}

func TestGetTrack(t *testing.T) {
	const testDataBaseSize = 20
	db.Storage.InitStorage(testDataBaseSize)
	storage := db.Storage.TrackStorage.Tracks

	cases := make([]getTrackTestCase, len(storage))
	for idx, track := range storage {
		cases[idx] = getTrackTestCase{
			id:     track.Id,
			track:  *GetTrackView(track.Id),
			status: http.StatusOK,
		}
	}

	for caseNum, item := range cases {
		url := Proto + Host + GetTrackUrlWithoutId + fmt.Sprint(item.id)
		req := httptest.NewRequest("GET", url, nil)

		w := httptest.NewRecorder()

		vars := map[string]string{
			"id": fmt.Sprint(item.id),
		}
		req = mux.SetURLVars(req, vars)
		GetTrack(w, req)

		if w.Code != item.status {
			t.Errorf("[%d] wrong StatusCode: got %d, expected %d",
				caseNum, w.Code, item.status)
		}

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		var response status.Success
		json.Unmarshal(body, &response)
		result := views.FromInterfaceToTrackView(response.Result)
		if result != item.track {
			t.Errorf("[%d] wrong Response: got %+v, expected %+v",
				caseNum, resp, item.track)
		}
	}
}

// ----------------------------------------------------------------------
type allTrackTestCase struct {
	tracks []views.Track
	status int
}

func TestGetTracks(t *testing.T) {
	const testDataBaseSize = 20
	db.Storage.InitStorage(testDataBaseSize)
	storage := db.Storage.TrackStorage.Tracks

	tracksViews := make([]views.Track, len(storage))
	for idx, track := range storage {
		tracksViews[idx] = *GetTrackView(track.Id)
	}

	url := Proto + Host + GetAllTracksUrl
	req := httptest.NewRequest("GET", url, nil)

	w := httptest.NewRecorder()

	GetTracks(w, req)

	testCase := allTrackTestCase{
		tracks: tracksViews,
		status: http.StatusOK,
	}

	if w.Code != testCase.status {
		t.Errorf("wrong StatusCode: got %d, expected %d", w.Code, testCase.status)
	}

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	var result status.Success

	json.Unmarshal(body, &result)
	data := result.Result.([]interface{})
	tracks := views.GetTracksViewsFromInterfaces(data)
	for idx, view := range testCase.tracks {
		if tracks[idx] != view {
			t.Errorf("wrong Response: got %+v, expected %+v", tracks, view)
		}
	}
}

// ----------------------------------------------------------------------
type createTrackTestCase struct {
	status int
}

func TestCreateTrack(t *testing.T) {
	const testDataBaseSize = 20
	db.Storage.InitStorage(testDataBaseSize)

	track := models.Track{
		Id:             0,
		AlbumId:        5,
		AuthorId:       8,
		Title:          "what are you want?",
		Duration:       0,
		Mp4:            "some source",
		CoverId:        50,
		CountLikes:     500,
		CountListening: 5000,
	}

	url := Proto + Host + GetAllTracksUrl
	dataToSend, _ := json.Marshal(track)
	req := httptest.NewRequest("POST", url, bytes.NewBuffer(dataToSend))

	w := httptest.NewRecorder()
	CreateTrack(w, req)

	testCase := createTrackTestCase{
		status: http.StatusOK,
	}

	if w.Code != testCase.status {
		t.Errorf("wrong StatusCode: got %d, expected %d", w.Code, testCase.status)
	}

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	var result status.Success
	json.Unmarshal(body, &result)
	expected := db.SuccessWrapper(track.Title, db.SuccessCreatedTrack)
	if result != status.MakeSuccess(expected) {
		t.Errorf("wrong Response: got %+v, expected %+v",
			result, expected)
	}
}

// ----------------------------------------------------------------------
type deleteTrackTestCase struct {
	status int
}

func TestDeleteTrack(t *testing.T) {
	const testDataBaseSize = 20
	db.Storage.InitStorage(testDataBaseSize)

	sizeBefore := len(db.Storage.TrackStorage.Tracks)
	sizeAfter := sizeBefore - 1

	trackToDelete := db.Storage.TrackStorage.Tracks[0]

	url := Proto + Host + DeleteTrackUrl
	req := httptest.NewRequest("DELETE", url, nil)

	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": fmt.Sprint(trackToDelete.Id),
	}
	req = mux.SetURLVars(req, vars)
	DeleteTrack(w, req)

	testCase := deleteTrackTestCase{
		status: http.StatusOK,
	}

	if w.Code != testCase.status {
		t.Errorf("wrong StatusCode: got %d, expected %d", w.Code, testCase.status)
	}

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	var result status.Success
	json.Unmarshal(body, &result)
	expected := db.SuccessWrapper(trackToDelete.Id, db.SuccessDeletedTrack)
	if result != status.MakeSuccess(expected) {
		t.Errorf("wrong Response: got %+v, expected %+v",
			result, expected)
	}
	require.Equal(t, sizeAfter, len(db.Storage.TrackStorage.Tracks))
}

// ----------------------------------------------------------------------
type updateTrackTestCase struct {
	status int
}

func TestUpdateTrack(t *testing.T) {
	const testDataBaseSize = 20
	db.Storage.InitStorage(testDataBaseSize)

	track := models.Track{
		Id:             0,
		AlbumId:        5,
		AuthorId:       8,
		Title:          "what are you want?",
		Duration:       0,
		Mp4:            "some source",
		CoverId:        50,
		CountLikes:     500,
		CountListening: 5000,
	}

	url := Proto + Host + UpdateTrackUrl
	dataToSend, _ := json.Marshal(track)
	req := httptest.NewRequest("PUT", url, bytes.NewBuffer(dataToSend))

	w := httptest.NewRecorder()
	UpdateTrack(w, req)

	testCase := updateTrackTestCase{
		status: http.StatusOK,
	}

	if w.Code != testCase.status {
		t.Errorf("wrong StatusCode: got %d, expected %d", w.Code, testCase.status)
	}

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	var result status.Success
	json.Unmarshal(body, &result)
	expected := db.SuccessWrapper(track.Id, db.SuccessUpdatedTrack)
	if result != status.MakeSuccess(expected) {
		t.Errorf("wrong Response: got %+v, expected %+v",
			result, expected)
	}
	require.Equal(t, track.Title, db.Storage.TrackStorage.Tracks[0].Title)
}

// ----------------------------------------------------------------------
type popularTrackTestCase struct {
	tracks []views.Track
	status int
}

func TestPopularTracks(t *testing.T) {
	const testDataBaseSize = 20
	db.Storage.InitStorage(testDataBaseSize)
	storage := db.Storage.TrackStorage.Tracks

	sort.SliceStable(storage, func(i int, j int) bool {
		return storage[i].CountListening > storage[j].CountListening
	})

	tracksViews := make([]views.Track, len(storage))
	for idx, _ := range storage {
		tracksViews[idx] = *GetTrackView(uint64(idx))
	}

	url := Proto + Host + GetPopularTracksUrl
	req := httptest.NewRequest("GET", url, nil)

	w := httptest.NewRecorder()

	GetPopularTracks(w, req)

	testCase := popularTrackTestCase{
		tracks: tracksViews,
		status: http.StatusOK,
	}

	if w.Code != testCase.status {
		t.Errorf("wrong StatusCode: got %d, expected %d", w.Code, testCase.status)
	}

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	var result status.Success

	json.Unmarshal(body, &result)

	data := result.Result.([]interface{})
	tracks := views.GetTracksViewsFromInterfaces(data)
	for idx, track := range testCase.tracks {
		if tracks[idx] != track {
			t.Errorf("wrong Response: got %+v, expected %+v", tracks, track)
		}
	}
}
