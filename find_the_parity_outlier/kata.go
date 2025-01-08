package kata

// FindOutlier https://www.codewars.com/kata/5526fc09a1bbd946250002dc/train/go
func FindOutlier(integers []int) int {
	var even, odd *int
	var oddCount, evenCount uint

	for _, i := range integers {
		value := i // See https://github.com/kyoh86/exportloopref

		if i%2 == 0 {
			evenCount++
			even = &value
		} else {
			oddCount++
			odd = &value
		}

		if evenCount > 1 && odd != nil {
			return *odd
		}

		if oddCount > 1 && even != nil {
			return *even
		}
	}

	panic("no outlier found")
}
