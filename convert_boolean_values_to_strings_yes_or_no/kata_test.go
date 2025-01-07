package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/convert_boolean_values_to_strings_yes_or_no"
	"testing"
)

var _ = Describe("Test Example", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(BoolToWord(true)).To(Equal("Yes"))
		Expect(BoolToWord(false)).To(Equal("No"))
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Convert boolean values to strings 'Yes' or 'No'.")
}
