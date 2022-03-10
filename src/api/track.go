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

func getAllTracks(trackRep db.TrackRep) (*[]models.Track, error) {
	return trackRep.GetAllSongs()
}

func addTrackToStorage(trackRep db.TrackRep, song models.Track) error {
	return trackRep.Insert(&song)
}

func updateTrackInStorage(trackRep db.TrackRep, song models.Track) error {
	return trackRep.Update(&song)
}

func deleteTrackFromStorageByID(trackRep db.TrackRep, id uint64) error {
	return trackRep.Delete(id)
}

func getTrackByIDFromStorage(trackRep db.TrackRep, id uint64) (*models.Track, error) {
	return trackRep.SelectByID(id)
}

func getPopularTacks(trackRep db.TrackRep) (*[]models.Track, error) {
	return trackRep.GetPopularSongs()
}

func GetTrackView(id uint64) *views.Track {
	currentTrack, err := getTrackByIDFromStorage(&db.Storage.TrackStorage, id)

	if err != nil {
		//status.WriteError(w, err, http.StatusBadRequest)
		return nil
	}

	currentTrackArtist, _ := getArtistByIDFromStorage(&db.Storage.ArtistStorage, currentTrack.AuthorId)

	currentTrackView := views.Track{
		Title:  currentTrack.Title,
		Artist: currentTrackArtist.Name,
		Cover:  "assets/" + "track_" + fmt.Sprint(currentTrack.Id) + ".png",
	}

	return &currentTrackView
}

// GetTracks godoc
// @Summary      GetTracks
// @Description  getting all tracks
// @Tags     track
// @Accept	 application/json
// @Produce  application/json
// @Success  200 {object} status.Success
// @Failure 400 {object} status.Error "Data is invalid"
// @Failure 405 {object} status.Error "Method is not allowed"
// @Router   /api/v1/tracks/ [get]
func GetTracks(w http.ResponseWriter, r *http.Request) {
	storage := &db.Storage.TrackStorage
	storage.Mutex.RLock()
	defer storage.Mutex.RUnlock()
	tracks, err := getAllTracks(storage)
	if err != nil {
		status.WriteError(w, err, http.StatusBadRequest)
		return
	}
	if *tracks == nil {
		*tracks = []models.Track{}
	}

	trackViews := make([]views.Track, len(*tracks))

	for i, track := range *tracks {
		view := GetTrackView(track.Id)
		if view == nil {
			status.WriteError(w, errors.New(db.TrackIsNotExist), http.StatusBadRequest)
		}
		trackViews[i] = *view
	}
	status.WriteSuccess(w, trackViews)
}

// CreateTrack godoc
// @Summary      CreateTrack
// @Description  creating new track
// @Tags     track
// @Accept	 application/json
// @Produce  application/json
// @Param    Track body models.Track true  "params of new track. Id will be set automatically."
// @Success  200 {object} status.Success
// @Failure 400 {object} status.Error "Data is invalid"
// @Failure 405 {object} status.Error "Method is not allowed"
// @Router   /api/v1/tracks/ [post]
func CreateTrack(w http.ResponseWriter, r *http.Request) {
	storage := &db.Storage.TrackStorage
	newSong := &models.Track{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		status.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(body, newSong); err != nil {
		status.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = db.CheckSong(newSong); err != nil {
		status.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = addTrackToStorage(storage, *newSong); err != nil {
		status.WriteError(w, err, http.StatusBadRequest)
		return
	}

	status.WriteSuccess(w, db.SuccessWrapper(newSong.Title, db.SuccessCreatedTrack))
}

// UpdateTrack godoc
// @Summary      UpdateTrack
// @Description  updating track by id
// @Tags     track
// @Accept	 application/json
// @Produce  application/json
// @Param    Track body models.Track true  "id of updating song and params of it."
// @Success  200 {object} status.Success
// @Failure 400 {object} status.Error "Data is invalid"
// @Failure 405 {object} status.Error "Method is not allowed"
// @Router   /api/v1/tracks/ [put]
func UpdateTrack(w http.ResponseWriter, r *http.Request) {
	storage := &db.Storage.TrackStorage
	newSong := &models.Track{}
	newSong.Id = uint64(len(storage.Tracks))
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		status.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(body, newSong); err != nil {
		status.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = db.CheckSong(newSong); err != nil {
		status.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = updateTrackInStorage(storage, *newSong); err != nil {
		status.WriteError(w, err, http.StatusBadRequest)
		return
	}

	status.WriteSuccess(w, db.SuccessWrapper(newSong.Id, db.SuccessUpdatedTrack))
}

// GetTrack godoc
// @Summary      GetTrack
// @Description  getting track by id
// @Tags     track
// @Accept	 application/json
// @Produce  application/json
// @Param    id path integer true  "id of track which need to be getted"
// @Success  200 {object} models.Track
// @Failure 400 {object} status.Error "Data is invalid"
// @Failure 405 {object} status.Error "Method is not allowed"
// @Router   /api/v1/tracks/{id} [get]
func GetTrack(w http.ResponseWriter, r *http.Request) {
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
	currentTrackView := GetTrackView(uint64(id))

	if currentTrackView == nil {
		status.WriteError(w, errors.New(db.TrackIsNotExist), http.StatusBadRequest)
		return
	}

	status.WriteSuccess(w, currentTrackView)
}

// DeleteTrack godoc
// @Summary      DeleteTrack
// @Description  deleting track by id
// @Tags     track
// @Accept	 application/json
// @Produce  application/json
// @Param    id path integer true  "id of track which need to be deleted"
// @Success  200 {object} status.Success
// @Failure 400 {object} status.Error "Data is invalid"
// @Failure 405 {object} status.Error "Method is not allowed"
// @Router   /api/v1/tracks/{id} [delete]
func DeleteTrack(w http.ResponseWriter, r *http.Request) {
	storage := &db.Storage.TrackStorage
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
	err = deleteTrackFromStorageByID(storage, uint64(id))
	if err != nil {
		status.WriteError(w, err, http.StatusBadRequest)
		return
	}

	status.WriteSuccess(w, db.SuccessWrapper(id, db.SuccessDeletedTrack))
}

// GetPopularTracks godoc
// @Summary      GetPopularTracks
// @Description  getting top20 popular tracks
// @Tags     track
// @Accept	 application/json
// @Produce  application/json
// @Success  200 {object} status.Success
// @Failure 400 {object} status.Error "Data is invalid"
// @Failure 405 {object} status.Error "Method is not allowed"
// @Router   /api/v1/tracks/popular [get]
func GetPopularTracks(w http.ResponseWriter, r *http.Request) {
	storage := &db.Storage.TrackStorage
	storage.Mutex.RLock()
	defer storage.Mutex.RUnlock()
	tracks, err := getPopularTacks(storage)
	if err != nil {
		status.WriteError(w, err, http.StatusBadRequest)
		return
	}
	trackViews := make([]views.Track, len(*tracks))

	for i, _ := range *tracks {
		view := GetTrackView(uint64(i))
		if view == nil {
			status.WriteError(w, errors.New(db.TrackIsNotExist), http.StatusBadRequest)
		}
		trackViews[i] = *view
	}

	status.WriteSuccess(w, trackViews)
}
