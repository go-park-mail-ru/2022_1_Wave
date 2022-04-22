package auth_http

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2022_1_Wave/config"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/auth/usecase"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/auth/proto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/tools/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

const (
	invalidUserJSON = "invalid json"
	sessionIdKey    = "session_id"
)

type AuthHandler struct {
	AuthUseCase proto.AuthorizationClient
}

//var Handler AuthHandler
//var M *http_middleware.HttpMiddleware

var csrfTokenExpire = time.Hour * 1

func formCookie(sessionId string) *http.Cookie {
	return &http.Cookie{
		Name:     sessionIdKey,
		Value:    sessionId,
		Expires:  time.Now().Add(AuthUseCase.SessionExpire),
		HttpOnly: true,
	}
}

func MakeHandler(authUseCase proto.AuthorizationClient) AuthHandler {
	return AuthHandler{
		AuthUseCase: authUseCase,
	}
}

func updateCookieAndCSRF(c echo.Context, sessionId string) {
	c.SetCookie(formCookie(sessionId))

	csrfToken := utils.CreateCSRF(sessionId, int64(csrfTokenExpire))
	c.Response().Header().Set(echo.HeaderXCSRFToken, csrfToken)
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
	cookie, _ := c.Cookie(config.C.SessionIDKey)
	loginData.Session = &proto.Session{SessionId: cookie.Value}
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

	updateCookieAndCSRF(c, loginResult.NewSession.SessionId)

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
	cookie, _ := c.Cookie(sessionIdKey)

	var session proto.Session
	session.SessionId = cookie.Value

	_, _ = a.AuthUseCase.Logout(context.Background(), &session)

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

	cookie, _ := c.Cookie(sessionIdKey)

	err = a.AuthUseCase.SignUp(&user, cookie.Value)
	if err != nil {
		return c.JSON(http.StatusBadRequest, getErrorAuthResponse(err))
	}

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
	cookie, err := c.Cookie(sessionIdKey)
	var csrfToken string
	if err == nil && a.AuthUseCase.IsSession(cookie.Value) { // уже есть сессия, выставляем csrf как session_id
		csrfToken = utils.CreateCSRF(cookie.Value, int64(csrfTokenExpire))
	} else { // нет сессии - создаем
		sessionId, err := a.AuthUseCase.GetUnauthorizedSession()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, getErrorAuthResponse(err))
		}

		csrfToken = utils.CreateCSRF(sessionId, int64(csrfTokenExpire))
		c.SetCookie(formCookie(sessionId))
	}

	c.Response().Header().Set(echo.HeaderXCSRFToken, csrfToken)
	return c.JSON(http.StatusOK, getSuccessGetCSRFResponse())
}
