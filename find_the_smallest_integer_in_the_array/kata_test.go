package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/find_the_smallest_integer_in_the_array"
	"testing"
)

var _ = Describe("Sample Test Cases", func() {
	It("should work for sample tests", func() {
		Expect(Expect(SmallestIntegerFinder([]int{34, 15, 88, 2})).To(Equal(2)))
		Expect(Expect(SmallestIntegerFinder([]int{34, -345, -1, 100})).To(Equal(-345)))
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Find the smallest integer in the array")
}
