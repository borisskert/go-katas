package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func dotest1(s string, shift int, exp []string) {
	var ans = Encode(s, shift)
	Expect(ans).To(Equal(exp))
}
func dotest2(arr []string, exp string) {
	var ans = Decode(arr)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Tests", func() {

	It("should handle basic cases Encode", func() {
		var u = "I should have known that you would have a perfect answer for me!!!"
		var v = []string{"ijJ tipvme ibw", "f lopxo uibu z", "pv xpvme ibwf ", "b qfsgfdu botx", "fs gps nf!!!"}
		dotest1(u, 1, v)

		u = "abcdefghjuty12"
		v = []string{"abbc", "defg", "hikv", "uz12"}
		dotest1(u, 1, v)

	})
	It("should handle basic cases Decode", func() {
		var u = "I should have known that you would have a perfect answer for me!!!"
		var v = []string{"ijJ tipvme ibw", "f lopxo uibu z", "pv xpvme ibwf ", "b qfsgfdu botx", "fs gps nf!!!"}
		dotest2(v, u)

	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Second Variation on Caesar Cipher")
}
