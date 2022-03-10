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
type getAlbumTestCase struct {
	id     uint64
	album  views.Album
	status int
}

//func TestGetAlbumsServer(t *testing.T) {
//	const testDataBaseSize = 20
//	db.Storage.InitStorage(testDataBaseSize)
//	storage := db.Storage.AlbumStorage.Albums
//
//	albumsViews := make([]views.Album, len(storage))
//	for idx, album := range storage {
//		albumsViews[idx] = *GetAlbumView(album.Id)
//	}
//
//	router := InitRouter()
//	apiServer := httptest.NewServer(router)
//	defer apiServer.Close()
//	res, err := http.Get(apiServer.URL + GetAllAlbumsUrl)
//
//	if err != nil {
//		t.Errorf("err: %#v", err)
//	}
//
//	testCase := allAlbumTestCase{
//		albums: albumsViews,
//		status: http.StatusOK,
//	}
//
//	if res.StatusCode != testCase.status {
//		t.Errorf("wrong StatusCode: got %d, expected %d", w.Code, testCase.status)
//	}
//
//	resp := w.Result()
//	body, _ := ioutil.ReadAll(resp.Body)
//
//	var result status.Success
//
//	json.Unmarshal(body, &result)
//
//	data := result.Result.([]interface{})
//	albums := views.GetAlbumsViewsFromInterfaces(data)
//	for idx, view := range testCase.albums {
//		if albums[idx] != view {
//			t.Errorf("wrong Response: got %+v, expected %+v", albums, view)
//		}
//	}
//
//}

func TestGetAlbum(t *testing.T) {
	const testDataBaseSize = 20
	db.Storage.InitStorage(testDataBaseSize)
	storage := db.Storage.AlbumStorage.Albums

	cases := make([]getAlbumTestCase, len(storage))
	for idx, album := range storage {
		cases[idx] = getAlbumTestCase{
			id:     album.Id,
			album:  *GetAlbumView(album.Id),
			status: http.StatusOK,
		}
	}

	for caseNum, item := range cases {
		url := Proto + Host + GetAlbumUrlWithoutId + fmt.Sprint(item.id)
		req := httptest.NewRequest("GET", url, nil)

		w := httptest.NewRecorder()

		vars := map[string]string{
			"id": fmt.Sprint(item.id),
		}
		req = mux.SetURLVars(req, vars)
		GetAlbum(w, req)

		if w.Code != item.status {
			t.Errorf("[%d] wrong StatusCode: got %d, expected %d",
				caseNum, w.Code, item.status)
		}

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		var response status.Success
		json.Unmarshal(body, &response)
		result := views.FromInterfaceToAlbumView(response.Result)
		if result != item.album {
			t.Errorf("[%d] wrong Response: got %+v, expected %+v",
				caseNum, result, item.album)
		}
	}
}

// ----------------------------------------------------------------------
type allAlbumTestCase struct {
	albums []views.Album
	status int
}

func TestGetAlbums(t *testing.T) {
	const testDataBaseSize = 20
	db.Storage.InitStorage(testDataBaseSize)
	storage := db.Storage.AlbumStorage.Albums

	albumsViews := make([]views.Album, len(storage))
	for idx, album := range storage {
		albumsViews[idx] = *GetAlbumView(album.Id)
	}

	url := Proto + Host + GetAllAlbumsUrl
	req := httptest.NewRequest("GET", url, nil)

	w := httptest.NewRecorder()

	GetAlbums(w, req)

	testCase := allAlbumTestCase{
		albums: albumsViews,
		status: http.StatusOK,
	}

	if w.Code != testCase.status {
		t.Errorf("wrong StatusCode: got %d, expected %d", w.Code, testCase.status)
	}

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	var response status.Success

	json.Unmarshal(body, &response)
	data := response.Result.([]interface{})
	albums := views.GetAlbumsViewsFromInterfaces(data)
	for idx, view := range testCase.albums {
		if albums[idx] != view {
			t.Errorf("wrong Response: got %+v, expected %+v", albums, view)
		}
	}
}

// ----------------------------------------------------------------------
type createAlbumTestCase struct {
	status int
}

