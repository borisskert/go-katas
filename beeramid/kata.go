package kata

// Beeramid / https://www.codewars.com/kata/51e04f6b544cf3f6550000c1/train/go
func Beeramid(bonus int, price float64) int {
	if float64(bonus) < price {
		return 0
	}

	return beeramid(float64(bonus), price, 0)
}

func beeramid(bonus float64, price float64, level int) int {
	if bonus < 0 {
		return level - 1
	}

	bonus -= float64(level+1) * float64(level+1) * price

	return beeramid(bonus, price, level+1)
}
