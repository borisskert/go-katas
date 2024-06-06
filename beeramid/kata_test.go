// TODO: replace with your own tests (TDD). An example to get you started is included below.
// Ginkgo BDD Testing Framework <http://onsi.github.io/ginkgo/>
// Gomega Matcher Library <http://onsi.github.io/gomega/>

package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/beeramid"
	"testing"
)

var _ = Describe("Tests", func() {
	It("Sample tests", func() {
		Expect(Beeramid(9, 2.0)).To(Equal(1))
		Expect(Beeramid(21, 1.5)).To(Equal(3))
		Expect(Beeramid(-1, 4.0)).To(Equal(0))
		Expect(Beeramid(293, 5.5625)).To(Equal(4))
		Expect(Beeramid(25, 1.9375)).To(Equal(2))
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Beeramid")
}
