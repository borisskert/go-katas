package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/count_ip_addresses"
	"testing"
)

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		Expect(IpsBetween("10.0.0.0", "10.0.0.50")).To(Equal(50))
		Expect(IpsBetween("20.0.0.10", "20.0.1.0")).To(Equal(246))
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Count IP Addresses")
}
