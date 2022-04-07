package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

var _ = Describe("Test Sample", func() {
	It("should test that the solution returns the correct value for single word inputs", func() {
		Expect(SpinWords("Welcome")).To(Equal("emocleW"))
		Expect(SpinWords("to")).To(Equal("to"))
		Expect(SpinWords("CodeWars")).To(Equal("sraWedoC"))
	})
	It("should test that the solution returns the correct value for multiple word outputs", func() {
		Expect(SpinWords("Hey fellow warriors")).To(Equal("Hey wollef sroirraw"))
		Expect(SpinWords("Burgers are my favorite fruit")).To(Equal("sregruB are my etirovaf tiurf"))
		Expect(SpinWords("Pizza is the best vegetable")).To(Equal("azziP is the best elbategev"))
	})
})

func TestAdder(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Run SpinWords tests")
}
