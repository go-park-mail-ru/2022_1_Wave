package middleware

import (
	"errors"
	"fmt"
	auth_domain "github.com/go-park-mail-ru/2022_1_Wave/websocket-server/auth"
	"github.com/labstack/echo/v4"
	"net/http"
)

type HttpMiddleware struct {
	authAgent auth_domain.AuthAgent
}

type middlewareResponse struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

const (
	statusFAIL = "FAIL"
)

var (
	ErrNotAuth = errors.New("available only to authorized users")
)

func InitMiddleware(authAgent auth_domain.AuthAgent) *HttpMiddleware {
	return &HttpMiddleware{authAgent: authAgent}
}

func getErrorMiddlewareResponse(err error) *middlewareResponse {
	return &middlewareResponse{
		Status: statusFAIL,
		Error:  err.Error(),
	}
}

func (m *HttpMiddleware) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("session_id")
		fmt.Println("cookie = ", cookie)
		if err != nil || !m.authAgent.IsAuthSession(cookie.Value) {
			return c.JSON(http.StatusUnauthorized, getErrorMiddlewareResponse(ErrNotAuth))
		}

		return next(c)
	}
}
