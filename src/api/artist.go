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

func getAllArtists(artistRep db.ArtistRep) (*[]models.Artist, error) {
	return artistRep.GetAllArtists()
}

func addArtistToStorage(artistRep db.ArtistRep, artist models.Artist) error {
	return artistRep.Insert(&artist)
}

func updateArtistInStorage(artistRep db.ArtistRep, artist models.Artist) error {
	return artistRep.Update(&artist)
}

func deleteArtistFromStorageByID(artistRep db.ArtistRep, id uint64) error {
	return artistRep.Delete(id)
}

func getArtistByIDFromStorage(artistRep db.ArtistRep, id uint64) (*models.Artist, error) {
	return artistRep.SelectByID(id)
}

func getPopularArtists(artistRep db.ArtistRep) (*[]models.Artist, error) {
	return artistRep.GetPopularArtists()
}

func GetArtistView(id uint64) *views.Artist {
	currentArtist, err := getArtistByIDFromStorage(&db.Storage.ArtistStorage, id)

	if err != nil {
		//status.WriteError(w, err, http.StatusBadRequest)
		return nil
	}

	currentArtistView := views.Artist{
		Name:  currentArtist.Name,
		Cover: "assets/" + "artist_" + fmt.Sprint(currentArtist.Id) + ".png",
	}

	return &currentArtistView
}

// GetArtists godoc
// @Summary      GetArtists
// @Description  getting all artists
// @Tags     artist
// @Accept	 application/json
// @Produce  application/json
// @Success  200 {object} status.Success
// @Failure 400 {object} status.Error "Data is invalid"
// @Failure 405 {object} status.Error "Method is not allowed"
// @Router   /api/v1/artists/ [get]
func GetArtists(w http.ResponseWriter, _ *http.Request) {
	artistStorage := &db.Storage.ArtistStorage
	artistStorage.Mutex.RLock()
	defer artistStorage.Mutex.RUnlock()
	artists, err := getAllArtists(artistStorage)
	if err != nil {
		status.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if artists == nil {
		*artists = []models.Artist{}
	}

	artistViews := make([]views.Artist, len(*artists))

	for i, artist := range *artists {
		view := GetArtistView(artist.Id)
		if view == nil {
			status.WriteError(w, errors.New(db.ArtistIsNotExist), http.StatusBadRequest)
		}
		artistViews[i] = *view
	}
	status.WriteSuccess(w, artistViews)
}

// CreateArtist godoc
// @Summary      CreateArtist
// @Description  creating new artist
// @Tags     artist
// @Accept	 application/json
// @Produce  application/json
// @Param    Artist body models.Artist true  "params of new artist. Id will be set automatically."
// @Success  200 {object} status.Success
// @Failure 400 {object} status.Error "Data is invalid"
// @Failure 405 {object} status.Error "Method is not allowed"
// @Router   /api/v1/artists/ [post]
func CreateArtist(w http.ResponseWriter, r *http.Request) {
	newArtist := &models.Artist{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		status.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(body, newArtist); err != nil {
		status.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = db.CheckArtist(newArtist); err != nil {
		status.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = addArtistToStorage(&db.Storage.ArtistStorage, *newArtist); err != nil {
		status.WriteError(w, err, http.StatusBadRequest)
		return
	}

	status.WriteSuccess(w, db.SuccessWrapper(newArtist.Name, db.SuccessCreatedArtist))
}

// UpdateArtist godoc
// @Summary      UpdateArtist
// @Description  updating artist by id
// @Tags     artist
// @Accept	 application/json
// @Produce  application/json
// @Param    Artist body models.Artist true  "id of updating artist and params of it."
// @Success  200 {object} status.Success
// @Failure 400 {object} status.Error "Data is invalid"
// @Failure 405 {object} status.Error "Method is not allowed"
// @Router   /api/v1/artists/ [put]
func UpdateArtist(w http.ResponseWriter, r *http.Request) {
	newArtist := &models.Artist{}
	newArtist.Id = uint64(len(db.Storage.ArtistStorage.Artists))
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		status.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(body, newArtist); err != nil {
		status.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = db.CheckArtist(newArtist); err != nil {
		status.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = updateArtistInStorage(&db.Storage.ArtistStorage, *newArtist); err != nil {
		status.WriteError(w, err, http.StatusBadRequest)
		return
	}

	status.WriteSuccess(w, db.SuccessWrapper(newArtist.Id, db.SuccessUpdatedArtist))
}

// GetArtist godoc
// @Summary      GetArtist
// @Description  getting artist by id
// @Tags     artist
// @Accept	 application/json
// @Produce  application/json
// @Param    id path integer true  "id of artist which need to be getted"
// @Success  200 {object} models.Artist
// @Failure 400 {object} status.Error "Data is invalid"
// @Failure 405 {object} status.Error "Method is not allowed"
// @Router   /api/v1/artists/{id} [get]
func GetArtist(w http.ResponseWriter, r *http.Request) {
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
	currentArtistView := GetArtistView(uint64(id))

	if currentArtistView == nil {
		status.WriteError(w, errors.New(db.ArtistIsNotExist), http.StatusBadRequest)
		return
	}

	status.WriteSuccess(w, currentArtistView)
}

// DeleteArtist godoc
// @Summary      DeleteArtist
// @Description  deleting artist by id
// @Tags     artist
// @Accept	 application/json
// @Produce  application/json
// @Param    id path integer true  "id of artist which need to be deleted"
// @Success  200 {object} status.Success
// @Failure 400 {object} status.Error "Data is invalid"
// @Failure 405 {object} status.Error "Method is not allowed"
// @Router   /api/v1/artists/{id} [delete]
func DeleteArtist(w http.ResponseWriter, r *http.Request) {
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
	err = deleteArtistFromStorageByID(&db.Storage.ArtistStorage, uint64(id))
	if err != nil {
		status.WriteError(w, err, http.StatusBadRequest)
		return
	}
	status.WriteSuccess(w, db.SuccessWrapper(id, db.SuccessDeletedArtist))
}

// GetPopularArtists godoc
// @Summary      GetPopularArtists
// @Description  getting top20 popular artists
// @Tags     artist
// @Accept	 application/json
// @Produce  application/json
// @Success  200 {object} status.Success
// @Failure 400 {object} status.Error "Data is invalid"
// @Failure 405 {object} status.Error "Method is not allowed"
// @Router   /api/v1/artists/popular [get]
func GetPopularArtists(w http.ResponseWriter, r *http.Request) {
	storage := &db.Storage.ArtistStorage
	storage.Mutex.RLock()
	defer storage.Mutex.RUnlock()
	artists, err := getPopularArtists(storage)
	if err != nil {
		status.WriteError(w, err, http.StatusBadRequest)
		return
	}

	artistViews := make([]views.Artist, len(*artists))

	for i, artist := range *artists {
		view := GetArtistView(artist.Id)
		if view == nil {
			status.WriteError(w, errors.New(db.ArtistIsNotExist), http.StatusBadRequest)
		}
		artistViews[i] = *view
	}
	status.WriteSuccess(w, artistViews)
}
