package kata

import "strings"

// FindShort https://www.codewars.com/kata/57cebe1dc6fdc20c57000ac9/train/go
func FindShort(s string) int {
	words := strings.Fields(s)

	return findMinimumOn(words, func(it string) int {
		return len(it)
	})
}

func findMinimumOn(s []string, fn func(it string) int) int {
	minimum := fn(s[0])

	for _, it := range s {
		if fn(it) < minimum {
			minimum = fn(it)
		}
	}

	return minimum
}
