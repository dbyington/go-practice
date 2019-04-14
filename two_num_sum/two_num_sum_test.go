package two_num_sum

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TwoNumSum", func() {
	Context("with an empty slice", func() {
		It("should return false", func() {
			Expect(twoNumSum([]int{}, 0)).To(BeFalse())
		})
	})
	Context("with {5}, 5", func() {
		It("should return false", func() {
			Expect(twoNumSum([]int{5}, 5)).To(BeFalse())
        })
    })
    Context("with {5, 5}, 10", func() {
        It("should return true", func() {
            Expect(twoNumSum([]int{5, 5}, 10)).To(BeTrue())
        })
    })
    Context("with int{1, 2, 3}, 5", func() {
        It("should return true", func() {
			Expect(twoNumSum([]int{1, 2, 3}, 5)).To(BeTrue())
        })
    })
    Context("with {10, 15, 3, 7}, 17", func() {
        It("should return true", func() {
			Expect(twoNumSum([]int{10, 15, 3, 7}, 17)).To(BeTrue())
        })
    })
    Context("with {10, 15, 3, 7}, 24", func() {
        It("should return false", func() {
			Expect(twoNumSum([]int{10, 15, 3, 7}, 24)).To(BeFalse())
        })
    })
})
