package auth_http

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/auth/proto"
	auth_usecase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/auth/usecase"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/tools/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

const (
	invalidUserJSON = "invalid json"
	SessionIdKey    = "session_id"
	CsrfTokenKey    = echo.HeaderXCSRFToken
)

type AuthHandler struct {
	AuthUseCase proto.AuthorizationClient
}

//var Handler AuthHandler
//var M *http_middleware.HttpMiddleware

var csrfTokenExpire = time.Hour * 1

func formSessionCookie(sessionId string) *http.Cookie {
	return &http.Cookie{
		Name:     SessionIdKey,
		Value:    sessionId,
		Expires:  time.Now().Add(auth_usecase.SessionExpire),
		HttpOnly: true,
	}
}

func formCSRFCookie(csrfToken string) *http.Cookie {
	return &http.Cookie{
		Name:     CsrfTokenKey,
		Value:    csrfToken,
		Expires:  time.Now().Add(csrfTokenExpire),
		HttpOnly: true,
	}
}

func deleteCookie(sessionId string) *http.Cookie {
	return &http.Cookie{
		Name:     SessionIdKey,
		Value:    sessionId,
		Expires:  time.Now().Add(-auth_usecase.SessionExpire),
		HttpOnly: true,
	}
}

func MakeHandler(authUseCase proto.AuthorizationClient) AuthHandler {
	return AuthHandler{
		AuthUseCase: authUseCase,
	}
}

// Login godoc
// @Summary      Login
// @Description  login user
// @Tags         auth
// @Accept       application/json
// @Produce      application/json
// @Param        User body domain.User  true  "username/email and password"
// @Success      200    {object}  webUtils.Success
// @Failure      422    {object}  webUtils.Error  "invalid json"
// @Failure      400    {object}  webUtils.Error  "invalid login or password"
// @Router       /api/v1/login/ [post]
func (a *AuthHandler) Login(c echo.Context) error {
	var user domain.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, getErrorAuthResponse(errors.New(invalidUserJSON)))
	}

	var loginResult *proto.LoginResult
	var loginData proto.LoginData

	loginData.Password = user.Password
	if user.Username != "" {
		loginData.Login = user.Username
		loginResult, err = a.AuthUseCase.Login(context.Background(), &loginData)
	} else {
		loginData.Login = user.Email
		loginResult, err = a.AuthUseCase.Login(context.Background(), &loginData)
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, getErrorAuthResponse(err))
	}

	c.SetCookie(formSessionCookie(loginResult.NewSession.SessionId))

	return c.JSON(http.StatusOK, getSuccessLoginResponse())
}

// Logout godoc
// @Summary      Logout
// @Description  logout user
// @Tags         auth
// @Accept       application/json
// @Produce      application/json
// @Success      200    {object}  webUtils.Success
// @Router       /api/v1/logout/ [post]
func (a *AuthHandler) Logout(c echo.Context) error {
	cookie, _ := c.Cookie(SessionIdKey)

	var session proto.Session
	session.SessionId = cookie.Value

	_, _ = a.AuthUseCase.Logout(context.Background(), &session)

	c.SetCookie(deleteCookie(session.SessionId))

	return c.JSON(http.StatusOK, getSuccessLogoutResponse())
}

// SignUp godoc
// @Summary      Signup
// @Description  Sign up user
// @Tags         auth
// @Accept       application/json
// @Produce      application/json
// @Param        User body domain.User  true  "Username, Email, Password required"
// @Success      200    {object}  webUtils.Success
// @Failure      422    {object}  webUtils.Error  "invalid json"
// @Failure      400    {object}  webUtils.Error  "invalid sign up"
// @Router       /api/v1/signup/ [post]
func (a *AuthHandler) SignUp(c echo.Context) error {
	var user domain.User
	err := c.Bind(&user)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusUnprocessableEntity, getErrorAuthResponse(errors.New(invalidUserJSON)))
	}

	var data proto.SignUpData
	data.User = &proto.User{Username: user.Username, Email: user.Email, Password: user.Password}

	signUpResult, err := a.AuthUseCase.SignUp(context.Background(), &data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, getErrorAuthResponse(err))
	}

	c.SetCookie(formSessionCookie(signUpResult.NewSession.SessionId))

	return c.JSON(http.StatusOK, getSuccessSignUpResponse())
}

// GetCSRF godoc
// @Summary      Getting csrf
// @Description  Exposes a csrf token and creates an unauthorized session
// @Tags         auth
// @Accept       application/json
// @Produce      application/json
// @Success      200    {object}  webUtils.Success
// @Failure      500    {object}  webUtils.Error  "Internal server error"
// @Router       /api/v1/get_csrf/ [post]
func (a *AuthHandler) GetCSRF(c echo.Context) error {
	csrfToken := utils.CreateCSRF()
	c.Response().Header().Set(echo.HeaderXCSRFToken, csrfToken)
	c.SetCookie(formCSRFCookie(csrfToken))

	return c.JSON(http.StatusOK, getSuccessGetCSRFResponse())
}