func TestCreateAlbum(t *testing.T) {
	const testDataBaseSize = 20
	db.Storage.InitStorage(testDataBaseSize)

	album := models.Album{
		Id:             0,
		Title:          "imagine",
		AuthorId:       5,
		CountLikes:     50,
		CountListening: 500,
		Date:           0,
		CoverId:        5000,
	}

	url := Proto + Host + GetAllAlbumsUrl
	dataToSend, _ := json.Marshal(album)
	req := httptest.NewRequest("POST", url, bytes.NewBuffer(dataToSend))

	w := httptest.NewRecorder()
	CreateAlbum(w, req)

	testCase := createAlbumTestCase{
		status: http.StatusOK,
	}

	if w.Code != testCase.status {
		t.Errorf("wrong StatusCode: got %d, expected %d", w.Code, testCase.status)
	}

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	var result status.Success
	json.Unmarshal(body, &result)

	expected := db.SuccessWrapper(album.Title, db.SuccessCreatedAlbum)

	if result != status.MakeSuccess(expected) {
		t.Errorf("wrong Response: got %+v, expected %+v",
			result, expected)
	}
}

// ----------------------------------------------------------------------
type deleteAlbumTestCase struct {
	status int
}

func TestDeleteAlbum(t *testing.T) {
	const testDataBaseSize = 20
	db.Storage.InitStorage(testDataBaseSize)

	sizeBefore := len(db.Storage.AlbumStorage.Albums)
	sizeAfter := sizeBefore - 1

	albumToDelete := db.Storage.AlbumStorage.Albums[0]

	url := Proto + Host + DeleteAlbumUrl
	req := httptest.NewRequest("DELETE", url, nil)

	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": fmt.Sprint(albumToDelete.Id),
	}
	req = mux.SetURLVars(req, vars)
	DeleteAlbum(w, req)

	testCase := deleteAlbumTestCase{
		status: http.StatusOK,
	}

	if w.Code != testCase.status {
		t.Errorf("wrong StatusCode: got %d, expected %d", w.Code, testCase.status)
	}

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	var result status.Success
	json.Unmarshal(body, &result)
	expected := db.SuccessWrapper(albumToDelete.Id, db.SuccessDeletedAlbum)
	if result != status.MakeSuccess(expected) {
		t.Errorf("wrong Response: got %+v, expected %+v",
			result, expected)
	}
	require.Equal(t, sizeAfter, len(db.Storage.AlbumStorage.Albums))
}

// ----------------------------------------------------------------------
type updateAlbumTestCase struct {
	status int
}

func TestUpdateAlbum(t *testing.T) {
	const testDataBaseSize = 20
	db.Storage.InitStorage(testDataBaseSize)

	album := models.Album{
		Id:             0,
		Title:          "imagine",
		AuthorId:       5,
		CountLikes:     50,
		CountListening: 500,
		Date:           0,
		CoverId:        5000,
	}

	url := Proto + Host + UpdateAlbumUrl
	dataToSend, _ := json.Marshal(album)
	req := httptest.NewRequest("PUT", url, bytes.NewBuffer(dataToSend))

	w := httptest.NewRecorder()
	UpdateAlbum(w, req)

	testCase := updateAlbumTestCase{
		status: http.StatusOK,
	}

	if w.Code != testCase.status {
		t.Errorf("wrong StatusCode: got %d, expected %d", w.Code, testCase.status)
	}

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	var result status.Success
	json.Unmarshal(body, &result)
	expected := db.SuccessWrapper(album.Id, db.SuccessUpdatedAlbum)
	if result != status.MakeSuccess(expected) {
		t.Errorf("wrong Response: got %+v, expected %+v",
			result, expected)
	}
	require.Equal(t, album.Title, db.Storage.AlbumStorage.Albums[0].Title)
}

// ----------------------------------------------------------------------
type popularAlbumTestCase struct {
	albums []views.Album
	status int
}

func TestPopularAlbums(t *testing.T) {
	const testDataBaseSize = 20
	db.Storage.InitStorage(testDataBaseSize)
	storage := db.Storage.AlbumStorage.Albums

	sort.SliceStable(storage, func(i int, j int) bool {
		return storage[i].CountListening > storage[j].CountListening
	})

	albumsViews := make([]views.Album, len(storage))
	for idx, _ := range storage {
		albumsViews[idx] = *GetAlbumView(uint64(idx))
	}

	url := Proto + Host + GetPopularAlbumsUrl
	req := httptest.NewRequest("GET", url, nil)

	w := httptest.NewRecorder()

	GetPopularAlbums(w, req)

	testCase := popularAlbumTestCase{
		albums: albumsViews,
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
	albums := views.GetAlbumsViewsFromInterfaces(data)
	for idx, album := range testCase.albums {
		if albums[idx] != album {
			t.Errorf("wrong Response: got %+v, expected %+v", albums, album)
		}
	}
}
