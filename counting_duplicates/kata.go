package kata

import "strings"

// duplicate_count https://www.codewars.com/kata/54bf1c2cd5b56cc47f0007a1/train/go
func duplicate_count(s1 string) int {
	s1 = strings.ToLower(s1)
	letters := map[rune]uint{}

	duplicates := 0

	for _, r := range s1 {
		if letters[r] == 1 {
			duplicates++
		}

		letters[r]++
	}

	return duplicates
}
