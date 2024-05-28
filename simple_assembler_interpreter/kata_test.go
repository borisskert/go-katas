package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/simple_assembler_interpreter"
	"testing"
)

var _ = Describe("Basic Tests", func() {
	It("should work for some basic example programs", func() {
		Expect(SimpleAssembler([]string{"mov a 5"})).To(Equal(map[string]int{"a": 5}))
		Expect(SimpleAssembler([]string{"mov a 5", "inc a", "dec a", "dec a", "jnz a -1", "inc a"})).To(Equal(map[string]int{"a": 1}))
		Expect(SimpleAssembler([]string{"mov a -10", "mov b a", "inc a", "dec b", "jnz a -2"})).To(Equal(map[string]int{"a": 0, "b": -20}))
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Simple assembler interpreter")
}
