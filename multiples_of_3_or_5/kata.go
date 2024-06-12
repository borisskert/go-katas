package kata

// Multiple3And5 / Multiples of 3 or 5
// https://www.codewars.com/kata/514b92a657cdc65150000006/train/go
func Multiple3And5(number int) int {
	sum := 0

	for i := 0; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	return sum
}

// Best practice should be:
//func Multiple3And5(number int) int {
//	number--
//	sum_3 := (number / 3) * (number/3 + 1) * 3 / 2
//	sum_5 := (number / 5) * (number/5 + 1) * 5 / 2
//	sum_15 := (number / 15) * (number/15 + 1) * 15 / 2
//	return sum_3 + sum_5 - sum_15
//}
