package webUtils

import (
	"encoding/json"
	"errors"
	"net/http"
)

const (
	OK   = "OK"
	FAIL = "FAIL"
)

type Success struct {
	Status string `json:"status"`
	Result interface{}
}

type Error struct {
	Status string `json:"status"`
	Err    string
}

func (err Error) makeError(msg string) Error {
	return Error{
		Status: FAIL,
		Err:    msg,
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
