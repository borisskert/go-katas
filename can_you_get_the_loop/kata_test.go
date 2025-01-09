package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/can_you_get_the_loop"
)

var _ = Describe("FindLoop", func() {
	It("finds loop in list of 4 nodes", func() {
		root := CreateChain(map[int]int{
			1: 2,
			2: 3,
			3: 4,
			4: 2,
		})

		Expect(FindLoop(root)).To(Equal(3))
	})

	It("finds loop in list of 2 nodes", func() {
		root := CreateChain(map[int]int{
			1: 2,
			2: 1,
		})

		Expect(FindLoop(root)).To(Equal(2))
	})

	It("finds loop in list of 1 node", func() {
		root := CreateChain(map[int]int{
			1: 1,
		})

		Expect(FindLoop(root)).To(Equal(1))
	})

	It("finds loop in list of 10 node", func() {
		root := CreateChain(map[int]int{
			4:    5,
			5:    2,
			2:    65,
			65:   8,
			8:    112,
			112:  -45,
			-45:  6,
			6:    3,
			3:    7,
			7:    2324,
			2324: 112,
		})

		Expect(FindLoop(root)).To(Equal(6))
	})

	It("finds loop in list of random nodes", func() {
		root := CreateRandomChain(1000, 3000)
		Expect(FindLoop(root)).To(Equal(3000))
	})
})
