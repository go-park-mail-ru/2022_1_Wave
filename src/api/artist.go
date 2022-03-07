package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2022_1_Wave/db"
	"github.com/go-park-mail-ru/2022_1_Wave/db/models"
	"github.com/go-park-mail-ru/2022_1_Wave/pkg/utils"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

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

func CreateArtist(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		UpdateArtist(w, r)
		return
	}
	if !utils.MethodsIsEqual(w, r.Method, http.MethodPost) {
		return
	}

	newArtist := &models.Artist{}
	newArtist.Id = uint64(len(db.Storage.ArtistStorage.Artists))
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(body, newArtist); err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = newArtist.CheckArtist(); err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = addArtistToStorage(&db.Storage.ArtistStorage, *newArtist); err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(utils.Success{
		Result: db.SuccessCreatedArtist + "(" + newArtist.Name + ")",
	})
	fmt.Println("artists storage now:", db.Storage.ArtistStorage.Artists)
}

func UpdateArtist(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		CreateArtist(w, r)
		return
	}
	if !utils.MethodsIsEqual(w, r.Method, http.MethodPut) {
		return
	}

	newArtist := &models.Artist{}
	newArtist.Id = uint64(len(db.Storage.ArtistStorage.Artists))
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(body, newArtist); err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = newArtist.CheckArtist(); err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = updateArtistInStorage(&db.Storage.ArtistStorage, *newArtist); err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(utils.Success{
		Result: db.SuccessUpdatedArtist + "(" + fmt.Sprint(newArtist.Id) + ")",
	})
	fmt.Println("artists storage now:", db.Storage.ArtistStorage.Artists)

}

func GetArtist(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		DeleteArtist(w, r)
		return
	}
	if r.Method == http.MethodPut {
		UpdateArtist(w, r)
		return
	}
	if !utils.MethodsIsEqual(w, r.Method, http.MethodGet) {
		return
	}
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
	currentArtist, err := getArtistByIDFromStorage(&db.Storage.ArtistStorage, uint64(id))
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(currentArtist)
	fmt.Println("artist storage now:", db.Storage.ArtistStorage.Artists)
}

func DeleteArtist(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		GetArtist(w, r)
		return
	}
	if r.Method == http.MethodPut {
		UpdateArtist(w, r)
		return
	}
	if !utils.MethodsIsEqual(w, r.Method, http.MethodDelete) {
		return
	}

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
	err = deleteArtistFromStorageByID(&db.Storage.ArtistStorage, uint64(id))
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(utils.Success{
		Result: db.SuccessDeletedArtist + "(" + fmt.Sprint(id) + ")",
	})

	fmt.Println("artists storage now:", db.Storage.ArtistStorage.Artists)
}
