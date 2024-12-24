package kata

// SquareSum https://www.codewars.com/kata/515e271a311df0350d00000f/train/go
func SquareSum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number * number
	}

	return sum
}
