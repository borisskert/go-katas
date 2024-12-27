package kata

// CountSheeps https://www.codewars.com/kata/54edbc7200b811e956000556/train/go
func CountSheeps(numbers []bool) int {
	count := 0

	for _, isSheep := range numbers {
		if isSheep {
			count++
		}
	}

	return count
}
