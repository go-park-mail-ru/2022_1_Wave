package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const secret = "secret_key"

func CreateCSRF(sessionId string, tokenExpTime int64) string {
	h := hmac.New(sha256.New, []byte(secret))
	data := fmt.Sprintf("%s:%d", sessionId, tokenExpTime)
	h.Write([]byte(data))
	token := hex.EncodeToString(h.Sum(nil)) + ":" + strconv.FormatInt(tokenExpTime, 10)
	return token
}

func CheckCSRF(sessionId string, token string) bool {
	tokenData := strings.Split(token, ":")
	if len(tokenData) != 2 {
		return false
	}

	tokenExp, err := strconv.ParseInt(tokenData[1], 10, 64)
	if err != nil {
		return false
	}

	if tokenExp < time.Now().Unix() {
		return false
	}

	h := hmac.New(sha256.New, []byte(secret))
	data := fmt.Sprintf("%s:%d", sessionId, tokenExp)
	h.Write([]byte(data))
	expectedMAC := h.Sum(nil)
	messageMAC, err := hex.DecodeString(tokenData[0])
	if err != nil {
		return false
	}

	return hmac.Equal(messageMAC, expectedMAC)
}
