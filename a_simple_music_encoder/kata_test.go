package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"sort"
	"testing"
)

var _ = Describe("Tests", func() {
	It("Sample tests", func() {
		Expect(Compress([]int{1, 2, 2, 3})).To(Equal("1,2*2,3"))
		Expect(Compress([]int{1, 3, 4, 5, 7})).To(Equal("1,3-5,7"))
		Expect(Compress([]int{1, 5, 4, 3, 7})).To(Equal("1,5-3,7"))
		Expect(Compress([]int{1, 10, 8, 6, 7})).To(Equal("1,10-6/2,7"))
		Expect(Compress([]int{4, 157, 157, 157, 50, 51, 52, 53})).To(Equal("4,157*3,50-53"))
		Expect(Compress([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
			11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
			21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
			31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
			41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
		})).To(Equal("1-50"))
	})

	It("should create List with copied items", func() {
		items := []int{5, 2, 6, 3, 4}
		list := ListFrom(items)

		sort.Ints(items) // check if `ListFrom(items)` returns a copy

		Expect(list.Length()).To(Equal(5))
		Expect(list.Get(0)).To(Equal(5))
		Expect(list.Get(1)).To(Equal(2))
		Expect(list.Get(2)).To(Equal(6))
		Expect(list.Get(3)).To(Equal(3))
		Expect(list.Get(4)).To(Equal(4))
	})

	It("should get Init from List", func() {
		items := []int{5, 2, 6, 3, 4}
		list := ListFrom(items)

		init := list.Init()

		sort.Ints(items) // check if `Init()` returns a copy

		Expect(init.Length()).To(Equal(4))
		Expect(init.Get(0)).To(Equal(5))
		Expect(init.Get(1)).To(Equal(2))
		Expect(init.Get(2)).To(Equal(6))
		Expect(init.Get(3)).To(Equal(3))
	})

	It("should get Tail from List", func() {
		items := []int{5, 2, 6, 3, 4}
		list := ListFrom(items)

		tail := list.Tail()

		sort.Ints(items) // check if `Tail()` returns a copy

		Expect(tail.Length()).To(Equal(4))
		Expect(tail.Get(0)).To(Equal(2))
		Expect(tail.Get(1)).To(Equal(6))
		Expect(tail.Get(2)).To(Equal(3))
		Expect(tail.Get(3)).To(Equal(4))
	})

	It("should create List of Lists", func() {
		five := ListOf(5)
		two := ListOf(2)
		six := ListOf(6)
		three := ListOf(3)
		four := ListOf(4)

		items := []List[int]{five, two, six, three, four}

		list := ListFrom(items)

		Expect(list.Length()).To(Equal(5))
		Expect(list.Get(0).Head()).To(Equal(5))
		Expect(list.Get(1).Head()).To(Equal(2))
		Expect(list.Get(2).Head()).To(Equal(6))
		Expect(list.Get(3).Head()).To(Equal(3))
		Expect(list.Get(4).Head()).To(Equal(4))
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "A Simple Music Encoder")
}
