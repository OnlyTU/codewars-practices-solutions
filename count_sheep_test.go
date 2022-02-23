package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sample Test Cases", func() {
	It("The solution should return the correct values for the sample test cases!", func() {
		Expect(countSheep(2)).To(Equal("1 sheep...2 sheep..."))
		Expect(countSheep(0)).To(Equal(""))
		Expect(countSheep(1)).To(Equal("1 sheep..."))
	})
})


