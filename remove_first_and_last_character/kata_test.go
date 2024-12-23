package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/remove_first_and_last_character"
	"testing"
)

var _ = Describe("RemoveChar", func() {
	Describe("Fixed Tests", func() {
		It("eloquent", func() {
			Expect(RemoveChar("eloquent")).To(Equal("loquen"))
		})
		It("country", func() {
			Expect(RemoveChar("country")).To(Equal("ountr"))
		})
		It("person", func() {
			Expect(RemoveChar("person")).To(Equal("erso"))
		})
		It("place", func() {
			Expect(RemoveChar("place")).To(Equal("lac"))
		})
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Remove First and Last Character")
}
