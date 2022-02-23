package kata_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "codewarrior/kata"
)

var _ = Describe("Test Example", func() {
  It("should test example values", func() {
    Expect(QuarterOf(3)).To(Equal(1))
    Expect(QuarterOf(8)).To(Equal(3))
    Expect(QuarterOf(11)).To(Equal(4))
  })
})


