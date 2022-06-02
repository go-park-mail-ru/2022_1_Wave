package utils

import (
	"github.com/google/uuid"
)

func CreateCSRF() string {
	return uuid.NewString()
}

func CheckCSRF(cookieCsrf string, headerCsrf string) bool {
	return cookieCsrf == headerCsrf
}
