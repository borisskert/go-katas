package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/help_the_bookseller"
	"testing"
)

func dotest(listArt []string, listCat []string, exp string) {
	var ans = StockList(listArt, listCat)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Tests StockList", func() {

	It("should handle basic cases", func() {
		var b = []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"}
		var c = []string{"A", "B", "C", "D"}
		dotest(b, c, "(A : 0) - (B : 1290) - (C : 515) - (D : 600)")

		b = []string{"ABAR 200", "CDXE 500", "BKWR 250", "BTSQ 890", "DRTY 600"}
		c = []string{"A", "B"}
		dotest(b, c, "(A : 200) - (B : 1140)")

	})

	It("should handle empty list", func() {
		var b = []string{}
		var c = []string{"B", "R", "D", "X"}
		dotest(b, c, "")
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Help the bookseller !")
}
