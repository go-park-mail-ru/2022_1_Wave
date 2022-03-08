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
//		Cover:        coverId,
//	}, nil
//}

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

// GetAlbums godoc
// @Summary      GetAlbums
// @Description  getting all albums
// @Tags     album
// @Accept	 application/json
// @Produce  application/json
// @Success  200 {object} utils.Success
// @Failure 400 {object} utils.Error "Data is invalid"
// @Failure 405 {object} utils.Error "Method is not allowed"
// @Router   /api/v1/albums/ [get]
func GetAlbums(w http.ResponseWriter, r *http.Request) {
	storage := &db.Storage.AlbumStorage
	storage.Mutex.RLock()
	defer storage.Mutex.RUnlock()
	albums, err := getAllAlbums(storage)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if *albums == nil {
		*albums = []models.Album{}
	}

	result, _ := json.MarshalIndent(albums, "", "    ")
	fmt.Println(utils.Success{
		Result: string(result)})
	json.NewEncoder(w).Encode(utils.Success{
		Result: string(result)})

}

// CreateAlbum godoc
// @Summary      CreateAlbum
// @Description  creating new album
// @Tags     album
// @Accept	 application/json
// @Produce  application/json
// @Param    Album body models.Album true  "params of new album. Id will be set automatically."
// @Success  200 {object} utils.Success
// @Failure 400 {object} utils.Error "Data is invalid"
// @Failure 405 {object} utils.Error "Method is not allowed"
// @Router   /api/v1/albums/ [post]
func CreateAlbum(w http.ResponseWriter, r *http.Request) {
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
}

// UpdateAlbum godoc
// @Summary      UpdateAlbum
// @Description  updating album by id
// @Tags     album
// @Accept	 application/json
// @Produce  application/json
// @Param    Album body models.Album true  "id of updating album and params of it."
// @Success  200 {object} utils.Success
// @Failure 400 {object} utils.Error "Data is invalid"
// @Failure 405 {object} utils.Error "Method is not allowed"
// @Router   /api/v1/albums/ [put]
func UpdateAlbum(w http.ResponseWriter, r *http.Request) {
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
}

// GetAlbum godoc
// @Summary      GetAlbum
// @Description  getting album by id
// @Tags     album
// @Accept	 application/json
// @Produce  application/json
// @Param    id path integer true  "id of album which need to be getted"
// @Success  200 {object} models.Album
// @Failure 400 {object} utils.Error "Data is invalid"
// @Failure 405 {object} utils.Error "Method is not allowed"
// @Router   /api/v1/albums/{id} [get]
func GetAlbum(w http.ResponseWriter, r *http.Request) {
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
}

// DeleteAlbum godoc
// @Summary      DeleteAlbum
// @Description  deleting album by id
// @Tags     album
// @Accept	 application/json
// @Produce  application/json
// @Param    id path integer true  "id of album which need to be deleted"
// @Success  200 {object} utils.Success
// @Failure 400 {object} utils.Error "Data is invalid"
// @Failure 405 {object} utils.Error "Method is not allowed"
// @Router   /api/v1/albums/{id} [delete]
func DeleteAlbum(w http.ResponseWriter, r *http.Request) {
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
}

// GetPopularAlbums godoc
// @Summary      GetPopularAlbums
// @Description  getting top20 popular albums
// @Tags     album
// @Accept	 application/json
// @Produce  application/json
// @Success  200 {object} utils.Success
// @Failure 400 {object} utils.Error "Data is invalid"
// @Failure 405 {object} utils.Error "Method is not allowed"
// @Router   /api/v1/albums/popular [get]
func GetPopularAlbums(w http.ResponseWriter, r *http.Request) {
	storage := &db.Storage.AlbumStorage
	storage.Mutex.RLock()
	defer storage.Mutex.RUnlock()
	albums, err := getPopularAlbums(storage)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}
	result, _ := json.MarshalIndent(albums, "", "    ")
	json.NewEncoder(w).Encode(utils.Success{
		Result: string(result)})
}
