package kata

// solution https://www.codewars.com/kata/51f2d1cafc9c0f745c00037d/train/go
func solution(str, ending string) bool {
	lenInput := len(str)
	lenEnding := len(ending)

	if lenEnding > lenInput {
		return false
	}

	return str[lenInput-lenEnding:] == ending
}

// #againwhatlearned
// Use `strings.HasSuffix(str, ending)` to check if a string ends with another string.
