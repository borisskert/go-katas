package kata

import "fmt"

// LastDigit / Last digit of a huge number
// https://www.codewars.com/kata/5518a860a73e708c0a000027/train/go
func LastDigit(as []int) int {
	fmt.Println(as)

	if len(as) == 0 {
		return 1
	}

	n := as[0] % 10

	if len(as) == 1 {
		return n
	}

	m := (as[1]-1)%4 + 1

	if n == 1 || m == 0 {
		return 1
	}

	if n == 0 || n == 5 || n == 6 {
		return n
	}

	//result := int(1)
	//
	//for i := len(as) - 1; i > 0; i-- {
	//	n := as[i]
	//	pow2 := veryFastPow2(n, result)
	//	result = (pow2-1)%4 + 1
	//}
	//
	//return veryFastPow2(as[0]%10, result) % 10

	result := _go(as, len(as)-1, 1)
	return result % 10
}

func _go(as []int, index int, result int) int {
	n := as[index]
	result2 := fastPow(n, result)

	if index == 0 {
		return result2
	}

	m := (result2-1)%4 + 1

	result3 := _go(as, index-1, m)

	return result3
}

func Pow(a, b uint) uint {
	if b == 0 {
		return 1
	}

	if a == 0 {
		return 0
	}

	return fastPow2(a, b)
}

//func Pow(a, b int, mod uint) int {
//	if b == 0 {
//		return 1
//	}
//
//	if a == 0 {
//		return 0
//	}
//
//	pow2 := fastPow2(uint(a), uint(b))
//	return int(pow2) % int(mod)
//}

//func fastPow(base, exp, mod uint) uint {
//	if exp == 0 {
//		return 0
//	}
//
//	if exp == 1 {
//		return base % mod
//	}
//
//	nextBase := (base * base) % mod
//
//	if exp%2 == 0 {
//		return fastPow(nextBase, exp/2, mod)
//	}
//
//	return (base * fastPow(nextBase, exp/2, mod)) % mod
//}

func fastPow(base, exp int) int {
	if exp == 0 {
		return 1
	}

	if exp == 1 {
		return base
	}

	if base == 0 || base == 1 {
		return base
	}

	return fastPowInt(base, exp)
}

func fastPowInt(base, exp int) int {
	if exp == 0 {
		return 0
	}

	if exp == 1 {
		return base
	}

	nextBase := base * base

	if exp%2 == 0 {
		return fastPowInt(nextBase, exp/2)
	}

	return base * fastPowInt(nextBase, exp/2)
}

func fastPow2(base, exp uint) uint {
	if exp == 0 {
		return 0
	}

	if exp == 1 {
		return base
	}

	nextBase := base * base

	if exp%2 == 0 {
		return fastPow2(nextBase, exp/2)
	}

	return base * fastPow2(nextBase, exp/2)
}

func veryFastPow2(n, m int) int {
	if m == 0 {
		return 1
	}

	if m == 1 || n == 0 || n == 1 || n == 5 || n == 6 {
		return n
	}

	rem := uint((m-1)%4) + 1

	if rem == 0 && m != 0 {
		return int(Pow(uint(n), 4))
	}

	return int(Pow(uint(n), rem))
}
