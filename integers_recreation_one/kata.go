package kata

import "math"

// ListSquared / Integers: Recreation One
// https://www.codewars.com/kata/55aa075506463dac6600010d/train/go
func ListSquared(m, n int) [][]int {
	var result = make([][]int, 0)

	for i := m; i <= n; i++ {
		divisors := divisors(i)
		sum := sumSquares(divisors)

		if isSquare(sum) {
			result = append(result, []int{i, sum})
		}
	}

	return result
}

func divisors(n int) []int {
	var result []int

	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			result = append(result, i)
			divisor := n / i

			if i != divisor {
				result = append(result, divisor)
			}
		}
	}

	return result
}

func sumSquares(nums []int) int {
	var sum int

	for _, num := range nums {
		sum += num * num
	}

	return sum
}

func isSquare(n int) bool {
	nAsFloat := float64(n)
	sqrtAsFloat := math.Sqrt(nAsFloat)

	sqrt := int(sqrtAsFloat)

	return sqrt*sqrt == n
}
