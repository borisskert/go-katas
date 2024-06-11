package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/last_digit_of_a_huge_number"
	"testing"
)

var _ = Describe("Test Example", func() {
	It("should handle basic cases", func() {
		Expect(LastDigit([]int{})).To(Equal(1))
		Expect(LastDigit([]int{0, 0})).To(Equal(1))    // 0 ^ 0
		Expect(LastDigit([]int{0, 0, 0})).To(Equal(0)) // 0^(0 ^ 0) = 0^1 = 0
		Expect(LastDigit([]int{1, 2})).To(Equal(1))
		Expect(LastDigit([]int{3, 4, 5})).To(Equal(1))
		Expect(LastDigit([]int{4, 3, 6})).To(Equal(4))
		Expect(LastDigit([]int{6, 21})).To(Equal(6))
		Expect(LastDigit([]int{7, 6})).To(Equal(9))
		Expect(LastDigit([]int{7, 6, 21})).To(Equal(1))
		Expect(LastDigit([]int{7, 6, 21})).To(Equal(9))
		Expect(LastDigit([]int{12, 30, 21})).To(Equal(6))
		Expect(LastDigit([]int{12, 30, 21})).To(Equal(4))
		Expect(LastDigit([]int{2, 0, 1})).To(Equal(1))
		Expect(LastDigit([]int{2, 2, 2, 0})).To(Equal(4))
		Expect(LastDigit([]int{937640, 767456, 981242})).To(Equal(0))
		Expect(LastDigit([]int{123232, 694022, 140249})).To(Equal(6))
		Expect(LastDigit([]int{499942, 898102, 846073})).To(Equal(6))
	})

	//It("should handle random cases", func() {
	//	source := rand.NewSource(time.Now().UnixNano())
	//	rand := rand.New(source)
	//
	//	var r1 int = rand.Intn(100)
	//	var r2 int = rand.Intn(10)
	//	var pow int = int(math.Pow(float64(r1%10), float64(r2)))
	//
	//	Expect(LastDigit([]int{})).To(Equal(1))
	//	Expect(LastDigit([]int{r1})).To(Equal(r1 % 10))
	//	Expect(LastDigit([]int{r1, r2})).To(Equal(pow % 10))
	//})
	//
	//It("should fast-pow numbers", func() {
	//	Expect(Pow(0, 0, 10)).To(Equal(1))
	//	Expect(Pow(1, 0, 10)).To(Equal(1))
	//	Expect(Pow(0, 1, 10)).To(Equal(0))
	//	Expect(Pow(1, 1, 10)).To(Equal(1))
	//	Expect(Pow(2, 1, 10)).To(Equal(2))
	//	Expect(Pow(2, 2, 10)).To(Equal(4))
	//	Expect(Pow(2, 3, 10)).To(Equal(8))
	//	Expect(Pow(2, 4, 10)).To(Equal(6))
	//	Expect(Pow(2, 5, 10)).To(Equal(2))
	//	Expect(Pow(2, 6, 10)).To(Equal(4))
	//	Expect(Pow(2, 7, 10)).To(Equal(8))
	//	Expect(Pow(2, 8, 10)).To(Equal(6))
	//	Expect(Pow(2, 9, 10)).To(Equal(2))
	//	Expect(Pow(2, 10, 10)).To(Equal(4))
	//	Expect(Pow(2, 11, 10)).To(Equal(8))
	//	Expect(Pow(2, 12, 10)).To(Equal(6))
	//	Expect(Pow(2, 13, 10)).To(Equal(2))
	//	Expect(Pow(2, 14, 10)).To(Equal(4))
	//	Expect(Pow(2, 15, 10)).To(Equal(8))
	//	Expect(Pow(2, 16, 10)).To(Equal(6))
	//	Expect(Pow(2, 17, 10)).To(Equal(2))
	//	Expect(Pow(2, 18, 10)).To(Equal(4))
	//	Expect(Pow(2, 19, 10)).To(Equal(8))
	//	Expect(Pow(2, 20, 10)).To(Equal(6))
	//	Expect(Pow(2, 21, 10)).To(Equal(2))
	//	Expect(Pow(2, 22, 10)).To(Equal(4))
	//	Expect(Pow(2, 23, 10)).To(Equal(8))
	//	Expect(Pow(2, 24, 10)).To(Equal(6))
	//	Expect(Pow(2, 25, 10)).To(Equal(2))
	//})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Last digit of a huge number")
}
