package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/greed_is_good"
	"testing"
)

var _ = Describe("Scorer Function", func() {
	It("should value this as worthless", func() {
		Expect(Score([5]int{2, 3, 4, 6, 2})).To(Equal(0), "Incorrect answer for dice=[2, 3, 4, 6, 2]")
	})

	It("should value this triplet correctly", func() {
		Expect(Score([5]int{4, 4, 4, 3, 3})).To(Equal(400), "Incorrect answer for dice=[4, 4, 4, 3, 3]")
	})

	It("should value this mixed set correctly", func() {
		Expect(Score([5]int{2, 4, 4, 5, 4})).To(Equal(450), "Incorrect answer for dice=[2, 4, 4, 5, 4]")
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Greed is Good")
}
