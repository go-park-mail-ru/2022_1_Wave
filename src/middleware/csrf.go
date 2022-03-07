package middleware

import (
	"net/http"

	"crypto/sha256"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2022_1_Wave/config"
	"github.com/gorilla/csrf"
)

var csrfMiddleware func(http.Handler) http.Handler

func init() {
	// 256 bits / 8 = 32 bytes
	hash := sha256.New()
	key32Bytes := hash.Sum([]byte(config.C.CSRFKey))
	csrfMiddleware = csrf.Protect(key32Bytes,
		csrf.HttpOnly(true),
		csrf.SameSite(csrf.SameSiteMode(http.SameSiteStrictMode)),
		csrf.Secure(false),
		csrf.ErrorHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(`{"message": "Forbidden - CSRF token invalid"}`))
		})),
	)
}

// Проверить POST запрос на наличие валидного CSRF токена.
func CSRF(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if service.CheckCSRF(r) {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, `{"error": "invalid csrf"}`, http.StatusUnauthorized)
		}
	})
}
