package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/convert_a_string_to_a_number"
	"testing"
)

var _ = Describe("Tests", func() {
	It("Sample tests", func() {
		Expect(StringToNumber("1234")).To(Equal(1234))
		Expect(StringToNumber("605")).To(Equal(605))
		Expect(StringToNumber("1405")).To(Equal(1405))
		Expect(StringToNumber("-7")).To(Equal(-7))
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Convert a String to a Number!")
}
