package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/human_readable_time"
	"testing"
)

var _ = Describe("Tests", func() {
	It("Sample tests", func() {
		Expect(HumanReadableTime(0)).To(Equal("00:00:00"))
		Expect(HumanReadableTime(59)).To(Equal("00:00:59"))
		Expect(HumanReadableTime(60)).To(Equal("00:01:00"))
		Expect(HumanReadableTime(90)).To(Equal("00:01:30"))
		Expect(HumanReadableTime(3599)).To(Equal("00:59:59"))
		Expect(HumanReadableTime(3600)).To(Equal("01:00:00"))
		Expect(HumanReadableTime(45296)).To(Equal("12:34:56"))
		Expect(HumanReadableTime(86399)).To(Equal("23:59:59"))
		Expect(HumanReadableTime(86400)).To(Equal("24:00:00"))
		Expect(HumanReadableTime(359999)).To(Equal("99:59:59"))
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Human Readable Time")
}
