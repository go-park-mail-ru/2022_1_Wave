package api

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2022_1_Wave/db/models"
	"github.com/go-park-mail-ru/2022_1_Wave/service"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type UserGetResponse struct {
	Status string      `json:"status"`
	Result models.User `json:"result"`
}

// GetSelfUser godoc
// @Summary      GetSelfUser
// @Description  get user
// @Tags     user
// @Accept	 application/json
// @Produce  application/json
// @Param    id path integer true  "id of user which need to be getted"
// @Success  200 {object} forms.User
// @Failure 401 {object} forms.Result "unauthorized"
// @Failure 401 {object} forms.Result "invalid csrf"
// @Router   /net/v1/users/self [get]
func GetSelfUser(w http.ResponseWriter, r *http.Request) {
	user, _ := service.GetSession(r)

	userFromDb, _ := models.MyUserStorage.SelectByID(user.UserId)
	/*if err != nil {
		http.Error(w, `{"status": "FAIL", "error": "no user with such id"}`, http.StatusNotFound)
		return
	}*/

	userFromDbCopy := *userFromDb
	userFromDbCopy.Password = ""

	response := &UserGetResponse{
		Status: "OK",
		Result: userFromDbCopy,
	}

	json.NewEncoder(w).Encode(response)
}

// GetUser godoc
// @Summary      GetUser
// @Description  get user by cookie
// @Tags     user
// @Accept	 application/json
// @Produce  application/json
// @Success  200 {object} forms.User
// @Failure 400 {object} forms.Result "invalid id"
// @Failure 404 {object} forms.Result "user not found"
// @Router   /net/v1/users/{id} [get]
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, `{"status": "FAIL", "error": "invalid id"}`, http.StatusBadRequest)
		return
	}

	userFromDb, err := models.MyUserStorage.SelectByID(uint(userId))

	if err != nil {
		http.Error(w, `{"status": "FAIL", "error": "user not found"}`, http.StatusNotFound)
		return
	}

	userFromDbCopy := *userFromDb
	userFromDbCopy.Password = ""

	response := &UserGetResponse{
		Status: "OK",
		Result: userFromDbCopy,
	}

	json.NewEncoder(w).Encode(response)
}
