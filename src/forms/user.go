package forms

import (
	"encoding/json"
	"errors"
	"net/http"
)

type User struct {
	ID       int    `json:"id" example:"5125112"`
	Username string `json:"name" example:"Martin"`
	Email    string `json:"email" example:"hello@example.com"`
	Password string `json:"password" example:"1fsgh2rfafas"`
}

func UserUnmarshal(r *http.Request) (*User, error) {
	decoder := json.NewDecoder(r.Body)
	userToLogin := new(User)
	err := decoder.Decode(userToLogin)
	if err != nil {
		return nil, errors.New("Invalid json")
	}

	return userToLogin, nil
}
