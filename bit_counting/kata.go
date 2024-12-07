package kata

// CountBits https://www.codewars.com/kata/526571aae218b8ee490006f4/train/go
func CountBits(n uint) int {
	if n == 0 {
		return 0
	}

	if n%2 == 1 {
		return 1 + CountBits(n/2)
	}

	return CountBits(n / 2)
}

// #againwhatlearned
// Use `bits.OnesCount(n)` to count the number of bits set to 1 in a number.
// This is a built-in function in the `math/bits` package.
