package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/multiples_of_3_or_5"
	"testing"
)

var _ = Describe("Multiples of 3 and 5", func() {

	It("should handle basic cases", func() {
		Expect(Multiple3And5(10)).To(Equal(23))
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Multiples of 3 or 5")
}
