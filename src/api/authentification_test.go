package api

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2022_1_Wave/config"
	"github.com/go-park-mail-ru/2022_1_Wave/db"
	"github.com/go-park-mail-ru/2022_1_Wave/forms"
	"github.com/go-park-mail-ru/2022_1_Wave/middleware"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

type TestCase struct {
	RequestFunc func(url string) (*forms.Result, error)
	Status      int
	Result      *forms.Result
}

func getSession(url string) (string, *http.Cookie, error) {
	req, _ := http.NewRequest("GET", url, nil)
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", nil, err
	}

	return resp.Header.Get("X-CSRF-TOKEN"), resp.Cookies()[0], nil
}

func doPostRequest(url string, body io.Reader, cookie *http.Cookie, headers map[string]string) (*http.Response, error) {
	req, _ := http.NewRequest("POST", url, body)
	for key, val := range headers {
		req.Header.Add(key, val)
	}
	if cookie != nil {
		req.AddCookie(cookie)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func getResultFormFromJson(r io.Reader) (*forms.Result, error) {
	body, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	resultForm := &forms.Result{}
	err = json.Unmarshal(body, &resultForm)

	if err != nil {
		return nil, err
	}

	return resultForm, nil
}

func TestLoginErrorCase(t *testing.T) {
	config.LoadConfig("../../config/config.toml")
	ts := httptest.NewServer(http.HandlerFunc(middleware.CSRF(middleware.NotAuth(Login))))
	serverGetCSRF := httptest.NewServer(http.HandlerFunc(GetCSRF))
	defer ts.Close()
	defer serverGetCSRF.Close()

	cases := []TestCase{
		{ // не работает без csrf токена
			Status: http.StatusUnauthorized,
			Result: &forms.Result{
				Status: "FAIL",
				Error:  "invalid csrf",
			},
			RequestFunc: func(url string) (*forms.Result, error) {
				resp, err := doPostRequest(url, nil, nil, map[string]string{})

				if err != nil {
					return nil, err
				}

				return getResultFormFromJson(resp.Body)
			},
		},
		{ // валидные куки и csrf токен, но неправильный пользователь
			Status: http.StatusBadRequest,
			Result: &forms.Result{
				Status: "FAIL",
				Error:  "invalid login or password",
			},
			RequestFunc: func(url string) (*forms.Result, error) {
				csrf, cookie, err := getSession(serverGetCSRF.URL)
				if err != nil {
					return nil, err
				}

				user := `{"username": "not_exist", "password": "user"}`
				resp, err := doPostRequest(url, strings.NewReader(user), cookie, map[string]string{"X-CSRF-TOKEN": csrf})
				if err != nil {
					return nil, err
				}

				return getResultFormFromJson(resp.Body)
			},
		},
		{
			Status: http.StatusBadRequest,
			Result: &forms.Result{
				Status: "FAIL",
				Error:  "invalid json",
			},
			RequestFunc: func(url string) (*forms.Result, error) {
				csrf, cookie, err := getSession(serverGetCSRF.URL)
				if err != nil {
					return nil, err
				}

				user := `{"not vaild: json"}`
				resp, err := doPostRequest(url, strings.NewReader(user), cookie, map[string]string{"X-CSRF-TOKEN": csrf})
				if err != nil {
					return nil, err
				}

				return getResultFormFromJson(resp.Body)
			},
		},
		{ // если уже вошли, повторый вызов с авторизованной сессией вернет ошибку
			Status: http.StatusBadRequest,
			Result: &forms.Result{
				Status: "FAIL",
				Error:  "available only to unauthorized users",
			},
			RequestFunc: func(url string) (*forms.Result, error) {
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

				user := `{"username": "admin", "password": "admin"}`
				resp, err := doPostRequest(url, strings.NewReader(user), cookie, map[string]string{"X-CSRF-TOKEN": csrf})
				if err != nil {
					return nil, err
				}

				resp, err = doPostRequest(url, strings.NewReader(user), cookie, map[string]string{"X-CSRF-TOKEN": csrf})
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
}

func TestLoginSuccessCase(t *testing.T) {
	config.LoadConfig("../../config/config.toml")
	ts := httptest.NewServer(http.HandlerFunc(middleware.CSRF(middleware.NotAuth(Login))))
	serverGetCSRF := httptest.NewServer(http.HandlerFunc(GetCSRF))
	defer ts.Close()
	defer serverGetCSRF.Close()

	cases := []TestCase{
		{ // работает вход по username/password
			Status: http.StatusOK,
			Result: &forms.Result{
				Status: "OK",
				Result: "you are login",
			},
			RequestFunc: func(url string) (*forms.Result, error) {
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

				user := `{"username": "admin", "password": "admin"}`
				resp, err := doPostRequest(url, strings.NewReader(user), cookie, map[string]string{"X-CSRF-TOKEN": csrf})
				if err != nil {
					return nil, err
				}

				return getResultFormFromJson(resp.Body)
			},
		},
		{ // работает вход по email/password
			Status: http.StatusOK,
			Result: &forms.Result{
				Status: "OK",
				Result: "you are login",
			},
			RequestFunc: func(url string) (*forms.Result, error) {
				csrf, cookie, err := getSession(serverGetCSRF.URL)
				if err != nil {
					return nil, err
				}

				db.MyUserStorage.Insert(&db.User{
					ID:       1,
					Username: "admin2",
					Email:    "admin2@samsabaka.ru",
					Password: "admin2",
				})

				user := `{"email": "admin2@samsabaka.ru", "password": "admin2"}`
				resp, err := doPostRequest(url, strings.NewReader(user), cookie, map[string]string{"X-CSRF-TOKEN": csrf})
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
}
