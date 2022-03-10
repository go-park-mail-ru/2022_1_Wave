package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2022_1_Wave/db"
	"github.com/go-park-mail-ru/2022_1_Wave/db/models"
	"github.com/go-park-mail-ru/2022_1_Wave/db/views"
	"github.com/go-park-mail-ru/2022_1_Wave/pkg/status"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func getAllAlbums(albumRep db.AlbumRep) (*[]models.Album, error) {
	return albumRep.GetAllAlbums()
}

func addAlbumToStorage(albumRep db.AlbumRep, album models.Album) error {
	return albumRep.Insert(&album)
}

func updateAlbumInStorage(albumRep db.AlbumRep, album models.Album) error {
	return albumRep.Update(&album)
}

func deleteAlbumFromStorageByID(albumRep db.AlbumRep, id uint64) error {
	return albumRep.Delete(id)
}

func getAlbumByIDFromStorage(albumRep db.AlbumRep, id uint64) (*models.Album, error) {
	return albumRep.SelectByID(id)
}

func getPopularAlbums(albumRep db.AlbumRep) (*[]models.Album, error) {
	return albumRep.GetPopularAlbums()
}

func GetAlbumView(id uint64) *views.Album {
	album, err := getAlbumByIDFromStorage(&db.Storage.AlbumStorage, id)

	if err != nil {
		//status.WriteError(w, err, http.StatusBadRequest)
		return nil
	}

	artist, _ := getArtistByIDFromStorage(&db.Storage.ArtistStorage, album.AuthorId)

	currentAlbumView := views.Album{
		Title:  album.Title,
		Artist: artist.Name,
		Cover:  "assets/" + "album_" + fmt.Sprint(album.CoverId) + ".png",
	}

	return &currentAlbumView
}

