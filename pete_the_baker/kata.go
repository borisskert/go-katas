package kata

import "math"

// Cakes / Pete, the baker
// https://www.codewars.com/kata/525c65e51bf619685c000059/train/go
func Cakes(recipe, available map[string]int) int {
	minimum := math.MaxInt

	for ingredient, amount := range recipe {
		if available[ingredient] == 0 {
			return 0
		}

		possible := available[ingredient] / amount

		if possible < minimum {
			minimum = possible
		}
	}

	return minimum
}
