package kata

// Disemvowel https://www.codewars.com/kata/52fba66badcd10859f00097e/train/go
func Disemvowel(comment string) string {
	var result string

	for _, r := range comment {
		if _, ok := vowels[r]; !ok {
			result += string(r)
		}
	}

	return result
}

var vowels = map[rune]struct{}{
	'a': {},
	'A': {},
	'e': {},
	'E': {},
	'i': {},
	'I': {},
	'o': {},
	'O': {},
	'u': {},
	'U': {},
}
