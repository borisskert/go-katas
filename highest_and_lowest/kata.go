package kata

import (
	"strconv"
	"strings"
)

// HighAndLow https://www.codewars.com/kata/554b4ac871d6813a03000035/train/go
func HighAndLow(in string) string {
	values := toInts(words(in))

	highest, lowest := highestAndLowest(values)

	return strconv.Itoa(highest) + " " + strconv.Itoa(lowest)
}

func highestAndLowest(in []int) (int, int) {
	lowest := in[0]
	highest := in[0]

	for _, v := range in {
		if v > highest {
			highest = v
		}

		if v < lowest {
			lowest = v
		}
	}

	return highest, lowest
}

func words(in string) []string {
	return strings.Split(in, " ")
}

func toInts(in []string) []int {
	out := make([]int, len(in))

	for index, v := range in {
		intValue, err := strconv.Atoi(v)

		if err != nil {
			panic(err)
		}

		out[index] = intValue
	}

	return out
}

// #againwhatlearned
// Use strings.Fields(in) instead of strings.Split(in, " ") to split a string into words.
