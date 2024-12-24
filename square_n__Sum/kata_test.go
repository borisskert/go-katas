package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/square_n__Sum"
	"testing"
)

var _ = Describe("Sample Tests", func() {
	It("Testing [1,2]", func() { Expect(SquareSum([]int{1, 2})).To(Equal(5)) })
	It("Testing [0,3,4,5]", func() { Expect(SquareSum([]int{0, 3, 4, 5})).To(Equal(50)) })
	It("Testing []", func() { Expect(SquareSum([]int{})).To(Equal(0)) })
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Square(n) Sum")
}
