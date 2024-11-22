package kata

// EvenOrOdd https://www.codewars.com/kata/53da3dbb4a5168369a0000fe/train/go
func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	}

	return "Odd"
}
