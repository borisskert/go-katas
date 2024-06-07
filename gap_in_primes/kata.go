package kata

// Gap / Gap in Primes
// https://www.codewars.com/kata/561e9c843a2ef5a40c0000a4/train/go
func Gap(g, m, n int) []int {
	a := NextPrime(m)
	b := NextPrime(a + 2)

	for b <= n {
		if b-a == g {
			return []int{a, b}
		}

		a = b
		b = NextPrime(a + 2)
	}

	return nil
}

func IsPrime(i int) bool {
	if i <= 1 {
		return false
	}

	if i <= 3 {
		return true
	}

	if i%2 == 0 || i%3 == 0 {
		return false
	}

	for j := 5; j*j <= i; j += 6 {
		if i%j == 0 || i%(j+2) == 0 {
			return false
		}
	}

	return true
}

func NextPrime(i int) int {
	if i%2 == 0 {
		i++
	}

	for {
		if IsPrime(i) {
			return i
		}

		i += 2
	}
}
