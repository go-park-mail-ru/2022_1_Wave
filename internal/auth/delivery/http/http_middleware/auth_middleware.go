package http_middleware

import (
	"errors"
	"github.com/go-park-mail-ru/2022_1_Wave/config"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/tools/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type HttpMiddleware struct {
	authUseCase domain.AuthUseCase
}

type middlewareResponse struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

const (
	statusFAIL = "FAIL"
)

var (
	ErrNoSession   = errors.New("your don't have a session")
	ErrInvalidCSRF = errors.New("invalid csrf")
	ErrNotAuth     = errors.New("available only to authorized users")
	ErrAuth        = errors.New("available only to unauthorized users")
)

func InitMiddleware(authUseCase domain.AuthUseCase) *HttpMiddleware {
	return &HttpMiddleware{authUseCase: authUseCase}
}

func getErrorMiddlewareResponse(err error) *middlewareResponse {
	return &middlewareResponse{
		Status: statusFAIL,
		Error:  err.Error(),
	}
}

func (m *HttpMiddleware) IsSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie(config.C.SessionIDKey)
		if err != nil || !m.authUseCase.IsSession(cookie.Value) {
			return c.JSON(http.StatusUnauthorized, getErrorMiddlewareResponse(ErrNoSession))
		}

		return next(c)
	}
}

func (m *HttpMiddleware) CSRF(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie(config.C.SessionIDKey)
		if err != nil || !m.authUseCase.IsSession(cookie.Value) {
			return c.JSON(http.StatusUnauthorized, getErrorMiddlewareResponse(ErrNoSession))
		}

		csrf := c.Request().Header.Get(echo.HeaderXCSRFToken)
		if !utils.CheckCSRF(cookie.Value, csrf) {
			return c.JSON(http.StatusUnauthorized, getErrorMiddlewareResponse(ErrInvalidCSRF))
		}

		return next(c)
	}
}

func (m *HttpMiddleware) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie(config.C.SessionIDKey)
		if err != nil || !m.authUseCase.IsAuthSession(cookie.Value) {
			return c.JSON(http.StatusUnauthorized, getErrorMiddlewareResponse(ErrNotAuth))
		}

		return next(c)
	}
}

func (m *HttpMiddleware) NotAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie(config.C.SessionIDKey)
		if err != nil || !m.authUseCase.IsSession(cookie.Value) || m.authUseCase.IsAuthSession(cookie.Value) {
			return c.JSON(http.StatusUnauthorized, getErrorMiddlewareResponse(ErrAuth))
		}

		return next(c)
	}
}
