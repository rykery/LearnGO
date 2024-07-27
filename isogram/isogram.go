package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	s := strings.ToLower(word)
	for i, c := range s {
		if unicode.IsLetter(c) && strings.ContainsRune(s[i+1:], c) {
			return false
		}
	}
	return true
}
