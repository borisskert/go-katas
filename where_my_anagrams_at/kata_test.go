package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func dotest(word string, words, exp []string) {
	var ans = Anagrams(word, words)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {
	It("should handle basic cases", func() {
		dotest("abba", []string{"aabb", "abcd", "bbaa", "dada"}, []string{"aabb", "bbaa"})
		dotest("laser", []string{"lazing", "lazy", "lacer"}, nil)
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Where my anagrams at?")
}
