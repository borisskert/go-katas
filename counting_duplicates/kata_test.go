package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func dotest(prod string, exp int) {
	var ans = duplicate_count(prod)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		dotest("abcde", 0)
		dotest("abcdea", 1)
		dotest("abcdeaB11", 3)
		dotest("indivisibility", 1)
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Counting Duplicates")
}