// GetAlbums godoc
// @Summary      GetAlbums
// @Description  getting all albums
// @Tags     album
// @Accept	 application/json
// @Produce  application/json
// @Success  200 {object} status.Success
// @Failure 400 {object} status.Error "Data is invalid"
// @Failure 405 {object} status.Error "Method is not allowed"
// @Router   /api/v1/albums/ [get]
func GetAlbums(w http.ResponseWriter, r *http.Request) {
	storage := &db.Storage.AlbumStorage
	albums, err := getAllAlbums(storage)
	if err != nil {
		status.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if albums == nil {
		*albums = []models.Album{}
	}

	albumViews := make([]views.Album, len(*albums))

	for i, album := range *albums {
		view := GetAlbumView(album.Id)
		if view == nil {
			status.WriteError(w, errors.New(db.AlbumIsNotExist), http.StatusBadRequest)
		}
		albumViews[i] = *view
	}

	status.WriteSuccess(w, albumViews)
}

// CreateAlbum godoc
// @Summary      CreateAlbum
// @Description  creating new album
// @Tags     album
// @Accept	 application/json
// @Produce  application/json
// @Param    Album body models.Album true  "params of new album. Id will be set automatically."
// @Success  200 {object} status.Success
// @Failure 400 {object} status.Error "Data is invalid"
// @Failure 405 {object} status.Error "Method is not allowed"
// @Router   /api/v1/albums/ [post]
func CreateAlbum(w http.ResponseWriter, r *http.Request) {
	newAlbum := &models.Album{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		status.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(body, newAlbum); err != nil {
		status.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = db.CheckAlbum(newAlbum); err != nil {
		status.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = addAlbumToStorage(&db.Storage.AlbumStorage, *newAlbum); err != nil {
		status.WriteError(w, err, http.StatusBadRequest)
		return
	}

	status.WriteSuccess(w, db.SuccessWrapper(newAlbum.Title, db.SuccessCreatedAlbum))
}

// UpdateAlbum godoc
// @Summary      UpdateAlbum
// @Description  updating album by id
// @Tags     album
// @Accept	 application/json
// @Produce  application/json
// @Param    Album body models.Album true  "id of updating album and params of it."
// @Success  200 {object} status.Success
// @Failure 400 {object} status.Error "Data is invalid"
// @Failure 405 {object} status.Error "Method is not allowed"
// @Router   /api/v1/albums/ [put]
func UpdateAlbum(w http.ResponseWriter, r *http.Request) {
	newAlbum := &models.Album{}
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		status.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(body, newAlbum); err != nil {
		status.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = db.CheckAlbum(newAlbum); err != nil {
		status.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = updateAlbumInStorage(&db.Storage.AlbumStorage, *newAlbum); err != nil {
		status.WriteError(w, err, http.StatusBadRequest)
		return
	}
	status.WriteSuccess(w, db.SuccessWrapper(newAlbum.Id, db.SuccessUpdatedAlbum))
}

// GetAlbum godoc
// @Summary      GetAlbum
// @Description  getting album by id
// @Tags     album
// @Accept	 application/json
// @Produce  application/json
// @Param    id path integer true  "id of album which need to be getted"
// @Success  200 {object} models.Album
// @Failure 400 {object} status.Error "Data is invalid"
// @Failure 405 {object} status.Error "Method is not allowed"
// @Router   /api/v1/albums/{id} [get]
func GetAlbum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars[FieldId])
	if err != nil {
		status.WriteError(w, err, http.StatusBadRequest)
		return
	}

	//id--
	if id < 0 {
		status.WriteError(w, errors.New(db.IndexOutOfRange), http.StatusBadRequest)
		return
	}

	currentAlbumView := GetAlbumView(uint64(id))

	if currentAlbumView == nil {
		status.WriteError(w, errors.New(db.AlbumIsNotExist), http.StatusBadRequest)
		return
	}

	status.WriteSuccess(w, currentAlbumView)
}

// DeleteAlbum godoc
// @Summary      DeleteAlbum
// @Description  deleting album by id
// @Tags     album
// @Accept	 application/json
// @Produce  application/json
// @Param    id path integer true  "id of album which need to be deleted"
// @Success  200 {object} status.Success
// @Failure 400 {object} status.Error "Data is invalid"
// @Failure 405 {object} status.Error "Method is not allowed"
// @Router   /api/v1/albums/{id} [delete]
func DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars[FieldId])
	if err != nil {
		status.WriteError(w, err, http.StatusBadRequest)
		return
	}
	//id--
	if id < 0 {
		status.WriteError(w, errors.New(db.IndexOutOfRange), http.StatusBadRequest)
		return
	}
	err = deleteAlbumFromStorageByID(&db.Storage.AlbumStorage, uint64(id))
	if err != nil {
		status.WriteError(w, err, http.StatusBadRequest)
		return
	}
	status.WriteSuccess(w, db.SuccessWrapper(id, db.SuccessDeletedAlbum))
}

// GetPopularAlbums godoc
// @Summary      GetPopularAlbums
// @Description  getting top20 popular albums
// @Tags     album
// @Accept	 application/json
// @Produce  application/json
// @Success  200 {object} status.Success
// @Failure 400 {object} status.Error "Data is invalid"
// @Failure 405 {object} status.Error "Method is not allowed"
// @Router   /api/v1/albums/popular [get]
func GetPopularAlbums(w http.ResponseWriter, r *http.Request) {
	storage := &db.Storage.AlbumStorage
	storage.Mutex.RLock()
	defer storage.Mutex.RUnlock()
	albums, err := getPopularAlbums(storage)
	if err != nil {
		status.WriteError(w, err, http.StatusBadRequest)
		return
	}

	albumViews := make([]views.Album, len(*albums))

	for i, _ := range *albums {
		view := GetAlbumView(uint64(i))
		if view == nil {
			status.WriteError(w, errors.New(db.AlbumIsNotExist), http.StatusBadRequest)
		}
		albumViews[i] = *view
	}
	status.WriteSuccess(w, albumViews)
}
