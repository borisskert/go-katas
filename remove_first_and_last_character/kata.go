package kata

// RemoveChar https://www.codewars.com/kata/56bc28ad5bdaeb48760009b0/train/go
func RemoveChar(word string) string {
	return word[1 : len(word)-1]
}
