package middleware

import (
	"github.com/go-park-mail-ru/2022_1_Wave/service"
	"net/http"
)

// Проверить есть ли у клиента валидная сессия (токен сессии в куки).
func Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if service.IsAuthorized(r) {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, `{"status": "FAIL", "error": "unauthorized"}`, http.StatusUnauthorized)
		}
	})
}

func NotAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !service.IsAuthorized(r) {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, `{"status": "FAIL", "error": "available only to unauthorized users"}`, http.StatusBadRequest)
		}
	})
}
