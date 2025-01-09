package kata_test

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
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

var _ = Describe("FindLoop (Hidden)", func() {
	It("finds loop in list of random nodes (between 1 and 10 nodes)", func() {
		loopSize := uint(randomUint64Range(1, 10))
		root := CreateRandomChain(uint(randomUint64Range(1, 10)), loopSize)
		Expect(FindLoop(root)).To(Equal(int(loopSize)))
	})

	It("finds loop in list of random nodes (between 10 and 100 nodes)", func() {
		loopSize := uint(randomUint64Range(10, 100))
		root := CreateRandomChain(uint(randomUint64Range(10, 100)), loopSize)
		Expect(FindLoop(root)).To(Equal(int(loopSize)))
	})

	It("finds loop in list of random nodes (between 100 and 1000 nodes)", func() {
		loopSize := uint(randomUint64Range(100, 1000))
		root := CreateRandomChain(uint(randomUint64Range(100, 1000)), loopSize)
		Expect(FindLoop(root)).To(Equal(int(loopSize)))
	})

	It("finds loop in list of random nodes (between 1000 and 10000 nodes)", func() {
		loopSize := uint(randomUint64Range(1000, 10000))
		root := CreateRandomChain(uint(randomUint64Range(1000, 10000)), loopSize)
		Expect(FindLoop(root)).To(Equal(int(loopSize)))
	})

	It("finds loop in list of random nodes (between 10000 and 100000 nodes)", func() {
		loopSize := uint(randomUint64Range(10000, 100000))
		root := CreateRandomChain(uint(randomUint64Range(10000, 100000)), loopSize)
		Expect(FindLoop(root)).To(Equal(int(loopSize)))
	})

	It("finds loop in list of random nodes (between 100000 and 1000000 nodes)", func() {
		loopSize := uint(randomUint64Range(100000, 1000000))
		root := CreateRandomChain(uint(randomUint64Range(100000, 1000000)), loopSize)
		Expect(FindLoop(root)).To(Equal(int(loopSize)))
	})
})

// randomUint64Range generates a random uint64 number in the range [min, max).
func randomUint64Range(min, max uint64) uint64 {
	if min >= max {
		panic("min must be less than max")
	}

	// Compute the range size
	rangeSize := max - min

	// Generate a random number in [0, rangeSize) and add `min`
	randomOffset := randomUint64N(rangeSize)

	return min + randomOffset
}

// randomUint64N generates a random uint64 number in the range [0, n).
func randomUint64N(n uint64) uint64 {
	if n == 0 {
		panic("n must be greater than 0")
	}

	var num uint64
	for {
		// Generate a random uint64
		err := binary.Read(rand.Reader, binary.LittleEndian, &num)
		if err != nil {
			panic(fmt.Errorf("failed to generate random number: %v", err))
		}

		// If the number is in the range [0, n), return it
		if num < n {
			return num
		}

		return num % n
	}
}
