package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/integers_recreation_one"
	"testing"
)

func dotest(m, n int, exp [][]int) {
	var ans = ListSquared(m, n)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		dotest(1, 250, [][]int{{1, 1}, {42, 2500}, {246, 84100}})
		dotest(250, 500, [][]int{{287, 84100}})
		dotest(300, 600, [][]int{})
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integers: Recreation One")
}
