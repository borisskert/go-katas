package kata

import (
	"fmt"
	"unicode"
)

// FirstNonRepeating / First non-repeating character
// https://www.codewars.com/kata/52bc74d4ac05d0945d00054e/train/go
func FirstNonRepeating(str string) string {
	letters := NewRuneMap(str)

	for _, r := range str {
		if letters.IsUnique(r) {
			return fmt.Sprintf("%c", r)
		}
	}

	return ""
}

type RuneMap map[rune]int

func NewRuneMap(str string) RuneMap {
	rm := make(RuneMap)

	for _, r := range str {
		rm.add(r)
	}

	return rm
}

func (rm RuneMap) add(r rune) {
	key := unicode.ToLower(r)
	rm[key]++
}

func (rm RuneMap) IsUnique(r rune) bool {
	key := unicode.ToLower(r)
	return rm[key] == 1
}
