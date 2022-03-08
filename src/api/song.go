package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2022_1_Wave/db"
	"github.com/go-park-mail-ru/2022_1_Wave/db/models"
	"github.com/go-park-mail-ru/2022_1_Wave/db/views"
	"github.com/go-park-mail-ru/2022_1_Wave/pkg/utils"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func getAllSongs(songRep db.SongRep) (*[]models.Track, error) {
	return songRep.GetAllSongs()
}

func addSongToStorage(songRep db.SongRep, song models.Track) error {
	return songRep.Insert(&song)
}

func updateSongInStorage(songRep db.SongRep, song models.Track) error {
	return songRep.Update(&song)
}

func deleteSongFromStorageByID(songRep db.SongRep, id uint64) error {
	return songRep.Delete(id)
}

func getTrackByIDFromStorage(songRep db.SongRep, id uint64) (*models.Track, error) {
	return songRep.SelectByID(id)
}

func getPopularSongs(songRep db.SongRep) (*[]models.Track, error) {
	return songRep.GetPopularSongs()
}

// GetSongs godoc
// @Summary      GetSongs
// @Description  getting all songs
// @Tags     song
// @Accept	 application/json
// @Produce  application/json
// @Success  200 {object} utils.Success
// @Failure 400 {object} utils.Error "Data is invalid"
// @Failure 405 {object} utils.Error "Method is not allowed"
// @Router   /api/v1/songs/ [get]
func GetSongs(w http.ResponseWriter, r *http.Request) {
	storage := &db.Storage.SongStorage
	storage.Mutex.RLock()
	defer storage.Mutex.RUnlock()
	tracks, err := getAllSongs(storage)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}
	if *tracks == nil {
		*tracks = []models.Track{}
	}

	tracksViews := make([]views.Track, len(*tracks))

	for i, track := range *tracks {
		tracksViews[i].Title = track.Title
		artist, _ := getArtistByIDFromStorage(&db.Storage.ArtistStorage, track.AuthorId)
		tracksViews[i].Artist = artist.Name
		tracksViews[i].Cover = "assets/" + "track_" + fmt.Sprint(track.CoverId) + ".png"
		fmt.Println(tracksViews[i])
	}
	json.NewEncoder(w).Encode(utils.Success{
		Result: tracksViews})
}

// CreateSong godoc
// @Summary      CreateSong
// @Description  creating new song
// @Tags     song
// @Accept	 application/json
// @Produce  application/json
// @Param    Track body models.Track true  "params of new song. Id will be set automatically."
// @Success  200 {object} utils.Success
// @Failure 400 {object} utils.Error "Data is invalid"
// @Failure 405 {object} utils.Error "Method is not allowed"
// @Router   /api/v1/songs/ [post]
func CreateSong(w http.ResponseWriter, r *http.Request) {
	storage := &db.Storage.SongStorage
	newSong := &models.Track{}
	newSong.Id = uint64(len(storage.Songs))
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(body, newSong); err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = newSong.CheckSong(); err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = addSongToStorage(storage, *newSong); err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(utils.Success{
		Result: db.SuccessCreatedSong + "(" + newSong.Title + ")",
	})
}

// UpdateSong godoc
// @Summary      UpdateSong
// @Description  updating song by id
// @Tags     song
// @Accept	 application/json
// @Produce  application/json
// @Param    Track body models.Track true  "id of updating song and params of it."
// @Success  200 {object} utils.Success
// @Failure 400 {object} utils.Error "Data is invalid"
// @Failure 405 {object} utils.Error "Method is not allowed"
// @Router   /api/v1/songs/ [put]
func UpdateSong(w http.ResponseWriter, r *http.Request) {
	storage := &db.Storage.SongStorage
	newSong := &models.Track{}
	newSong.Id = uint64(len(storage.Songs))
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(body, newSong); err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = newSong.CheckSong(); err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = updateSongInStorage(storage, *newSong); err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(utils.Success{
		Result: db.SuccessUpdatedSong + "(" + fmt.Sprint(newSong.Id) + ")",
	})
}

// GetSong godoc
// @Summary      GetSong
// @Description  getting song by id
// @Tags     song
// @Accept	 application/json
// @Produce  application/json
// @Param    id path integer true  "id of song which need to be getted"
// @Success  200 {object} models.Track
// @Failure 400 {object} utils.Error "Data is invalid"
// @Failure 405 {object} utils.Error "Method is not allowed"
// @Router   /api/v1/songs/{id} [get]
func GetSong(w http.ResponseWriter, r *http.Request) {
	storage := &db.Storage.SongStorage
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars[FieldId])
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}
	//id--
	if id < 0 {
		utils.WriteError(w, errors.New(db.IndexOutOfRange), http.StatusBadRequest)
		return
	}
	currentTrack, err := getTrackByIDFromStorage(storage, uint64(id))
	currentTrackArtist, _ := getArtistByIDFromStorage(&db.Storage.ArtistStorage, currentTrack.AuthorId)

	currentTrackView := views.Track{
		Title:  currentTrack.Title,
		Artist: currentTrackArtist.Name,
		Cover:  "assets/" + "track_" + fmt.Sprint(currentTrack.CoverId) + ".png",
	}

	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(currentTrackView)
}

// DeleteSong godoc
// @Summary      DeleteSong
// @Description  deleting song by id
// @Tags     song
// @Accept	 application/json
// @Produce  application/json
// @Param    id path integer true  "id of song which need to be deleted"
// @Success  200 {object} utils.Success
// @Failure 400 {object} utils.Error "Data is invalid"
// @Failure 405 {object} utils.Error "Method is not allowed"
// @Router   /api/v1/songs/{id} [delete]
func DeleteSong(w http.ResponseWriter, r *http.Request) {
	storage := &db.Storage.SongStorage
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars[FieldId])
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}
	//id--
	if id < 0 {
		utils.WriteError(w, errors.New(db.IndexOutOfRange), http.StatusBadRequest)
		return
	}
	err = deleteSongFromStorageByID(storage, uint64(id))
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(utils.Success{
		Result: db.SuccessDeletedSong + "(" + fmt.Sprint(id) + ")",
	})
}

// GetPopularSongs godoc
// @Summary      GetPopularSongs
// @Description  getting top20 popular songs
// @Tags     song
// @Accept	 application/json
// @Produce  application/json
// @Success  200 {object} utils.Success
// @Failure 400 {object} utils.Error "Data is invalid"
// @Failure 405 {object} utils.Error "Method is not allowed"
// @Router   /api/v1/songs/popular [get]
func GetPopularSongs(w http.ResponseWriter, r *http.Request) {
	storage := &db.Storage.SongStorage
	storage.Mutex.RLock()
	defer storage.Mutex.RUnlock()
	tracks, err := getPopularSongs(storage)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}
	trackViews := make([]views.Album, len(*tracks))

	for i, track := range *tracks {
		trackViews[i].Title = track.Title
		artist, _ := getArtistByIDFromStorage(&db.Storage.ArtistStorage, track.AuthorId)
		trackViews[i].Artist = artist.Name
		trackViews[i].Cover = "assets/" + "track_" + fmt.Sprint(track.CoverId) + ".png"
	}
	json.NewEncoder(w).Encode(utils.Success{
		Result: trackViews})
}
