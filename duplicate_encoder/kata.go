package kata

import (
	"strings"
)

// DuplicateEncode https://www.codewars.com/kata/54b42f9314d9229fd6000d9c/train/go
func DuplicateEncode(word string) string {
	encoder := NewEncoder(strings.ToLower(word))
	return encoder.Encode()
}

type Encoder struct {
	word  string
	runes map[rune]uint
}

func NewEncoder(word string) *Encoder {
	return &Encoder{
		word:  word,
		runes: countRunes(word),
	}
}

func countRunes(word string) map[rune]uint {
	counts := make(map[rune]uint)

	for _, r := range word {
		counts[r]++
	}

	return counts
}

func (e *Encoder) Encode() string {
	var result strings.Builder
	result.Grow(len(e.word))

	for _, r := range e.word {
		if e.runes[r] > 1 {
			result.WriteRune(')')
		} else {
			result.WriteRune('(')
		}
	}

	return result.String()
}
