package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

var _ = Describe("CreatePhoneNumber()", func() {
	It("basic test", func() {
		Expect(CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})).To(Equal("(123) 456-7890"))
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Create Phone Number")
}
