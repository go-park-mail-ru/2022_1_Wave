package utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

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

func WriteError(w http.ResponseWriter, err error, status int) {
	response, _ := json.Marshal(Error{}.makeError(err.Error()))
	http.Error(w, string(response), status)
	return
}

func MethodsIsEqual(w http.ResponseWriter, actualMethod string, expectedMethod string) bool {
	if actualMethod != expectedMethod {
		WriteError(w, errors.New("expected method "+expectedMethod), http.StatusMethodNotAllowed)
		return false
	}
	return true
}
