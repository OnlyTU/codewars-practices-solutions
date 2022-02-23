package kata_test
import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "codewarrior/kata"
)
var _ = Describe("Test Examples", func() {
	It("Should return the correct values for the sample test cases!", func() {
		Expect(Litres(2)).To(Equal(1))
    Expect(Litres(1.4)).To(Equal(0))
    Expect(Litres(12.3)).To(Equal(6))
    Expect(Litres(0.82)).To(Equal(0))
    Expect(Litres(11.8)).To(Equal(5))
    Expect(Litres(1787)).To(Equal(893))
    Expect(Litres(0)).To(Equal(0))
  })
})
