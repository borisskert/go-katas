package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func dotest(arr []string, exp []string) {
	var ans = DirReduc(arr)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Tests DirReduc", func() {

	It("should handle basic cases", func() {
		var a = []string{"NORTH", "SOUTH", "EAST", "WEST"}
		dotest(a, []string{})
		a = []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"}
		dotest(a, []string{"WEST"})
		a = []string{"NORTH", "WEST", "SOUTH", "EAST"}
		dotest(a, []string{"NORTH", "WEST", "SOUTH", "EAST"})
		a = []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "NORTH"}
		dotest(a, []string{"NORTH"})

	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Directions Reduction")
}
