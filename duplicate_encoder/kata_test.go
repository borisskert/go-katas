package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/duplicate_encoder"
	"testing"
)

var _ = Describe("Test Example", func() {
	It("should return the correct value", func() {
		Expect(DuplicateEncode("din")).To(Equal("((("))
		Expect(DuplicateEncode("recede")).To(Equal("()()()"))
		Expect(DuplicateEncode("(( @")).To(Equal("))(("))
	})

	It("should ignore case", func() {
		Expect(DuplicateEncode("Success")).To(Equal(")())())"))
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Duplicate Encoder")
}
