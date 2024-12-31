package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

var _ = Describe("Example test:", func() {
	It("Should work for fixed tests", func() {
		Expect(solution("", "")).To(Equal(true))
		Expect(solution(" ", "")).To(Equal(true))
		Expect(solution("abc", "c")).To(Equal(true))
		Expect(solution("banana", "ana")).To(Equal(true))
		Expect(solution("a", "z")).To(Equal(false))
		Expect(solution("", "t")).To(Equal(false))
		Expect(solution("$a = $b + 1", "+1")).To(Equal(false))
		Expect(solution("    ", "   ")).To(Equal(true))
		Expect(solution("1oo", "100")).To(Equal(false))
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "String ends with?")
}
