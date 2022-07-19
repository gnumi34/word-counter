package utils

import (
	"strings"
	"unicode"
)

func NonLetterRemover(input string) string {
	var b strings.Builder
	b.Grow(len(input))
	for _, char := range input {
		if unicode.IsLetter(char) {
			b.WriteRune(char)
		}
	}
	return strings.ToLower(b.String())
}
