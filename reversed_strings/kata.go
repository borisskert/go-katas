package kata

import "strings"

// Solution https://www.codewars.com/kata/5168bb5dfe9a00b126000018/train/go
func Solution(word string) string {
	var sb strings.Builder

	for i := len(word) - 1; i >= 0; i-- {
		sb.WriteByte(word[i])
	}

	return sb.String()
}
