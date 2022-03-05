package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2022_1_Wave/db"
	"io"
	"net/http"
)

type Album struct {
	id             int
	title          string
	authorId       int
	countLikes     int
	countListening int
	date           int
	coverId        int
}

func MakeAlbum(id int, title string, authorId int, countLikes int, countListening int, date int, coverId int) (Album, error) {
	if id < 0 || authorId < 0 {
		return Album{}, errors.New("invalid id")
	}

	if len(title) > db.AlbumTitleLen {
		title = title[:db.AlbumTitleLen]
	}

	return Album{
		id:             id,
		title:          title,
		authorId:       authorId,
		countLikes:     countLikes,
		countListening: countListening,
		date:           date,
		coverId:        coverId,
	}, nil
}

type Success struct {
	Result string `json:"result"`
}

type Error struct {
	Err string `json:"error"`
}

func (err Error) makeError(msg string) Error {
	return Error{
		Err: msg,
	}
}

func writeError(w http.ResponseWriter, msg string, status int) {
	response, _ := json.Marshal(Error{}.makeError(msg))
	http.Error(w, string(response), status)
	return
}

// expected
const (
	expectedPost = "Expected method POST."
)

// errors
const (
	invalidBody = "Invalid body format."
	invalidJson = "Error to unpacking json."
)

// success
const (
	successCreatedAlbum = "Success created album."
)

func (api *Album) CreateAlbum(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method != http.MethodPost {
		writeError(w, expectedPost, http.StatusMethodNotAllowed)
		return
	}
	album := &Album{}
	body, err := io.ReadAll(r.Body)

	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, album)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("created album=", album.coverId, album.title, album.authorId)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Success{
		Result: successCreatedAlbum,
	})
}
