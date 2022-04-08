package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

var _ = Describe("Tests", func() {
	It("Basic tests", func() {
		Describe("Empty arrays should have a max of 0", func() {
			Expect(MaximumSubarraySum([]int{})).To(Equal(0))
		})

		Describe("Example array should have a max of 6", func() {
			Expect(MaximumSubarraySum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})).To(Equal(6))
		})

		Describe("Example array with all negative values should have a max of 0", func() {
			Expect(MaximumSubarraySum([]int{-2, -1, -3, -4, -1, -2, -1, -5, -4})).To(Equal(0))
		})
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Maximum subarray sum")
}
