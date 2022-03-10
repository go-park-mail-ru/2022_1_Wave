package middleware

import (
	"github.com/go-park-mail-ru/2022_1_Wave/service"
	"net/http"
)

// Проверить POST запрос на наличие валидного CSRF токена.
func CSRF(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if service.CheckCSRF(r) {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, `{"status": "FAIL", "error": "invalid csrf"}`, http.StatusUnauthorized)
		}
	})
}
