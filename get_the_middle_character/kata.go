package kata

// GetMiddle https://www.codewars.com/kata/56747fd5cb988479af000028/train/go
func GetMiddle(s string) string {
	length := len(s)
	half := length / 2

	start := half + length%2 - 1
	end := half + 1

	return s[start:end]
}
