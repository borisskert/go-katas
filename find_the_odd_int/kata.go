package kata

// FindOdd / Find the odd int
// https://www.codewars.com/kata/54da5a58ea159efa38000836/train/go
func FindOdd(seq []int) int {
	xor := 0

	for _, n := range seq {
		xor ^= n
	}

	return xor
}
