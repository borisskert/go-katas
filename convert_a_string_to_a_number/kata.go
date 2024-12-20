package kata

import "strconv"

// StringToNumber https://www.codewars.com/kata/544675c6f971f7399a000e79/train/go
func StringToNumber(str string) int {
	i, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return i
}
