package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/what_century_is_it"
	"testing"
)

func dotest(y, expected string) {
	Expect(WhatCentury(y)).To(Equal(expected), "With year = \"%s\"", y)
}

var _ = Describe("Tests", func() {
	It("Sample tests", func() {
		dotest("2011", "21st")
		dotest("2154", "22nd")
		dotest("2259", "23rd")
		dotest("1234", "13th")
		dotest("1023", "11th")
		dotest("2000", "20th")
		dotest("2700", "27th")
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "What century is it?")
}
