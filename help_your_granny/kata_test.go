package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func dotest(arrFriends []string, ftwns map[string]string, h map[string]float64, exp int) {
	var ans = Tour(arrFriends, ftwns, h)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Tests Tour", func() {
	It("should handle basic cases", func() {

		var friends1 = []string{"A1", "A2", "A3", "A4", "A5"}
		var fTowns1 = map[string]string{"A1": "X1", "A2": "X2", "A3": "X3", "A4": "X4"}
		var dist1 = map[string]float64{"X1": 100.0, "X2": 200.0, "X3": 250.0, "X4": 300.0}
		dotest(friends1, fTowns1, dist1, 889)

		friends1 = []string{"B1", "B2"}
		fTowns1 = map[string]string{"B1": "Y1", "B2": "Y2", "B3": "Y3", "B4": "Y4", "B5": "Y5"}
		dist1 = map[string]float64{"Y1": 50.0, "Y2": 70.0, "Y3": 90.0, "Y4": 110.0, "Y5": 150.0}
		dotest(friends1, fTowns1, dist1, 168)

	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Help your granny!")
}
