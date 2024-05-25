package kata

// Smallest https://www.codewars.com/kata/573992c724fc289553000e95/train/go
func Smallest(n int64) []int64 {
	possibleResults := PossibleResults(n)
	minimum := Min(possibleResults)

	return []int64{minimum.number, int64(minimum.pair.i), int64(minimum.pair.j)}
}

type Result struct {
	number int64
	pair   IntPair
}

func PossibleResults(n int64) []Result {
	var results = make([]Result, 0)

	size := Length(n)
	pairs := CartesianProduct(size, size)

	for _, pair := range pairs {
		number := pair.MoveDigit(n)
		results = append(results, Result{number, pair})
	}

	return results
}

type IntPair struct {
	i uint8
	j uint8
}

func CartesianProduct(a uint8, b uint8) []IntPair {
	var pairs []IntPair

	for i := uint8(0); i < a; i++ {
		for j := uint8(0); j < b; j++ {
			pairs = append(pairs, IntPair{i, j})
		}
	}

	return pairs
}

func (ctx IntPair) MoveDigit(n int64) int64 {
	if ctx.i == ctx.j {
		return n
	}

	digits := ToDigits(n)
	length := len(digits)

	from := uint8(length-1) - ctx.i
	to := uint8(length-1) - ctx.j

	digit := digits[from]

	digits = append(digits[:from], digits[from+1:]...)
	digits = append(digits[:to], append([]int64{digit}, digits[to:]...)...)

	return FromDigits(digits)
}

func ToDigits(n int64) []int64 {
	var digits []int64

	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	return digits
}

func FromDigits(digits []int64) int64 {
	var n int64

	for i := len(digits) - 1; i >= 0; i-- {
		n = n*10 + digits[i]
	}

	return n
}

func Length(n int64) uint8 {
	return uint8(len(ToDigits(n)))
}

func Min(values []Result) Result {
	found := values[0]

	for _, current := range values {
		if current.number < found.number {
			found = current
		}
	}

	return found
}
