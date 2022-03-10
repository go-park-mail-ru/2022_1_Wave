package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-park-mail-ru/2022_1_Wave/db"
	"github.com/go-park-mail-ru/2022_1_Wave/db/models"
	"github.com/go-park-mail-ru/2022_1_Wave/db/views"
	"github.com/go-park-mail-ru/2022_1_Wave/pkg/utils"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"
)

// ----------------------------------------------------------------------
type getArtistTestCase struct {
	id     uint64
	artist views.Artist
	status int
}

func TestGetArtist(t *testing.T) {
	const testDataBaseSize = 20
	db.Storage.InitStorage(testDataBaseSize)
	storage := db.Storage.ArtistStorage.Artists

	cases := make([]getArtistTestCase, len(storage))
	for idx, artist := range storage {
		cases[idx] = getArtistTestCase{
			id:     artist.Id,
			artist: *GetArtistView(artist.Id),
			status: http.StatusOK,
		}
	}

	for caseNum, item := range cases {
		url := Proto + Host + GetArtistUrlWithoutId + fmt.Sprint(item.id)
		req := httptest.NewRequest("GET", url, nil)

		w := httptest.NewRecorder()

		vars := map[string]string{
			"id": fmt.Sprint(item.id),
		}
		req = mux.SetURLVars(req, vars)
		GetArtist(w, req)

		if w.Code != item.status {
			t.Errorf("[%d] wrong StatusCode: got %d, expected %d",
				caseNum, w.Code, item.status)
		}

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		var result views.Artist
		json.Unmarshal(body, &result)
		if result != item.artist {
			t.Errorf("[%d] wrong Response: got %+v, expected %+v",
				caseNum, resp, item.artist)
		}
	}
}

// ----------------------------------------------------------------------
type allArtistTestCase struct {
	artists []views.Artist
	status  int
}

func TestGetArtists(t *testing.T) {
	const testDataBaseSize = 20
	db.Storage.InitStorage(testDataBaseSize)
	storage := db.Storage.ArtistStorage.Artists

	artistsViews := make([]views.Artist, len(storage))
	for idx, artist := range storage {
		artistsViews[idx] = *GetArtistView(artist.Id)
	}

	url := Proto + Host + GetAllArtistsUrl
	req := httptest.NewRequest("GET", url, nil)

	w := httptest.NewRecorder()

	GetArtists(w, req)

	testCase := allArtistTestCase{
		artists: artistsViews,
		status:  http.StatusOK,
	}

	if w.Code != testCase.status {
		t.Errorf("wrong StatusCode: got %d, expected %d", w.Code, testCase.status)
	}

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	var result utils.Success

	json.Unmarshal(body, &result)

	data := result.Result.([]interface{})
	albums := views.SetArtistsViewsFromInterfaces(data)
	for idx, view := range testCase.artists {
		if albums[idx] != view {
			t.Errorf("wrong Response: got %+v, expected %+v", albums, view)
		}
	}
}

// ----------------------------------------------------------------------
type createArtistTestCase struct {
	status int
}

func TestCreateArtist(t *testing.T) {
	const testDataBaseSize = 20
	db.Storage.InitStorage(testDataBaseSize)

	artist := models.Artist{
		Id:             0,
		Name:           "monster gammy",
		Photo:          "some_photo",
		CountFollowers: 500,
		CountListening: 5000,
	}

	url := Proto + Host + GetAllArtistsUrl
	dataToSend, _ := json.Marshal(artist)
	req := httptest.NewRequest("POST", url, bytes.NewBuffer(dataToSend))

	w := httptest.NewRecorder()
	CreateArtist(w, req)

	testCase := createArtistTestCase{
		status: http.StatusOK,
	}

	if w.Code != testCase.status {
		t.Errorf("wrong StatusCode: got %d, expected %d", w.Code, testCase.status)
	}

	expected := utils.Success{
		Result: db.SuccessCreatedArtist + "(" + artist.Name + ")",
	}
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	var result utils.Success
	json.Unmarshal(body, &result)
	if result != expected {
		t.Errorf("wrong Response: got %+v, expected %+v",
			result, expected)
	}
}

