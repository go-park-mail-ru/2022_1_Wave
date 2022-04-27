package auth_http_test

import (
	"context"
	"encoding/json"
	"github.com/bxcodec/faker"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	auth_http "github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/auth/delivery/http"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/auth/mocks"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/auth/proto"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
	var mockUser1 domain.User
	var mockUser2 domain.User
	err := faker.FakeData(&mockUser1)
	assert.NoError(t, err)
	mockUser1.Email = ""

	err = faker.FakeData(&mockUser2)
	assert.NoError(t, err)
	mockUser2.Username = ""

	sessionId := "some-session-id"

	mockUseCase := new(mocks.AuthorizationClient)
	var loginData proto.LoginData
	loginData.Login = mockUser1.Username
	loginData.Password = mockUser1.Password

	mockUseCase.On("Login", context.Background(), &loginData).Return(&proto.LoginResult{NewSession: &proto.Session{SessionId: sessionId}}, nil)
	loginData.Login += "a"

	mockUseCase.On("Login", context.Background(), &loginData).Return(&proto.LoginResult{NewSession: &proto.Session{SessionId: sessionId}}, nil)

	loginData.Login = mockUser2.Email
	loginData.Password = mockUser2.Password
	mockUseCase.On("Login", context.Background(), &loginData).Return(&proto.LoginResult{NewSession: &proto.Session{SessionId: sessionId}}, nil)

	e := echo.New()
	jsonUser, _ := json.Marshal(&mockUser1)
	req, err := http.NewRequest(echo.POST, "/login", strings.NewReader(string(jsonUser)))
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

	req.AddCookie(cookie)
	rec = httptest.NewRecorder()

	c = e.NewContext(req, rec)
	c.SetPath("/login")

	err = handler.Login(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	mockUser1.Username += "a"
	jsonUser, _ = json.Marshal(&mockUser1)
	req, err = http.NewRequest(echo.POST, "/login", strings.NewReader(string(jsonUser)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	req.AddCookie(cookie)
	rec = httptest.NewRecorder()

	c = e.NewContext(req, rec)
	c.SetPath("/login")

	err = handler.Login(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	req, err = http.NewRequest(echo.POST, "/login", strings.NewReader("aboba"))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	req.AddCookie(cookie)
	rec = httptest.NewRecorder()

	c = e.NewContext(req, rec)
	c.SetPath("/login")

	err = handler.Login(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
}

/*
func TestLogout(t *testing.T) {
	sessionId := "some-session-id"

	mockUseCase := new(mocks.AuthUseCase)
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
	handler := authHttp.AuthHandler{
		AuthUseCase: mockUseCase,
	}

	err = handler.Logout(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestSignUp(t *testing.T) {
	var mockUser1 domain.User
	var mockUser2 domain.User
	err := faker.FakeData(&mockUser1)
	assert.NoError(t, err)
	mockUser1.Email = ""

	err = faker.FakeData(&mockUser2)
	assert.NoError(t, err)
	mockUser2.Username = ""

	sessionId := "some-session-id"

	mockUseCase := new(mocks.AuthUseCase)
	mockUseCase.On("SignUp", &mockUser1, sessionId).Return(nil)
	mockUseCase.On("SignUp", &mockUser2, sessionId).Return(errors.New("error"))

	e := echo.New()
	jsonUser, _ := json.Marshal(&mockUser1)
	req, err := http.NewRequest(echo.POST, "/signup", strings.NewReader(string(jsonUser)))
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
	c.SetPath("/signup")
	handler := authHttp.AuthHandler{
		AuthUseCase: mockUseCase,
	}

	err = handler.SignUp(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	jsonUser, _ = json.Marshal(&mockUser2)
	req, err = http.NewRequest(echo.POST, "/signup", strings.NewReader(string(jsonUser)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	req.AddCookie(cookie)
	rec = httptest.NewRecorder()

	c = e.NewContext(req, rec)
	c.SetPath("/signup")

	err = handler.SignUp(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	req, err = http.NewRequest(echo.POST, "/signup", strings.NewReader("aboba"))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	req.AddCookie(cookie)
	rec = httptest.NewRecorder()

	c = e.NewContext(req, rec)
	c.SetPath("/signup")

	err = handler.SignUp(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
}

func TestGetCSRF(t *testing.T) {
	sessionId1 := "some-session-id"
	sessionId2 := "some-session-id2"

	mockUseCase := new(mocks.AuthUseCase)
	mockUseCase.On("IsSession", sessionId1).Return(true)
	mockUseCase.On("IsSession", sessionId2).Return(false)
	mockUseCase.On("GetUnauthorizedSession").Return(sessionId2, nil)

	e := echo.New()

	req, err := http.NewRequest(echo.POST, "/get_csrf", strings.NewReader(""))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sessionId1,
		HttpOnly: true,
	}
	req.AddCookie(cookie)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/get_csrf")

	handler := authHttp.AuthHandler{
		AuthUseCase: mockUseCase,
	}

	err = handler.GetCSRF(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEqual(t, rec.Header().Get(echo.HeaderXCSRFToken), "")
}
*/
