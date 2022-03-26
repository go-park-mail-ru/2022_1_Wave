package http

import (
	"errors"
	"github.com/go-park-mail-ru/2022_1_Wave/config"
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

func NewAuthHandler(e *echo.Echo, authUseCase domain.AuthUseCase) {
	handler := &AuthHandler{
		authUseCase: authUseCase,
	}

	// TODO: навесить мидлвары
	e.POST("/login", handler.Login)
	e.POST("/logout", handler.Logout)
	e.POST("/signup", handler.SignUp)
}

func (a *AuthHandler) Login(c echo.Context) error {
	var user domain.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, getErrorAuthResponse(errors.New(invalidUserJSON)))
	}

	var sessionId string
	if user.Username != "" {
		sessionId, err = a.authUseCase.Login(user.Username, user.Password)
	} else {
		sessionId, err = a.authUseCase.Login(user.Email, user.Password)
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, getErrorAuthResponse(err))
	}

	cookie := &http.Cookie{
		Name:     config.C.SessionIDKey,
		Value:    sessionId,
		Expires:  time.Now().Add(sessionExpire),
		HttpOnly: true,
	}
	c.SetCookie(cookie)

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

	sessionId, err := a.authUseCase.SignUp(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, getErrorAuthResponse(err))
	}

	cookie := &http.Cookie{
		Name:     config.C.SessionIDKey,
		Value:    sessionId,
		Expires:  time.Now().Add(sessionExpire),
		HttpOnly: true,
	}
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, getSuccessSignUpResponse())
}
