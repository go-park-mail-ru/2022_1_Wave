package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2022_1_Wave/config"
	"github.com/go-park-mail-ru/2022_1_Wave/forms"
	"time"
)

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

func GetSession(c *gin.Context) (*forms.Session, error) {

	return nil, nil
}

func SetSessionSession(subjectSession *forms.Session, c *gin.Context) error {

	return nil
}
