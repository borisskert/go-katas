package kata

import (
	"sort"
)

type Word string
type Words []string

// Anagrams https://www.codewars.com/kata/523a86aa4230ebb5420001e1/train/go
func Anagrams(word string, words []string) []string {
	sorted := Word(word).Sorted()

	areAnagrams := func(b Word) bool {
		return sorted == b.Sorted()
	}

	return Words(words).Filter(areAnagrams)
}

func (ctx Word) Sorted() Word {
	sorted := []rune(ctx)

	sort.SliceStable(sorted, func(i, j int) bool {
		a := sorted[i]
		b := sorted[j]

		return a < b
	})

	return Word(sorted)
}

func (ctx Words) Filter(predicate func(x Word) bool) Words {
	var filtered Words

	for _, x := range ctx {
		if predicate(Word(x)) {
			filtered = append(filtered, x)
		}
	}

	return filtered
}

// #againwhatlearned
// You can sort a string using the "strings" import
//
//func sortString(word string) string {
//	w := strings.Split(word, "")
//	sort.Strings(w)
//	return strings.Join(w, "")
//}
