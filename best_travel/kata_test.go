package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go-katas/best_travel"
	"testing"
)

func dotest(t, k int, ls []int, exp int) {
	var ans = ChooseBestSum(t, k, ls)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		var ts = []int{50, 55, 56, 57, 58}
		dotest(163, 3, ts, 163)
		ts = []int{50}
		dotest(163, 3, ts, -1)
		ts = []int{91, 74, 73, 85, 73, 81, 87}
		dotest(230, 3, ts, 228)
		dotest(331, 2, ts, 178)
		dotest(331, 4, ts, 331)
		dotest(331, 5, ts, -1)
	})

	//It("Performance test", func() {
	//	ts := []int{931, 744, 763, 825, 713, 851, 867, 435, 726, 112, 847, 697, 547, 384, 897, 455, 242, 564, 873,
	//		586, 336, 934, 208, 557, 876, 644, 934, 753, 664, 725, 941, 754, 563, 972, 683, 636, 428, 943,
	//		653, 862, 929, 645, 251, 665, 120, 750, 420, 340, 931, 744, 763, 825, 713, 851, 867, 435, 726,
	//		112, 847, 697, 547, 384, 897, 455, 242, 564, 873, 586, 336, 934, 208, 557, 876, 644, 934, 753,
	//		664, 725, 941, 754, 563, 972, 683, 636, 428, 943, 653, 862, 929, 645, 251, 665, 120, 750, 420,
	//		340, 931, 744, 763, 825, 713, 851, 867, 435, 726, 112, 847, 697, 547, 384, 897, 455, 242, 564,
	//		873, 586, 336, 934, 208, 557, 876, 644, 934, 753, 664, 725, 941, 754, 563, 972, 683, 636, 428,
	//		943, 653, 862, 929, 645, 251, 665, 120, 750, 420, 340, 931, 744, 763}
	//
	//	dotest(3198, 6, ts, 3198)
	//})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Best Travel")
}
