package middleware

import (
	"net/http"

	"crypto/sha256"
	"github.com/NNKulickov/wave.music_backend/config"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
	adapter "github.com/gwatts/gin-adapter"
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
func CSRF() gin.HandlerFunc {
	return adapter.Wrap(csrfMiddleware)
}
