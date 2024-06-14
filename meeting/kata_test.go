package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/meeting"
	"testing"
)

func dotest(s string, exp string) {
	var ans = Meeting(s)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Tests Meeting", func() {

	It("should handle example case", func() {
		dotest("Fred:Corwill;Wilfred:Corwill;Barney:Tornbull;Betty:Tornbull;Bjon:Tornbull;Raphael:Corwill;Alfred:Corwill",
			"(CORWILL, ALFRED)(CORWILL, FRED)(CORWILL, RAPHAEL)(CORWILL, WILFRED)(TORNBULL, BARNEY)(TORNBULL, BETTY)(TORNBULL, BJON)")
	})

	It("should handle basic case 1", func() {
		dotest("Alexis:Wahl;John:Bell;Victoria:Schwarz;Abba:Dorny;Grace:Meta;Ann:Arno;Madison:STAN;Alex:Cornwell;Lewis:Kern;Megan:Stan;Alex:Korn",
			"(ARNO, ANN)(BELL, JOHN)(CORNWELL, ALEX)(DORNY, ABBA)(KERN, LEWIS)(KORN, ALEX)(META, GRACE)(SCHWARZ, VICTORIA)(STAN, MADISON)(STAN, MEGAN)(WAHL, ALEXIS)")
	})

	It("should handle basic case 2", func() {
		dotest("John:Gates;Michael:Wahl;Megan:Bell;Paul:Dorries;James:Dorny;Lewis:Steve;Alex:Meta;Elizabeth:Russel;Anna:Korn;Ann:Kern;Amber:Cornwell",
			"(BELL, MEGAN)(CORNWELL, AMBER)(DORNY, JAMES)(DORRIES, PAUL)(GATES, JOHN)(KERN, ANN)(KORN, ANNA)(META, ALEX)(RUSSEL, ELIZABETH)(STEVE, LEWIS)(WAHL, MICHAEL)")
	})

	It("Schwarz before Stan", func() {
		dotest("Stan:Schwarz;Stan:Stan", "(SCHWARZ, STAN)(STAN, STAN)")
		dotest("Stan:Stan;Stan:Schwarz", "(SCHWARZ, STAN)(STAN, STAN)")
		dotest("Schwarz:Stan;Stan:Stan", "(STAN, SCHWARZ)(STAN, STAN)")
		dotest("Stan:Stan;Schwarz:Stan", "(STAN, SCHWARZ)(STAN, STAN)")
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Meeting")
}
