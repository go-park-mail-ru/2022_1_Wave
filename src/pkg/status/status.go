package status

import (
	"encoding/json"
	"net/http"
)

const (
	OK   = "OK"
	FAIL = "FAIL"
)

type Success struct {
	Status string `json:"status" example:"OK"`
	Result interface{}
}

type Error struct {
	Status string `json:"status" example:"FAIL"`
	Error  string
}

func (err Error) MakeError(msg string) Error {
	return Error{
		Status: FAIL,
		Error:  msg,
	}
}

func MakeSuccess(obj interface{}) Success {
	return Success{
		Status: OK,
		Result: obj,
	}
}

func WriteSuccess(w http.ResponseWriter, result interface{}) {
	json.NewEncoder(w).Encode(MakeSuccess(result))
}

func WriteError(w http.ResponseWriter, err error, status int) {
	response, _ := json.Marshal(Error{}.MakeError(err.Error()))
	http.Error(w, string(response), status)
	return
}

//func MethodsIsEqual(w http.ResponseWriter, actualMethod string, expectedMethod string) bool {
//	if actualMethod != expectedMethod {
//		WriteError(w, errors.New("expected method "+expectedMethod), http.StatusMethodNotAllowed)
//		return false
//	}
//	return true
//}
