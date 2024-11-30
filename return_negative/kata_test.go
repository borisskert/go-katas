package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/return_negative"
	"testing"
)

var _ = Describe("Test Example", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(MakeNegative(42)).To(Equal(-42))
		Expect(MakeNegative(0)).To(Equal(0))
		Expect(MakeNegative(-9)).To(Equal(-9))
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Return Negative")
}