// ----------------------------------------------------------------------
type deleteArtistsTestCase struct {
	status int
}

func TestDeleteArtist(t *testing.T) {
	const testDataBaseSize = 20
	db.Storage.InitStorage(testDataBaseSize)

	sizeBefore := len(db.Storage.ArtistStorage.Artists)
	sizeAfter := sizeBefore - 1

	artistToDelete := db.Storage.ArtistStorage.Artists[0]

	url := Proto + Host + DeleteArtistUrl
	req := httptest.NewRequest("DELETE", url, nil)

	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": fmt.Sprint(artistToDelete.Id),
	}
	req = mux.SetURLVars(req, vars)
	DeleteArtist(w, req)

	testCase := deleteArtistsTestCase{
		status: http.StatusOK,
	}

	if w.Code != testCase.status {
		t.Errorf("wrong StatusCode: got %d, expected %d", w.Code, testCase.status)
	}

	expected := utils.Success{
		Result: db.SuccessDeletedArtist + "(" + fmt.Sprint(artistToDelete.Id) + ")",
	}
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	var result utils.Success
	json.Unmarshal(body, &result)
	if result != expected {
		t.Errorf("wrong Response: got %+v, expected %+v",
			result, expected)
	}
	require.Equal(t, sizeAfter, len(db.Storage.ArtistStorage.Artists))
}

// ----------------------------------------------------------------------
type updateArtistTestCase struct {
	status int
}

func TestUpdateArtist(t *testing.T) {
	const testDataBaseSize = 20
	db.Storage.InitStorage(testDataBaseSize)

	artist := models.Artist{
		Id:             0,
		Name:           "monster gammy",
		Photo:          "some_photo",
		CountFollowers: 500,
		CountListening: 5000,
	}

	url := Proto + Host + UpdateArtistUrl
	dataToSend, _ := json.Marshal(artist)
	req := httptest.NewRequest("PUT", url, bytes.NewBuffer(dataToSend))

	w := httptest.NewRecorder()
	UpdateArtist(w, req)

	testCase := updateArtistTestCase{
		status: http.StatusOK,
	}

	if w.Code != testCase.status {
		t.Errorf("wrong StatusCode: got %d, expected %d", w.Code, testCase.status)
	}

	expected := utils.Success{
		Result: db.SuccessUpdatedArtist + "(" + fmt.Sprint(artist.Id) + ")",
	}
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	var result utils.Success
	json.Unmarshal(body, &result)
	if result != expected {
		t.Errorf("wrong Response: got %+v, expected %+v",
			result, expected)
	}
	require.Equal(t, artist.Name, db.Storage.ArtistStorage.Artists[0].Name)
}

// ----------------------------------------------------------------------
type popularArtistTestCase struct {
	artists []views.Artist
	status  int
}

func TestPopularArtists(t *testing.T) {
	const testDataBaseSize = 20
	db.Storage.InitStorage(testDataBaseSize)
	storage := db.Storage.ArtistStorage.Artists

	sort.SliceStable(storage, func(i int, j int) bool {
		return storage[i].CountListening > storage[j].CountListening
	})

	artistsViews := make([]views.Artist, len(storage))
	for idx, _ := range storage {
		artistsViews[idx] = *GetArtistView(uint64(idx))
	}

	url := Proto + Host + GetPopularArtistsUrl
	req := httptest.NewRequest("GET", url, nil)

	w := httptest.NewRecorder()

	GetPopularArtists(w, req)

	testCase := popularArtistTestCase{
		artists: artistsViews,
		status:  http.StatusOK,
	}

	if w.Code != testCase.status {
		t.Errorf("wrong StatusCode: got %d, expected %d", w.Code, testCase.status)
	}

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	var result utils.Success

	json.Unmarshal(body, &result)

	data := result.Result.([]interface{})
	albums := views.SetArtistsViewsFromInterfaces(data)
	for idx, artist := range testCase.artists {
		if albums[idx] != artist {
			t.Errorf("wrong Response: got %+v, expected %+v", albums, artist)
		}
	}
}
