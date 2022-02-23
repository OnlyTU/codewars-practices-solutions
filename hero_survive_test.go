package kata_test
import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "codewarrior/kata"
)

var _ = Describe("Fixed tests", func() {
   It("should test that the solution returns the correct value", func() {
     Expect(Hero(10, 5)).To(Equal(true))
     Expect(Hero(7, 4)).To(Equal(false))
     Expect(Hero(4, 5)).To(Equal(false))
     Expect(Hero(100, 40)).To(Equal(true))
     Expect(Hero(1500, 751)).To(Equal(false))
     Expect(Hero(0, 1)).To(Equal(false))
   })
})


