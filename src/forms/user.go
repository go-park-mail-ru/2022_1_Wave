package forms

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"regexp"
)

var emailRegex, _ = regexp.Compile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

type User struct {
	ID       int    `json:"id" example:"5125112"`
	Username string `json:"username" example:"Martin"`
	Email    string `json:"email" example:"hello@example.com"`
	Password string `json:"password" example:"1fsgh2rfafas"`
}

func UserUnmarshal(r *http.Request) (*User, error) {
	decoder := json.NewDecoder(r.Body)
	userToLogin := new(User)
	err := decoder.Decode(userToLogin)
	if err != nil {
		return nil, errors.New("invalid json")
	}

	return userToLogin, nil
}

func (user User) IsValid() bool {
	if user.Username == "" || user.Email == "" || user.Password == "" {
		log.Println("empty")
		return false
	}

	if !emailRegex.MatchString(user.Email) {
		log.Println("invalid email")
		return false
	}

	return true
}
