package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/bit_counting"
	"testing"
)

var _ = Describe("CountBits()", func() {
	It("basic tests", func() {
		Expect(CountBits(0)).To(Equal(0))
		Expect(CountBits(4)).To(Equal(1))
		Expect(CountBits(7)).To(Equal(3))
		Expect(CountBits(9)).To(Equal(2))
		Expect(CountBits(10)).To(Equal(2))
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bit Counting")
}
