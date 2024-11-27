// TODO: replace with your own tests (TDD). An example to get you started is included below.
// Ginkgo BDD Testing Framework <http://onsi.github.io/ginkgo/>
// Gomega Matcher Library <http://onsi.github.io/gomega/>

package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/array_diff"
	"testing"
)

var _ = Describe("Sample tests", func() {
	It("should handle basic cases", func() {
		var emptyArr []int
		Expect(ArrayDiff([]int{1, 2}, []int{1})).To(Equal([]int{2}))
		Expect(ArrayDiff([]int{1, 2, 2}, []int{1})).To(Equal([]int{2, 2}))
		Expect(ArrayDiff([]int{1, 2, 2}, []int{2})).To(Equal([]int{1}))
		Expect(ArrayDiff([]int{1, 2, 2}, emptyArr)).To(Equal([]int{1, 2, 2}))
		Expect(ArrayDiff(emptyArr, []int{1, 2})).To(BeEmpty())
		Expect(ArrayDiff([]int{1, 2, 3}, []int{1, 2})).To(Equal([]int{3}))
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Array.diff")
}
