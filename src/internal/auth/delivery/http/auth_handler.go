package http

import (
	"errors"
	"github.com/go-park-mail-ru/2022_1_Wave/config"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/auth/delivery/http/middleware"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

const (
	invalidUserJSON = "invalid json"
)

type AuthHandler struct {
	authUseCase domain.AuthUseCase
}

var sessionExpire, _ = time.ParseDuration(config.C.SessionExpires)

func formCookie(sessionId string) *http.Cookie {
	return &http.Cookie{
		Name:     config.C.SessionIDKey,
		Value:    sessionId,
		Expires:  time.Now().Add(sessionExpire),
		HttpOnly: true,
	}
}

func NewAuthHandler(e *echo.Echo, authUseCase domain.AuthUseCase, m *middleware.HttpMiddleware) {
	handler := &AuthHandler{
		authUseCase: authUseCase,
	}

	// TODO: навесить мидлвары
	e.POST("/login", handler.Login, m.IsSession, m.CSRF)
	e.POST("/logout", handler.Logout, m.IsSession, m.CSRF)
	e.POST("/signup", handler.SignUp, m.IsSession, m.CSRF)
	e.POST("/get_csrf", handler.GetCSRF)
}

func (a *AuthHandler) Login(c echo.Context) error {
	var user domain.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, getErrorAuthResponse(errors.New(invalidUserJSON)))
	}

	cookie, _ := c.Cookie(config.C.SessionIDKey)
	if user.Username != "" {
		err = a.authUseCase.Login(user.Username, user.Password, cookie.Value)
	} else {
		err = a.authUseCase.Login(user.Email, user.Password, cookie.Value)
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, getErrorAuthResponse(err))
	}

	return c.JSON(http.StatusOK, getSuccessLoginResponse())
}

func (a *AuthHandler) Logout(c echo.Context) error {
	cookie, _ := c.Cookie(config.C.SessionIDKey)

	_ = a.authUseCase.Logout(cookie.Value)

	return nil
}

func (a *AuthHandler) SignUp(c echo.Context) error {
	var user domain.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, getErrorAuthResponse(errors.New(invalidUserJSON)))
	}

	cookie, _ := c.Cookie(config.C.SessionIDKey)

	err = a.authUseCase.SignUp(&user, cookie.Value)
	if err != nil {
		return c.JSON(http.StatusBadRequest, getErrorAuthResponse(err))
	}

	return c.JSON(http.StatusOK, getSuccessSignUpResponse())
}

func (a *AuthHandler) GetCSRF(c echo.Context) error {
	cookie, err := c.Cookie(config.C.SessionIDKey)
	var csrfToken string
	if err == nil && a.authUseCase.IsSession(cookie.Value) { // уже есть сессия, выставляем csrf как session_id
		csrfToken = cookie.Value
	} else { // нет сессии - создаем
		sessionId, err := a.authUseCase.GetUnauthorizedSession()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, getErrorAuthResponse(err))
		}

		csrfToken = sessionId
		c.SetCookie(formCookie(sessionId))
	}

	c.Response().Header().Set(echo.HeaderXCSRFToken, csrfToken)
	return nil
}
