package kata_test
import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "codewarrior/kata"
)
var _ = Describe("Sample Test Cases", func() {
  It("should work for sample tests", func() {
    Expect(Expect(SmallestIntegerFinder([]int{34, 15, 88, 2})).To(Equal(2)))
    Expect(Expect(SmallestIntegerFinder([]int{34, -345, -1, 100})).To(Equal(-345)))
  })
})


