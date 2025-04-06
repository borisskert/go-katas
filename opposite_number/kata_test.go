package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/opposite_number"
	"testing"
)

var _ = Describe("opposite function", func() {
	It("should invert number", func() {
		Expect(Opposite(1)).To(Equal(-1))
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Opposite number")
}
