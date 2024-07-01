package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/bananas"
	"sort"
	"testing"
)

func arrFromMap(m map[string]bool) []string {
	arr := []string{}
	for k := range m {
		arr = append(arr, k)
	}
	sort.Strings(arr)
	return arr
}

func dotest(s string, expected map[string]bool) {
	userAns := Bananas(s)
	if len(expected) == 0 {
		Expect(userAns).To(BeEmpty())
	} else {
		actualMap := make(map[string]bool)
		for _, a := range userAns {
			actualMap[a] = true
		}
		actual := arrFromMap(actualMap)
		exp := arrFromMap(expected)
		Expect(actual).To(Equal(exp))
	}
}

var _ = Describe("Tests", func() {
	It("Example 0", func() {
		dotest("", map[string]bool{})
	})

	It("Example 1", func() {
		dotest("banana", map[string]bool{"banana": true})
	})

	It("Example 2", func() {
		dotest("bbananana", map[string]bool{
			"-b--anana": true,
			"b-anana--": true,
			"-banan--a": true,
			"-banana--": true,
			"b---anana": true,
			"b-a--nana": true,
			"-ba--nana": true,
			"-ban--ana": true,
			"b-anan--a": true,
			"b-an--ana": true,
			"b-ana--na": true,
			"-bana--na": true,
		})
	})

	It("Example 3", func() {
		dotest("bananaaa", map[string]bool{
			"banan--a": true,
			"banana--": true,
			"banan-a-": true,
		})
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bananas")
}
