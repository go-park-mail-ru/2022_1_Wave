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

//func makeAlbum(id int, title string, authorId int, countLikes int, countListening int, date int, coverId int) (Album, error) {
//	if id < 0 || authorId < 0 {
//		return Album{}, errors.New("invalid id")
//	}
//
//	if len(title) > db.AlbumTitleLen {
//		title = title[:db.AlbumTitleLen]
//	}
//
//	return Album{
//		Id:             id,
//		Title:          title,
//		AuthorId:       authorId,
//		CountLikes:     countLikes,
//		CountListening: countListening,
//		Date:           date,
//		CoverId:        coverId,
//	}, nil
//}

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

func CreateAlbum(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		UpdateAlbum(w, r)
		return
	}
	if !utils.MethodsIsEqual(w, r.Method, http.MethodPost) {
		return
	}

	newAlbum := &models.Album{}
	newAlbum.Id = uint64(len(db.Storage.AlbumStorage.Albums))
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(body, newAlbum); err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = newAlbum.CheckAlbum(); err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = addAlbumToStorage(&db.Storage.AlbumStorage, *newAlbum); err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(utils.Success{
		Result: db.SuccessCreatedAlbum + "(" + newAlbum.Title + ")",
	})
	fmt.Println("albums storage now:", db.Storage.AlbumStorage.Albums)
}

func UpdateAlbum(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		CreateAlbum(w, r)
		return
	}
	if !utils.MethodsIsEqual(w, r.Method, http.MethodPut) {
		return
	}

	newAlbum := &models.Album{}
	newAlbum.Id = uint64(len(db.Storage.AlbumStorage.Albums))
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(body, newAlbum); err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = newAlbum.CheckAlbum(); err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err = updateAlbumInStorage(&db.Storage.AlbumStorage, *newAlbum); err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(utils.Success{
		Result: db.SuccessUpdatedAlbum + "(" + fmt.Sprint(newAlbum.Id) + ")",
	})
	fmt.Println("albums storage now:", db.Storage.AlbumStorage.Albums)

}

func GetAlbum(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		DeleteAlbum(w, r)
		return
	}
	if r.Method == http.MethodPut {
		UpdateAlbum(w, r)
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
	currentAlbum, err := getAlbumByIDFromStorage(&db.Storage.AlbumStorage, uint64(id))
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(currentAlbum)
	fmt.Println("albums storage now:", db.Storage.AlbumStorage.Albums)
}

func DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		GetAlbum(w, r)
		return
	}
	if r.Method == http.MethodPut {
		UpdateAlbum(w, r)
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
	err = deleteAlbumFromStorageByID(&db.Storage.AlbumStorage, uint64(id))
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(utils.Success{
		Result: db.SuccessDeletedAlbum + "(" + fmt.Sprint(id) + ")",
	})

	fmt.Println("albums storage now:", db.Storage.AlbumStorage.Albums)
}
