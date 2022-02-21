/*

Question:

Given an array of integers, return a new array with each value doubled.

For example:

[1, 2, 3] --> [2, 4, 6]


Solution Explanation:

- Multiplying each element of the given array by 2 and creating a new array without breaking the order.

*/


package kata

func Maps(x []int) []int {
  var s []int
  for i:=0; i<len(x); i++{
    temp:=2*x[i]
    s = append(s, temp)
  }
  return s
}

/*

Test:

var _ = Describe("Example tests", func() {
  It("should test that the solution returns the correct value", func() {
    Expect(Maps([]int{1, 2, 3})).To(Equal([]int{2, 4, 6}))
    Expect(Maps([]int{4, 1, 1, 1, 4})).To(Equal([]int{8, 2, 2, 2, 8}))
    Expect(Maps([]int{2, 2, 2, 2, 2, 2})).To(Equal([]int{4, 4, 4, 4, 4, 4}))
  })
})

*/
