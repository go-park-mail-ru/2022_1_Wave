package utils

import (
	"math/rand"
)

func RandomRune() string {
	return string('a' + rune(rand.Intn('z'-'a'+1)))
}

func RandomWord(maxLen int) string {
	word := ""
	for i := 0; i < maxLen; i++ {
		word += RandomRune()
	}
	return word
}
