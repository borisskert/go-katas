package kata

// ChooseBestSum ... https://www.codewars.com/kata/55e7280b40e1c4a06d0000aa/train/go
func ChooseBestSum(t, k int, ls []int) int {
	return Go(t, k, 0, 0, ls)
}

func Go(max, count, index, result int, list []int) int {
	if count == 0 && result <= max {
		return result
	}

	if result > max || index >= len(list) {
		return -1
	}

	a := Go(max, count, index+1, result, list)
	b := Go(max, count-1, index+1, result+list[index], list)

	if a < 0 {
		return b
	}

	if b < 0 {
		return a
	}

	if a > b {
		return a
	}

	return b
}
