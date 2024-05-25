package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/find_the_smallest"
	"testing"
)

func dotest(n int64, exp []int64) {
	var ans = Smallest(n)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Tests Smallest", func() {

	It("should handle basic cases", func() {
		dotest(261235, []int64{126235, 2, 0})
		dotest(209917, []int64{29917, 0, 1})

	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Find the smallest")
}
