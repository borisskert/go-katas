package kata

// RepeatStr https://www.codewars.com/kata/57a0e5c372292dd76d000d7e/train/go
func RepeatStr(repetitions int, value string) string {
	var result string

	for counter := 0; counter < repetitions; counter++ {
		result += value
	}

	return result
}

// #againwhatlearned
// use `strings.Repeat` function of the `strings` package
// ```
// return strings.Repeat(value, repetitions)
// ```
