package utils

import (
	"math/rand"
)

func RandomRune() string {
	return string('a' + rune(rand.Intn('z'-'a'+1)))
}

func RandomWord(maxLen int64) string {
	word := ""
	for i := int64(0); i < maxLen; i++ {
		word += RandomRune()
	}
	return word
}
