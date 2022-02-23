package kata_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "codewarrior/kata"
)
var _ = Describe("Basic Tests", func() {
    It("Testing for 1", func() { Expect(NthEven(1)).To(Equal(0)) })
    It("Testing for 2", func() { Expect(NthEven(2)).To(Equal(2)) })
    It("Testing for 3", func() { Expect(NthEven(3)).To(Equal(4)) })
    It("Testing for 100", func() { Expect(NthEven(100)).To(Equal(198)) })
    It("Testing for 1298734", func() { Expect(NthEven(1298734)).To(Equal(2597466)) })
})


