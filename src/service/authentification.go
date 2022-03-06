package service

import (
	"errors"
	"github.com/NNKulickov/wave.music_backend/config"
	"math/rand"
	"net/http"
	"time"
)

type userSessions map[string]uint

var Sessions userSessions
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// возвращает id пользователя по токену сессии
func GetSession(r *http.Request) (uint, error) {
	session, err := r.Cookie(config.C.SessionIDKey)
	if err == nil && session != nil {
		session_id, err := Sessions[session.Value]
		if err {
			if session.Expires.Sub(time.Now()) > 0 {
				return session_id, nil
			} else {
				return session_id, errors.New("")
			}
		}
	}

	return 0, errors.New("auth error")
}

// добавит запись о сессии и вернет сгенерированный идентификатор сессии
func SetNewSession(userId uint) string {
	sessionId := randStringRunes(32)
	Sessions[sessionId] = userId

	return sessionId
}

func DeleteSession(sessionId string) error {
	if _, ok := Sessions[sessionId]; !ok {
		return errors.New("session does not exist")
	}

	delete(Sessions, sessionId)

	return nil
}

/*

func SetSessionToken(req forms.SessionTokenRequest) (string, error) {
	now := time.Now().Unix()
	claims := forms.SessionClaims{
		UserID:   4,
		UserName: "some",
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now,
			ExpiresAt: now + config.C.TokenMaxAge,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := []byte(config.C.SessionKey)
	ss, err := token.SignedString(key)
	return ss, err
}

func ParseSessionToken(subjectTokenString string) (*forms.SessionClaims, error) {
	token, err := jwt.ParseWithClaims(subjectTokenString, &forms.SessionClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.C.SessionKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*forms.SessionClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("token is invalid")
	}
}

func SetSessionSession(subjectSession *forms.Session, c *gin.Context) error {

	return nil
}
*/
