package kata

// PositiveSum https://www.codewars.com/kata/5715eaedb436cf5606000381/train/go
func PositiveSum(numbers []int) int {
	sum := 0

	for _, n := range numbers {
		if n > 0 {
			sum += n
		}
	}

	return sum
}
