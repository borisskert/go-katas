package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/find_the_next_perfect_square"
	"testing"
)

var _ = Describe("Test Example", func() {
	It("fixed tests", func() {
		Expect(FindNextSquare(int64(121))).To(Equal(int64(144)))
		Expect(FindNextSquare(int64(625))).To(Equal(int64(676)))
		Expect(FindNextSquare(int64(319225))).To(Equal(int64(320356)))
		Expect(FindNextSquare(int64(15241383936))).To(Equal(int64(15241630849)))
		Expect(FindNextSquare(int64(155))).To(Equal(int64(-1)))
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Find the next perfect square!")
}
