package kata

import "math"

// FindNextSquare https://www.codewars.com/kata/56269eb78ad2e4ced1000013/train/go
func FindNextSquare(sq int64) int64 {
	root, isPerfect := sqrtInt64(sq)
	if !isPerfect {
		return -1
	}

	next := root + 1

	return next * next
}

func sqrtInt64(n int64) (int64, bool) {
	root := int64(math.Sqrt(float64(n)))
	return root, root*root == n
}
