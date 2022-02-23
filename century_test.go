package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Basic tests", func() {
	It("should return the correct values", func() {
		Expect(century(int(1990))).To(Equal(20))
		Expect(century(int(1705))).To(Equal(18))
		Expect(century(int(1900))).To(Equal(19))
		Expect(century(int(1601))).To(Equal(17))
		Expect(century(int(2000))).To(Equal(20))
		Expect(century(int(89))).To(Equal(1))
	})
})


