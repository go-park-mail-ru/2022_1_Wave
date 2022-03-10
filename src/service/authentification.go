package service

import (
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2022_1_Wave/config"
	"github.com/google/uuid"
	"net/http"
	"sync"
	"time"
)

type Session struct {
	UserId       uint
	CSRFToken    string
	IsAuthorized bool
	Expire       time.Time
}

type userSessions map[string]Session

var Sessions userSessions = make(userSessions)

var sessionMutex sync.RWMutex

// возвращает id пользователя по токену сессии
func GetSession(r *http.Request) (*Session, error) {
	session, err := r.Cookie(config.C.SessionIDKey)
	if err == nil && session != nil {
		sessionMutex.RLock()
		sessionAuth, err := Sessions[session.Value]
		sessionMutex.RUnlock()
		if err {
			if sessionAuth.Expire.Sub(time.Now()) > 0 {
				return &sessionAuth, nil
			} else {
				DeleteSession(session.Value)
				return &sessionAuth, errors.New("session expired")
			}
		}
	}

	return nil, errors.New("auth error")
}

func SetNewUnauthorizedSession() (*http.Cookie, string) {
	sessionId := uuid.NewString()
	csrfToken := uuid.NewString()
	expireTime := time.Now().Add(10 * time.Hour)

	sessionMutex.Lock()
	Sessions[sessionId] = Session{
		CSRFToken:    csrfToken,
		IsAuthorized: false,
		Expire:       expireTime,
	}
	sessionMutex.Unlock()

	cookie := &http.Cookie{
		Name:     config.C.SessionIDKey,
		Value:    sessionId,
		Expires:  expireTime,
		HttpOnly: true,
	}

	fmt.Println(config.C.Domain)

	return cookie, csrfToken
}

func AuthorizeUser(sessionId string, userId uint) {
	sessionMutex.Lock()
	defer sessionMutex.Unlock()
	nowSession, ok := Sessions[sessionId]
	if !ok {
		return
	}
	nowSession.UserId = userId
	nowSession.IsAuthorized = true
	Sessions[sessionId] = nowSession
}

// добавит запись о сессии и вернет сгенерированный идентификатор сессии и csrf-токен
func SetNewSession(userId uint) (*http.Cookie, string) {
	sessionId := uuid.NewString()
	csrfToken := uuid.NewString()
	expireTime := time.Now().Add(10 * time.Hour)

	sessionMutex.Lock()
	Sessions[sessionId] = Session{
		UserId:       userId,
		CSRFToken:    csrfToken,
		IsAuthorized: true,
		Expire:       expireTime,
	}
	sessionMutex.Unlock()

	cookie := &http.Cookie{
		Name:     config.C.SessionIDKey,
		Value:    sessionId,
		Expires:  expireTime,
		HttpOnly: true,
	}

	return cookie, csrfToken
}

func SetNewCSRFToken(sessionId string) string {
	csrfToken := uuid.NewString()

	sessionMutex.Lock()
	session, _ := Sessions[sessionId]
	session.CSRFToken = csrfToken
	Sessions[sessionId] = session
	sessionMutex.Unlock()

	return csrfToken
}

func DeleteSession(sessionId string) error {
	sessionMutex.Lock()
	defer sessionMutex.Unlock()
	if _, ok := Sessions[sessionId]; !ok {
		return errors.New("session does not exist")
	}

	delete(Sessions, sessionId)

	return nil
}

func IsAuthorized(r *http.Request) bool {
	authorized := false
	session, err := r.Cookie(config.C.SessionIDKey)
	if err == nil && session != nil {
		sessionMutex.RLock()
		val, ok := Sessions[session.Value]
		sessionMutex.RUnlock()
		if ok {
			authorized = val.IsAuthorized
		}
	}

	return authorized
}

func CheckCSRF(r *http.Request) bool {
	sessionID, err := r.Cookie(config.C.SessionIDKey)
	if err != nil {
		return false
	}

	sessionMutex.RLock()
	session, ok := Sessions[sessionID.Value]
	sessionMutex.RUnlock()
	if !ok {
		return false
	}

	csrfToken := r.Header.Get("X-CSRF-TOKEN")
	return csrfToken != "" && session.CSRFToken == csrfToken
}
