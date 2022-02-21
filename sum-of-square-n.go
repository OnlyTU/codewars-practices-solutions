 /*
 
Question:

Complete the square sum function so that it squares each number passed into it and then sums the results together.

For example, for [1, 2, 2] it should return 9 because 1^2 + 2^2 + 2^2 = 9.

 
Solution Explanation:

- If we multiply each element of the given array by itself and add up the obtained values, we can find the expected answer.

*/


package kata

func SquareSum(numbers []int) int {
  result := 0
  
  for _, n := range numbers {
    result += n * n
  }
  
  return result
}

 /*
 
Test:

var _ = Describe("Sample Tests",func() {
    It("Testing [1,2]",func() {Expect(SquareSum([]int{1,2})).To(Equal(5))})
    It("Testing [0,3,4,5]",func() {Expect(SquareSum([]int{0,3,4,5})).To(Equal(50))})
    It("Testing []",func() {Expect(SquareSum([]int{})).To(Equal(0))})
})

*/
