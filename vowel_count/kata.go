package kata

// GetCount https://www.codewars.com/kata/54ff3102c1bad923760001f3/train/go
func GetCount(str string) (count int) {
	for _, r := range str {
		if _, ok := vowels[r]; ok {
			count++
		}
	}

	return
}

var vowels = map[rune]struct{}{
	'a': {},
	'e': {},
	'i': {},
	'o': {},
	'u': {},
}
