package http

import (
	"github.com/go-park-mail-ru/2022_1_Wave/config"
	domain2 "github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type AuthHandler struct {
	authUseCase domain2.AuthUseCase
}

var sessionExpire, _ = time.ParseDuration(config.C.SessionExpires)

func NewAuthHandler(e *echo.Echo, authUseCase domain2.AuthUseCase) {
	handler := &AuthHandler{
		authUseCase: authUseCase,
	}
	e.POST("/login", handler.Login)
	e.POST("/logout", handler.Logout)
}

// TODO: функции, реализующие взаимодействие с AuthUseCase (бизнес логикой) посредством http

func (a *AuthHandler) Login(c echo.Context) error {
	var user domain2.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, getErrorAuthResponse(err))
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
	return nil
}
