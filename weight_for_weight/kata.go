package kata

import (
	"sort"
	"strings"
)

// OrderWeight https://www.codewars.com/kata/55c6126177c9441a570000cc/train/go
func OrderWeight(strng string) string {
	return Convert(strng).
		Sorted().
		ToString()
}

type Weight struct {
	crossSum int
	raw      string
}

type Weights struct {
	items []Weight
}

func Convert(strng string) Weights {
	words := strings.Fields(strng)

	var weights []Weight

	for _, raw := range words {
		weight := NewWeight(raw)
		weights = append(weights, weight)
	}

	return Weights{weights}
}

// Sorted sorts the weights by creating a new instance
func (ctx Weights) Sorted() Weights {
	sorted := make([]Weight, len(ctx.items))
	copy(sorted, ctx.items)

	sort.SliceStable(sorted, func(i, j int) bool {
		a := sorted[i]
		b := sorted[j]

		if a.crossSum == b.crossSum {
			return a.raw < b.raw
		}

		return a.crossSum < b.crossSum
	})

	return Weights{sorted}
}

func (ctx Weights) ToString() string {
	var newArray []string

	for _, it := range ctx.items {
		newArray = append(newArray, it.raw)
	}

	return strings.Join(newArray, " ")
}

func NewWeight(raw string) Weight {
	return Weight{crossSumOf(raw), raw}
}

func crossSumOf(raw string) int {
	var crossSum = 0

	for _, element := range raw {
		crossSum += int(element - '0')
	}

	return crossSum
}
