package kata

import (
	"strings"
	"unicode"
)

// Accum https://www.codewars.com/kata/5667e8f4e3f572a8f2000039/train/go
func Accum(s string) string {
	words := makeWords(s)
	return strings.Join(words, "-")
}

func makeWords(s string) []string {
	words := make([]string, len(s))

	for i, r := range s {
		words[i] = makeWord(i+1, unicode.ToLower(r))
	}

	return words
}

func makeWord(count int, r rune) string {
	word := strings.Repeat(string(r), count)
	return capitalize(word)
}

func capitalize(s string) string {
	return strings.ToUpper(s[0:1]) + s[1:]
}
