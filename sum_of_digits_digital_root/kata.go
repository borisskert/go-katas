package kata

// DigitalRoot / Sum of Digits / Digital Root
// https://www.codewars.com/kata/541c8630095125aba6000c00/train/go
func DigitalRoot(n int) int {
	if n < 10 {
		return n
	}

	digits := toDigits(n)
	sum := sum(digits)

	return DigitalRoot(sum)
}

func toDigits(n int) []int {
	var digits []int

	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	return digits
}

func sum(digits []int) int {
	sum := 0

	for _, digit := range digits {
		sum += digit
	}

	return sum
}

// Best practice should be:
// func DigitalRoot(n int) int {
//	 return (n - 1) % 9 + 1
// }
