package slice_factor

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("SliceFactor", func() {
	DescribeTable("with a valid slice",
	    func(input []int, expected []int) {
	        Expect(sliceFactor(input)).To(BeEquivalentTo(expected))
        },
		Entry("empty slice", []int{}, []int{}),
		Entry("example 1", []int{1, 2, 3, 4, 5}, []int{120, 60, 40, 30, 24}),
		Entry("example 2", []int{3, 2, 1}, []int{2, 3, 6}),
	)
})
