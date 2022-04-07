package helpers

import "golang.org/x/crypto/bcrypt"

func GetPasswordHash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func CheckPassword(passwordHash string, password string) bool {
	res := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return res == nil
}
