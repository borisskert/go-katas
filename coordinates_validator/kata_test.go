package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/coordinates_validator"
	"testing"
)

//var _ = Describe("Valid coordinates", func() {
//
//	validCoordinates := []string{
//		"-23, 25",
//		"4, -3",
//		"24.53525235, 23.45235",
//		"04, -23.234235",
//		"43.91343345, 143"}
//
//	It("should validate the coordinates", func() {
//		for _, coordinates := range validCoordinates {
//			Expect(IsValidCoordinates(coordinates)).To(Equal(true))
//		}
//	})
//})

var _ = Describe("Invalid coordinates", func() {

	invalidCoordinates := []string{
		//"23.234, - 23.4234",
		//"2342.43536, 34.324236",
		//"N23.43345, E32.6457",
		//"99.234, 12.324",
		//"6.325624, 43.34345.345",
		//"0, 1,2",
		//"0.342q0832, 1.2324",
		"23.245, 1e1",
	}

	It("should invalidate the coordinates", func() {
		for _, coordinates := range invalidCoordinates {
			Expect(IsValidCoordinates(coordinates)).To(Equal(false))
		}
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Coordinates Validator")
}
