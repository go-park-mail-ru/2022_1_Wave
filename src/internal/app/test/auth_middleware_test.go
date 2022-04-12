package test

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/auth/delivery/http/http_middleware"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain/mocks"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/tools/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestIsSession(t *testing.T) {
	result := false
	check := func(c echo.Context) error {
		result = true
		return nil
	}

	sessionId := "some-session-id"
	mockUseCase := new(mocks.AuthUseCase)
	mockUseCase.On("IsSession", sessionId).Return(true)

	e := echo.New()
	req, _ := http.NewRequest(echo.GET, "/is_session", strings.NewReader(""))
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sessionId,
		HttpOnly: true,
	}
	req.AddCookie(cookie)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/is_session")

	handler := http_middleware.InitMiddleware(mockUseCase)
	res := handler.IsSession(check)
	err := res(c)

	assert.NoError(t, err)
	assert.True(t, result)

	result = false

	req, _ = http.NewRequest(echo.GET, "/is_session", strings.NewReader(""))

	rec = httptest.NewRecorder()

	c = e.NewContext(req, rec)
	c.SetPath("/is_session")

	err = res(c)
	assert.False(t, result)
}

func TestCSRF(t *testing.T) {
	result := false
	check := func(c echo.Context) error {
		result = true
		return nil
	}

	sessionId := "some-session-id"
	mockUseCase := new(mocks.AuthUseCase)
	mockUseCase.On("IsSession", sessionId).Return(true)

	e := echo.New()
	req, _ := http.NewRequest(echo.GET, "/is_session", strings.NewReader(""))
	req.Header.Set(echo.HeaderXCSRFToken, utils.CreateCSRF(sessionId, int64(time.Hour)))
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sessionId,
		HttpOnly: true,
	}
	req.AddCookie(cookie)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/is_session")

	handler := http_middleware.InitMiddleware(mockUseCase)
	res := handler.CSRF(check)
	err := res(c)

	assert.NoError(t, err)
	assert.True(t, result)

	result = false

	req, _ = http.NewRequest(echo.GET, "/is_session", strings.NewReader(""))

	rec = httptest.NewRecorder()

	c = e.NewContext(req, rec)
	c.SetPath("/is_session")

	err = res(c)
	assert.False(t, result)

	result = false

	req, _ = http.NewRequest(echo.GET, "/is_session", strings.NewReader(""))
	cookie = &http.Cookie{
		Name:     "session_id",
		Value:    sessionId,
		HttpOnly: true,
	}
	req.AddCookie(cookie)
	rec = httptest.NewRecorder()

	c = e.NewContext(req, rec)
	c.SetPath("/is_session")

	err = res(c)
	assert.False(t, result)
}

func TestAuth(t *testing.T) {
	result := false
	check := func(c echo.Context) error {
		result = true
		return nil
	}

	sessionId := "some-session-id"
	mockUseCase := new(mocks.AuthUseCase)
	mockUseCase.On("IsAuthSession", sessionId).Return(true)
	mockUseCase.On("IsAuthSession", sessionId+"a").Return(false)

	e := echo.New()
	req, _ := http.NewRequest(echo.GET, "/is_session", strings.NewReader(""))
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sessionId,
		HttpOnly: true,
	}
	req.AddCookie(cookie)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/is_session")

	handler := http_middleware.InitMiddleware(mockUseCase)
	res := handler.Auth(check)
	err := res(c)

	assert.NoError(t, err)
	assert.True(t, result)

	result = false

	req, _ = http.NewRequest(echo.GET, "/is_session", strings.NewReader(""))
	cookie = &http.Cookie{
		Name:     "session_id",
		Value:    sessionId + "a",
		HttpOnly: true,
	}
	req.AddCookie(cookie)
	rec = httptest.NewRecorder()

	c = e.NewContext(req, rec)
	c.SetPath("/is_session")

	err = res(c)
	assert.False(t, result)
}

func TestNotAuth(t *testing.T) {
	result := false
	check := func(c echo.Context) error {
		result = true
		return nil
	}

	sessionId := "some-session-id"
	mockUseCase := new(mocks.AuthUseCase)
	mockUseCase.On("IsSession", sessionId).Return(true)
	mockUseCase.On("IsSession", sessionId+"a").Return(true)
	mockUseCase.On("IsAuthSession", sessionId).Return(true)
	mockUseCase.On("IsAuthSession", sessionId+"a").Return(false)

	e := echo.New()
	req, _ := http.NewRequest(echo.GET, "/is_session", strings.NewReader(""))
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sessionId,
		HttpOnly: true,
	}
	req.AddCookie(cookie)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/is_session")

	handler := http_middleware.InitMiddleware(mockUseCase)
	res := handler.NotAuth(check)
	_ = res(c)

	assert.False(t, result)

	result = false

	req, _ = http.NewRequest(echo.GET, "/is_session", strings.NewReader(""))
	cookie = &http.Cookie{
		Name:     "session_id",
		Value:    sessionId + "a",
		HttpOnly: true,
	}
	req.AddCookie(cookie)
	rec = httptest.NewRecorder()

	c = e.NewContext(req, rec)
	c.SetPath("/is_session")

	_ = res(c)
	assert.True(t, result)
}
