// Ginkgo BDD Testing Framework <http://onsi.github.io/ginkgo></http:>
// Gomega Matcher Library <http://onsi.github.io/gomega></http:>

package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/vowel_count"
	"testing"
)

var _ = Describe("Test Example", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(GetCount("abracadabra")).To(Equal(5))
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Vowel Count")
}
