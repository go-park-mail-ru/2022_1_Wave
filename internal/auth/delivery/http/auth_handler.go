package auth_http

import (
	"errors"
	auth_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/auth"
	auth_usecase "github.com/go-park-mail-ru/2022_1_Wave/internal/auth/usecase"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	user_microservice_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/user"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/tools/utils"
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
	AuthUseCase auth_domain.AuthUseCase
}

//var Handler AuthHandler
//var M *http_middleware.HttpMiddleware

var csrfTokenExpire = time.Hour * 1

func formSessionCookie(sessionId string) *http.Cookie {
	return &http.Cookie{
		Name:     SessionIdKey,
		Value:    sessionId,
		Expires:  time.Now().Add(auth_usecase.SessionExpires),
		HttpOnly: true,
	}
}

func formCSRFCookie(csrfToken string) *http.Cookie {
	return &http.Cookie{
		Name:     CsrfTokenKey,
		Value:    csrfToken,
		HttpOnly: true,
	}
}

func deleteCookie(sessionId string) *http.Cookie {
	return &http.Cookie{
		Name:     SessionIdKey,
		Value:    sessionId,
		Expires:  time.Now().Add(-auth_usecase.SessionExpires),
		HttpOnly: true,
	}
}

func MakeHandler(authUseCase auth_domain.AuthUseCase) AuthHandler {
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
// @Param        X-CSRF-TOKEN header string true "csrf-token"
// @Param        Cookie header string true "the same csrf-token by key X-CSRF-TOKEN"
// @Success      200    {object}  webUtils.Success
// @Failure      422    {object}  webUtils.Error  "invalid json"
// @Failure      400    {object}  webUtils.Error  "invalid login or password"
// @Router       /api/v1/login [post]
func (a *AuthHandler) Login(c echo.Context) error {
	var user domain.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, getErrorAuthResponse(errors.New(invalidUserJSON)))
	}

	var sessionId string
	if user.Username != "" {
		sessionId, err = a.AuthUseCase.Login(user.Username, user.Password)
	} else {
		sessionId, err = a.AuthUseCase.Login(user.Email, user.Password)
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, getErrorAuthResponse(err))
	}

	c.SetCookie(formSessionCookie(sessionId))

	return c.JSON(http.StatusOK, getSuccessLoginResponse())
}

// Logout godoc
// @Summary      Logout
// @Description  logout user
// @Tags         auth
// @Accept       application/json
// @Produce      application/json
// @Param        X-CSRF-TOKEN header string true "csrf-token"
// @Param        Cookie header string true "the same csrf-token by key X-CSRF-TOKEN"
// @Success      200    {object}  webUtils.Success
// @Router       /api/v1/logout [post]
func (a *AuthHandler) Logout(c echo.Context) error {
	cookie, _ := c.Cookie(SessionIdKey)

	_ = a.AuthUseCase.Logout(cookie.Value)

	c.SetCookie(deleteCookie(cookie.Value))

	return c.JSON(http.StatusOK, getSuccessLogoutResponse())
}

// SignUp godoc
// @Summary      Signup
// @Description  Sign up user
// @Tags         auth
// @Accept       application/json
// @Produce      application/json
// @Param        User body domain.User  true  "Username, Email, Password required"
// @Param        X-CSRF-TOKEN header string true "csrf-token"
// @Param        Cookie header string true "the same csrf-token by key X-CSRF-TOKEN"
// @Success      200    {object}  webUtils.Success
// @Failure      422    {object}  webUtils.Error  "invalid json"
// @Failure      400    {object}  webUtils.Error  "invalid sign up"
// @Router       /api/v1/signup [post]
func (a *AuthHandler) SignUp(c echo.Context) error {
	var user user_microservice_domain.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, getErrorAuthResponse(errors.New(invalidUserJSON)))
	}

	sessionId, err := a.AuthUseCase.SignUp(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, getErrorAuthResponse(err))
	}

	c.SetCookie(formSessionCookie(sessionId))

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
// @Router       /api/v1/get_csrf [get]
func (a *AuthHandler) GetCSRF(c echo.Context) error {
	csrfToken := utils.CreateCSRF()
	c.Response().Header().Set(echo.HeaderXCSRFToken, csrfToken)
	c.SetCookie(formCSRFCookie(csrfToken))

	return c.JSON(http.StatusOK, getSuccessGetCSRFResponse())
}
