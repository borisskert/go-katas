package kata

// JosephusSurvivor https://www.codewars.com/kata/555624b601231dc7a400017a/train/go
func JosephusSurvivor(n, k int) int {
	var x = 1
	var i = 1

	for ; i <= n; i++ {
		x += k - 1
		x = x%i + 1
	}

	return x
}
