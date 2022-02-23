package kata_test
import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "codewarrior/kata"
)

var _ = Describe("Basic Tests", func() {
    It("Move(0, 4)", func() { Expect(Move(0, 4)).To(Equal(8)) })
    It("Move(3, 6)", func() { Expect(Move(3, 6)).To(Equal(15)) })
    It("Move(2, 5)", func() { Expect(Move(2, 5)).To(Equal(12)) })
})


