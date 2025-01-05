package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/get_the_middle_character"
	"testing"
)

var _ = Describe("GetMiddle", func() {
	It("Tests", func() {
		Expect(GetMiddle("test")).To(Equal("es"))
		Expect(GetMiddle("testing")).To(Equal("t"))
		Expect(GetMiddle("middle")).To(Equal("dd"))
		Expect(GetMiddle("A")).To(Equal("A"))
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Get the Middle Character")
}
