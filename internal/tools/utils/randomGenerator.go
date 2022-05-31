package utils

import (
	"math/rand"
)

func RandomRune() string {
	letters := []rune{'a', 'A'}

	beginPos := rand.Intn(len(letters))
	beginLetter := letters[beginPos]

	return string(beginLetter + rune(rand.Intn('z'-'a'+1)))
}

func RandomWord(maxLen int64) string {
	word := ""
	for i := int64(0); i < maxLen; i++ {
		word += RandomRune()
	}
	return word
}

var runes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}
