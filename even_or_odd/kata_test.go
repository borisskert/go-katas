package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/even_or_odd"
	"testing"
)

var _ = Describe("Test Example", func() {

	It("should return \"Odd\" for odd positive numbers", func() {
		Expect(EvenOrOdd(1)).To(Equal("Odd"))
	})

	It("should return \"Even\" for even positive numbers", func() {
		Expect(EvenOrOdd(2)).To(Equal("Even"))
	})

	It("should return \"Odd\" for odd negative numbers", func() {
		Expect(EvenOrOdd(-1)).To(Equal("Odd"))
	})

	It("should return \"Even\" for even negative numbers", func() {
		Expect(EvenOrOdd(-2)).To(Equal("Even"))
	})

	It("should return \"Even\" for zero", func() {
		Expect(EvenOrOdd(0)).To(Equal("Even"))
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Even or Odd")
}
