package auth_http_test

import (
	"encoding/json"
	"errors"
	"github.com/bxcodec/faker"
	auth_http "github.com/go-park-mail-ru/2022_1_Wave/internal/auth/delivery/http"
	mocks2 "github.com/go-park-mail-ru/2022_1_Wave/internal/auth/mocks"
	user_microservice_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/user"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
	var mockUser1 user_microservice_domain.User
	var mockUser2 user_microservice_domain.User
	err := faker.FakeData(&mockUser1)
	assert.NoError(t, err)
	mockUser1.Email = ""

	err = faker.FakeData(&mockUser2)
	assert.NoError(t, err)
	mockUser2.Username = ""

	sessionId := "some-session-id"

	mockUseCase := new(mocks2.AuthUseCase)

	mockUseCase.On("Login", mockUser1.Username, mockUser1.Password).Return(sessionId, nil)

	mockUseCase.On("Login", mockUser1.Username+"a", mockUser1.Password).Return("", errors.New("error"))

	mockUseCase.On("Login", mockUser2.Email, mockUser2.Password).Return(sessionId, nil)

	e := echo.New()
	jsonUser, _ := json.Marshal(&mockUser1)
	req, err := http.NewRequest(echo.POST, "/login", strings.NewReader(string(jsonUser)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/login")
	handler := auth_http.AuthHandler{
		AuthUseCase: mockUseCase,
	}

	err = handler.Login(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	jsonUser, _ = json.Marshal(&mockUser2)
	req, err = http.NewRequest(echo.POST, "/login", strings.NewReader(string(jsonUser)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	//req.AddCookie(cookie)
	rec = httptest.NewRecorder()

	c = e.NewContext(req, rec)
	c.SetPath("/login")

	err = handler.Login(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEqual(t, rec.Header().Get("Set-Cookie"), "")

	mockUser1.Username += "a"
	jsonUser, _ = json.Marshal(&mockUser1)
	req, err = http.NewRequest(echo.POST, "/login", strings.NewReader(string(jsonUser)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	//req.AddCookie(cookie)
	rec = httptest.NewRecorder()

	c = e.NewContext(req, rec)
	c.SetPath("/login")

	err = handler.Login(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	req, err = http.NewRequest(echo.POST, "/login", strings.NewReader("aboba"))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	//req.AddCookie(cookie)
	rec = httptest.NewRecorder()

	c = e.NewContext(req, rec)
	c.SetPath("/login")

	err = handler.Login(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
}

func TestLogout(t *testing.T) {
	sessionId := "some-session-id"

	mockUseCase := new(mocks2.AuthUseCase)
	mockUseCase.On("Logout", sessionId).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/logout", strings.NewReader(""))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sessionId,
		HttpOnly: true,
	}
	req.AddCookie(cookie)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/logout")
	handler := auth_http.AuthHandler{
		AuthUseCase: mockUseCase,
	}

	err = handler.Logout(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestSignUp(t *testing.T) {
	var mockUser1 user_microservice_domain.User
	var mockUser2 user_microservice_domain.User
	err := faker.FakeData(&mockUser1)
	assert.NoError(t, err)
	mockUser1.Email = ""

	err = faker.FakeData(&mockUser2)
	assert.NoError(t, err)
	mockUser2.Username = ""

	sessionId := "some-session-id"
	mockUseCase := new(mocks2.AuthUseCase)
	mockUseCase.On("SignUp", &mockUser1).Return(sessionId, nil)
	mockUseCase.On("SignUp", &mockUser2).Return("", errors.New("error"))

	e := echo.New()
	jsonUser, _ := json.Marshal(&mockUser1)
	req, err := http.NewRequest(echo.POST, "/signup", strings.NewReader(string(jsonUser)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/signup")
	handler := auth_http.AuthHandler{
		AuthUseCase: mockUseCase,
	}

	err = handler.SignUp(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEqual(t, rec.Header().Get("Set-Cookie"), "")

	jsonUser, _ = json.Marshal(&mockUser2)
	req, err = http.NewRequest(echo.POST, "/signup", strings.NewReader(string(jsonUser)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec = httptest.NewRecorder()

	c = e.NewContext(req, rec)
	c.SetPath("/signup")

	err = handler.SignUp(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, rec.Header().Get("Set-Cookie"), "")

	req, err = http.NewRequest(echo.POST, "/signup", strings.NewReader("aboba"))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec = httptest.NewRecorder()

	c = e.NewContext(req, rec)
	c.SetPath("/signup")

	err = handler.SignUp(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	assert.Equal(t, rec.Header().Get("Set-Cookie"), "")
}

func TestGetCSRF(t *testing.T) {
	mockUseCase := new(mocks2.AuthUseCase)

	e := echo.New()

	req, err := http.NewRequest(echo.POST, "/get_csrf", strings.NewReader(""))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/get_csrf")

	handler := auth_http.AuthHandler{
		AuthUseCase: mockUseCase,
	}

	err = handler.GetCSRF(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEqual(t, rec.Header().Get(echo.HeaderXCSRFToken), "")
	assert.Equal(t, rec.Header().Get(echo.HeaderXCSRFToken), rec.Result().Cookies()[0].Value)
}
