package api

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2022_1_Wave/config"
	"github.com/go-park-mail-ru/2022_1_Wave/db"
	"github.com/go-park-mail-ru/2022_1_Wave/forms"
	"github.com/go-park-mail-ru/2022_1_Wave/middleware"
	"github.com/go-park-mail-ru/2022_1_Wave/service"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type UserSuccessTestCase struct {
	RequestFunc func(url string) (*UserGetResponse, error)
	Status      int
	Result      *UserGetResponse
}

type UserErrorTestCase struct {
	RequestFunc func(url string) (*forms.Result, error)
	Status      int
	Result      *forms.Result
}

func getUserFromJson(r io.Reader) (*UserGetResponse, error) {
	body, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	userForm := &UserGetResponse{}
	err = json.Unmarshal(body, &userForm)

	if err != nil {
		return nil, err
	}

	return userForm, nil
}

func TestSelfUserSuccessCase(t *testing.T) {
	config.LoadConfig("../../config/config.toml")
	ts := httptest.NewServer(http.HandlerFunc(middleware.CSRF(middleware.Auth(GetSelfUser))))
	serverGetCSRF := httptest.NewServer(http.HandlerFunc(GetCSRF))
	defer ts.Close()
	defer serverGetCSRF.Close()

	cases := []UserSuccessTestCase{
		{
			Status: http.StatusOK,
			Result: &UserGetResponse{
				Status: "OK",
				Result: db.User{
					ID:       1,
					Username: "admin",
					Email:    "admin@samsabaka.ru",
					Password: "",
				},
			},
			RequestFunc: func(url string) (*UserGetResponse, error) {
				csrf, cookie, err := getSession(serverGetCSRF.URL)
				if err != nil {
					return nil, err
				}

				db.MyUserStorage.Insert(&db.User{
					ID:       1,
					Username: "admin",
					Email:    "admin@samsabaka.ru",
					Password: "admin",
				})

				service.AuthorizeUser(cookie.Value, 1)

				resp, err := doGetRequest(url, cookie, map[string]string{"X-CSRF-TOKEN": csrf})
				if err != nil {
					return nil, err
				}

				return getUserFromJson(resp.Body)
			},
		},
	}

	for caseNum, item := range cases {
		resultForm, err := item.RequestFunc(ts.URL)

		if err != nil {
			t.Errorf("[%d] unexpected error : %#v", caseNum, err)
		}

		if !reflect.DeepEqual(resultForm, item.Result) {
			t.Errorf("[%d] wrong result, expected %#v, got %#v", caseNum, item.Result, resultForm)
		}
	}

	db.MyUserStorage.Delete(1)
}

func TestUserErrorCase(t *testing.T) {
	config.LoadConfig("../../config/config.toml")
	router := mux.NewRouter()
	router.HandleFunc("/{id:[0-9]+}", GetUser)
	ts := httptest.NewServer(router)
	defer ts.Close()

	cases := []UserErrorTestCase{
		{ // несуществующий id
			Status: http.StatusNotFound,
			Result: &forms.Result{
				Status: "FAIL",
				Error:  "user not found",
			},
			RequestFunc: func(url string) (*forms.Result, error) {
				resp, err := doGetRequest(url+"/10", nil, map[string]string{})
				if err != nil {
					return nil, err
				}

				return getResultFormFromJson(resp.Body)
			},
		},
	}

	for caseNum, item := range cases {
		resultForm, err := item.RequestFunc(ts.URL)

		if err != nil {
			t.Errorf("[%d] unexpected error : %#v", caseNum, err)
		}

		if !reflect.DeepEqual(resultForm, item.Result) {
			t.Errorf("[%d] wrong result, expected %#v, got %#v", caseNum, item.Result, resultForm)
		}
	}

	//db.MyUserStorage.Delete(1)
}

func TestUserSuccessCase(t *testing.T) {
	config.LoadConfig("../../config/config.toml")
	router := mux.NewRouter()
	router.HandleFunc("/{id:[0-9]+}", GetUser)
	ts := httptest.NewServer(router)
	defer ts.Close()

	cases := []UserSuccessTestCase{
		{ // просто получение пользователя
			Status: http.StatusNotFound,
			Result: &UserGetResponse{
				Status: "OK",
				Result: db.User{
					ID:       1,
					Username: "admin",
					Email:    "admin@samsabaka.ru",
					Password: "",
				},
			},
			RequestFunc: func(url string) (*UserGetResponse, error) {
				db.MyUserStorage.Insert(&db.User{
					ID:       1,
					Username: "admin",
					Email:    "admin@samsabaka.ru",
					Password: "admin",
				})

				resp, err := doGetRequest(url+"/1", nil, map[string]string{})
				if err != nil {
					return nil, err
				}

				return getUserFromJson(resp.Body)
			},
		},
	}

	for caseNum, item := range cases {
		resultForm, err := item.RequestFunc(ts.URL)

		if err != nil {
			t.Errorf("[%d] unexpected error : %#v", caseNum, err)
		}

		if !reflect.DeepEqual(resultForm, item.Result) {
			t.Errorf("[%d] wrong result, expected %#v, got %#v", caseNum, item.Result, resultForm)
		}
	}

	db.MyUserStorage.Delete(1)
}
