package kata

import (
	. "fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

var _ = Describe("Sample Tests", func() {
	It("should handle basic tests", func() {
		dotest(7, 3, 4)
		dotest(11, 19, 10)
		dotest(40, 3, 28)
		dotest(14, 2, 13)
		dotest(100, 1, 100)
	})
})

func dotest(n, k, e int) {
	Println(n, k)
	Expect(JosephusSurvivor(n, k)).To(Equal(e))
}

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Josephus Survivor")
}
