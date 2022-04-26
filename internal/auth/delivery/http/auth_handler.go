package authHttp

import (
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2022_1_Wave/config"
	AuthUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/auth/usecase"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/tools/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

const (
	invalidUserJSON = "invalid json"
	sessionIdKey    = "session_id"
)

type AuthHandler struct {
	AuthUseCase domain.AuthUseCase
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

func MakeHandler(authUseCase domain.AuthUseCase) AuthHandler {
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

	cookie, _ := c.Cookie(config.C.SessionIDKey)
	if user.Username != "" {
		err = a.AuthUseCase.Login(user.Username, user.Password, cookie.Value)
	} else {
		err = a.AuthUseCase.Login(user.Email, user.Password, cookie.Value)
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, getErrorAuthResponse(err))
	}

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

	_ = a.AuthUseCase.Logout(cookie.Value)

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
