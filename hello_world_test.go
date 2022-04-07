package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Sample", func() {
	It("should test something", func() {
		Expect(Greet("World")).To(Equal("Hello, World!"))
	})
})
