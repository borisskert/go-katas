package kata

// MakeNegative https://www.codewars.com/kata/55685cd7ad70877c23000102/train/go
func MakeNegative(x int) int {
	if x > 0 {
		return -x
	}

	return x
}
