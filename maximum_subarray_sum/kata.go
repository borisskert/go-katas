package kata

// MaximumSubarraySum https://www.codewars.com/kata/54521e9ec8e60bc4de000d6c/train/go
func MaximumSubarraySum(numbers []int) int {
	current := 0
	maximum := 0

	for _, number := range numbers {
		current = Max(0, current+number)
		maximum = Max(current, maximum)
	}

	return maximum
}

// Max https://www.pixelstech.net/article/1559993656-Why-no-max-min-function-for-integer-in-GoLang
func Max(a int, b int) int {
	if a >= b {
		return a
	}

	return b
}
