package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/convert_a_number_to_a_string"
	"testing"
)

func dotest(n int, expected string) {
	Expect(NumberToString(n)).To(Equal(expected), "With n = %d", n)
}

var _ = Describe("Tests", func() {
	It("Sample tests", func() {
		dotest(67, "67")
		dotest(79585, "79585")
		dotest(79585, "79585")
		dotest(3, "3")
		dotest(-1, "-1")
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Convert a Number to a String!")
}
