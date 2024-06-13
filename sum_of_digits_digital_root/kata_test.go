// Ginkgo BDD Testing Framework <http://onsi.github.io/ginkgo></http:>
// Gomega Matcher Library <http://onsi.github.io/gomega></http:>

package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/sum_of_digits_digital_root"
	"testing"
)

var _ = Describe("Test Example", func() {
	It("fixed tests", func() {
		Expect(DigitalRoot(16)).To(Equal(7))
		Expect(DigitalRoot(195)).To(Equal(6))
		Expect(DigitalRoot(992)).To(Equal(2))
		Expect(DigitalRoot(167346)).To(Equal(9))
		Expect(DigitalRoot(0)).To(Equal(0))
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sum of Digits / Digital Root")
}
