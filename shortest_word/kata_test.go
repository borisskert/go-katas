package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/shortest_word"
	"testing"
)

var _ = Describe("Test Example", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(FindShort("bitcoin take over the world maybe who knows perhaps")).To(Equal(3))
		Expect(FindShort("turns out random test cases are easier than writing out basic ones")).To(Equal(3))
		Expect(FindShort("Let's travel abroad shall we")).To(Equal(2))
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shortest Word")
}
