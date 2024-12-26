package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/remove_string_spaces"
	"testing"
)

var _ = Describe("Tests", func() {
	It("should work for sample tests", func() {
		Expect(NoSpace("8 j 8   mBliB8g  imjB8B8  jl  B")).To(Equal("8j8mBliB8gimjB8B8jlB"))
		Expect(NoSpace("8 8 Bi fk8h B 8 BB8B B B  B888 c hl8 BhB fd")).To(Equal("88Bifk8hB8BB8BBBB888chl8BhBfd"))
		Expect(NoSpace("8aaaaa dddd r     ")).To(Equal("8aaaaaddddr"))
		Expect(NoSpace("jfBm  gk lf8hg  88lbe8 ")).To(Equal("jfBmgklf8hg88lbe8"))
		Expect(NoSpace("8j aam")).To(Equal("8jaam"))
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Remove String Spaces")
}
