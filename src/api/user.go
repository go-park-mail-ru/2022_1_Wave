package api

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2022_1_Wave/db"
	"github.com/go-park-mail-ru/2022_1_Wave/service"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type UserGetResponse struct {
	Status string  `json:"status"`
	Body   db.User `json:"body"`
}

func GetSelfUser(w http.ResponseWriter, r *http.Request) {
	user, err := service.GetSession(r)
	if err != nil {
		http.Error(w, `{"error": "no auth"}`, http.StatusForbidden)
		return
	}

	userFromDb, err := db.MyUserStorage.SelectByID(user.UserId)
	if err != nil {
		http.Error(w, `{"error": "no user with such id"}`, http.StatusNotFound)
		return
	}

	userFromDbCopy := *userFromDb
	userFromDbCopy.Password = ""

	json.NewEncoder(w).Encode(userFromDbCopy)
}

// 127.0.0.1/api/v1/users/<id>
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, `{"error": "bad id"}`, http.StatusBadRequest)
		return
	}

	userFromDb, err := db.MyUserStorage.SelectByID(uint(userId))

	if err != nil {
		http.Error(w, `{"error": "user not found"}`, http.StatusNotFound)
		return
	}

	userFromDbCopy := *userFromDb
	userFromDbCopy.Password = ""

	json.NewEncoder(w).Encode(userFromDbCopy)
}
