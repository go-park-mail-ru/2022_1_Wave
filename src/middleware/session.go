package middleware

import (
	"github.com/NNKulickov/wave.music_backend/config"
	"github.com/NNKulickov/wave.music_backend/service"
	"net/http"
	"time"
)

// Проверить есть ли у клиента валидная сессия (токен сессии в куки).
func Session(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorized := false
		session, err := r.Cookie(config.C.SessionIDKey)
		if err == nil && session.Expires.Sub(time.Now()) <= 0 {
			service.DeleteSession(session.Value)
		}

		if err == nil && session != nil {
			_, authorized = service.Sessions[session.Value]
		}

		if authorized {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, `{"error": "unauthorized"}`, 401)
		}
	})
}
