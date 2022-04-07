package kata

import "strings"

// SpinWords https://www.codewars.com/kata/5264d2b162488dc400000001/train/go
func SpinWords(str string) string {
	words := strings.Fields(str)

	mapWord := func(word string) string {
		if len(word) >= 5 {
			return Reverse(word)
		}
		return word
	}

	spinWords := Map(words, mapWord)

	return strings.Join(spinWords, " ")
}

func Reverse(str string) string {
	chars := []rune(str)
	size := len(chars)
	reversed := make([]rune, size, size)

	for index := 0; index < size; index++ {
		reversed[index] = chars[size-index-1]
	}

	return string(reversed)
}

func Map(arr []string, fn func(it string) string) []string {
	var newArray []string

	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}

	return newArray
}
