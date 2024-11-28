// Ginkgo BDD Testing Framework <http://onsi.github.io/ginkgo></http:>
// Gomega Matcher Library <http://onsi.github.io/gomega></http:>

package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/highest_and_lowest"
	"testing"
)

var _ = Describe("Example Tests", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4")).To(Equal("42 -9"))
	})
	It("test 2", func() {
		Expect(HighAndLow("1 2 3")).To(Equal("3 1"))
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Highest and Lowest")
}
