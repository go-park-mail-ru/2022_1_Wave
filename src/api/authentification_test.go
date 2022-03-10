package api

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2022_1_Wave/config"
	"github.com/go-park-mail-ru/2022_1_Wave/db/models"
	"github.com/go-park-mail-ru/2022_1_Wave/forms"
	"github.com/go-park-mail-ru/2022_1_Wave/middleware"
	"github.com/go-park-mail-ru/2022_1_Wave/service"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

type AuthTestCase struct {
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

func doGetRequest(url string, cookie *http.Cookie, headers map[string]string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", url, nil)
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

	cases := []AuthTestCase{
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

				models.MyUserStorage.Insert(&models.User{
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

	models.MyUserStorage.Delete(1)
}

func TestLoginSuccessCase(t *testing.T) {
	config.LoadConfig("../../config/config.toml")
	ts := httptest.NewServer(http.HandlerFunc(middleware.CSRF(middleware.NotAuth(Login))))
	serverGetCSRF := httptest.NewServer(http.HandlerFunc(GetCSRF))
	defer ts.Close()
	defer serverGetCSRF.Close()

	cases := []AuthTestCase{
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

				models.MyUserStorage.Insert(&models.User{
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

				models.MyUserStorage.Insert(&models.User{
					ID:       2,
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

	models.MyUserStorage.Delete(1)
	models.MyUserStorage.Delete(2)
}

func TestSignUpErrorCase(t *testing.T) {
	config.LoadConfig("../../config/config.toml")
	ts := httptest.NewServer(http.HandlerFunc(SignUp))
	defer ts.Close()

	cases := []AuthTestCase{
		{ // не работает для невалидного формата данных
			Status: http.StatusBadRequest,
			Result: &forms.Result{
				Status: "FAIL",
				Error:  "invalid json",
			},
			RequestFunc: func(url string) (*forms.Result, error) {
				user := `{"username: "admin", "email": "admin@samsabaka.ru", "password": "admin"}`
				resp, err := doPostRequest(url, strings.NewReader(user), nil, map[string]string{})
				if err != nil {
					return nil, err
				}

				return getResultFormFromJson(resp.Body)
			},
		},
		{ // не работает для невалидных данных
			Status: http.StatusBadRequest,
			Result: &forms.Result{
				Status: "FAIL",
				Error:  "invalid fields",
			},
			RequestFunc: func(url string) (*forms.Result, error) {
				models.MyUserStorage.Insert(&models.User{
					ID:       1,
					Username: "admin",
					Email:    "admin@samsabaka",
					Password: "admin",
				})

				user := `{"username": "admin", "email": "admin@samsabaka", "password": "admin"}`
				resp, err := doPostRequest(url, strings.NewReader(user), nil, map[string]string{})
				if err != nil {
					return nil, err
				}

				return getResultFormFromJson(resp.Body)
			},
		},
		{ // не работает для уже существующего пользователя
			Status: http.StatusBadRequest,
			Result: &forms.Result{
				Status: "FAIL",
				Error:  "user already exist",
			},
			RequestFunc: func(url string) (*forms.Result, error) {
				models.MyUserStorage.Insert(&models.User{
					ID:       1,
					Username: "admin",
					Email:    "admin@samsabaka.ru",
					Password: "admin",
				})

				user := `{"username": "admin", "email": "admin2@samsabaka.ru", "password": "admin"}`
				resp, err := doPostRequest(url, strings.NewReader(user), nil, map[string]string{})
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

	models.MyUserStorage.Delete(1)
}

func TestSignUpSuccessCase(t *testing.T) {
	config.LoadConfig("../../config/config.toml")
	ts := httptest.NewServer(http.HandlerFunc(middleware.CSRF(middleware.NotAuth(SignUp))))
	serverGetCSRF := httptest.NewServer(http.HandlerFunc(GetCSRF))
	defer ts.Close()
	defer serverGetCSRF.Close()

	var signupCookie *http.Cookie
	var signupCSRFToken string

	cases := []AuthTestCase{
		{ // просто работает регистрация
			Status: http.StatusOK,
			Result: &forms.Result{
				Status: "OK",
				Result: "you are sign up",
			},
			RequestFunc: func(url string) (*forms.Result, error) {
				csrf, cookie, err := getSession(serverGetCSRF.URL)
				if err != nil {
					return nil, err
				}

				user := `{"username": "admin2", "email": "admin2@samsabaka.ru", "password": "admin2"}`
				resp, err := doPostRequest(url, strings.NewReader(user), cookie, map[string]string{"X-CSRF-TOKEN": csrf})
				if err != nil {
					return nil, err
				}

				signupCookie = cookie
				signupCSRFToken = csrf

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

	session, ok := service.Sessions[signupCookie.Value]
	require.True(t, ok)
	require.Equal(t, session.CSRFToken, signupCSRFToken)
	require.True(t, session.IsAuthorized)

	models.MyUserStorage.Delete(1)
}

func TestLogoutErrorCase(t *testing.T) {
	config.LoadConfig("../../config/config.toml")
	ts := httptest.NewServer(http.HandlerFunc(middleware.CSRF(middleware.Auth(Logout))))
	serverGetCSRF := httptest.NewServer(http.HandlerFunc(GetCSRF))
	defer ts.Close()
	defer serverGetCSRF.Close()

	cases := []AuthTestCase{
		{ // не работает без cookie и csrf токена
			Status: http.StatusUnauthorized,
			Result: &forms.Result{
				Status: "FAIL",
				Error:  "invalid csrf",
			},
			RequestFunc: func(url string) (*forms.Result, error) {

				resp, err := doPostRequest(url, strings.NewReader(""), nil, map[string]string{})
				if err != nil {
					return nil, err
				}

				return getResultFormFromJson(resp.Body)
			},
		},
		{ // не работает для не авторизованной сессии
			Status: http.StatusUnauthorized,
			Result: &forms.Result{
				Status: "FAIL",
				Error:  "unauthorized",
			},
			RequestFunc: func(url string) (*forms.Result, error) {
				csrf, cookie, err := getSession(serverGetCSRF.URL)
				if err != nil {
					return nil, err
				}

				resp, err := doPostRequest(url, strings.NewReader(""), cookie, map[string]string{"X-CSRF-TOKEN": csrf})

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

func TestLogoutSuccessCase(t *testing.T) {
	config.LoadConfig("../../config/config.toml")
	ts := httptest.NewServer(http.HandlerFunc(middleware.CSRF(middleware.Auth(Logout))))
	serverGetCSRF := httptest.NewServer(http.HandlerFunc(GetCSRF))
	defer ts.Close()
	defer serverGetCSRF.Close()

	var logoutCookie *http.Cookie

	cases := []AuthTestCase{
		{ // работает для авторизованной сессии
			Status: http.StatusOK,
			Result: &forms.Result{
				Status: "OK",
				Result: "you are logout",
			},
			RequestFunc: func(url string) (*forms.Result, error) {
				csrf, cookie, err := getSession(serverGetCSRF.URL)
				if err != nil {
					return nil, err
				}

				models.MyUserStorage.Insert(&models.User{
					ID:       1,
					Username: "admin3",
					Email:    "admin3@samsabaka.ru",
					Password: "admin3",
				})

				service.AuthorizeUser(cookie.Value, 1)

				logoutCookie = cookie

				resp, err := doPostRequest(url, strings.NewReader(""), cookie, map[string]string{"X-CSRF-TOKEN": csrf})

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

	// сессия удаляется из хранилища после логаута
	_, ok := service.Sessions[logoutCookie.Value]
	require.False(t, ok)

	models.MyUserStorage.Delete(1)
}
