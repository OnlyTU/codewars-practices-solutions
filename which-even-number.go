/*

Question:

Return the Nth Even Number

Example(Input --> Output):

1 --> 0 (the first even number is 0)
3 --> 4 (the 3rd even number is 4 (0, 2, 4))
100 --> 198
1298734 --> 2597466

The input will not be 0.


Solution Explanation:

- If we multiply the given number by 2 and subtract 2, we can find which even number it corresponds to. 

*/


package kata

func NthEven(n int) int {
  return 2 * n - 2
}

/*

Test:

var _ = Describe("Basic Tests", func() {
    It("Testing for 1", func() { Expect(NthEven(1)).To(Equal(0)) })
    It("Testing for 2", func() { Expect(NthEven(2)).To(Equal(2)) })
    It("Testing for 3", func() { Expect(NthEven(3)).To(Equal(4)) })
    It("Testing for 100", func() { Expect(NthEven(100)).To(Equal(198)) })
    It("Testing for 1298734", func() { Expect(NthEven(1298734)).To(Equal(2597466)) })
})

*/
