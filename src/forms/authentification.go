package forms

import (
	"github.com/dgrijalva/jwt-go"
)

type SessionTokenRequest struct {
	UserID   uint   `binding:"required"`
	UserName string `binding:"required"`
}

type SessionClaims struct {
	UserID   uint   `json:"SubjectID"`
	UserName string `json:"UserName"`
	jwt.StandardClaims
}

type Session struct {
	SubjectID uint
	PeriodID  uint
}
