package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/disemvowel_trolls"
	"testing"
)

var _ = Describe("Sample Test", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(Disemvowel("This website is for losers LOL!")).To(Equal("Ths wbst s fr lsrs LL!"))
		Expect(Disemvowel("UoIeA")).To(Equal(""))
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Directions Reduction")
}
