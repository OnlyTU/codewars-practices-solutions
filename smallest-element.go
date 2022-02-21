/*

Question:

Given an array of integers your solution should find the smallest integer.

For example:

Given [34, 15, 88, 2] your solution will return 2
Given [34, -345, -1, 100] your solution will return -345
You can assume, for the purpose of this kata, that the supplied array will not be empty.

Solution Explanation:

- We can assume the first element of the given array as the smallest and compare it with the other elements to find the smallest number.

*/


package kata

func SmallestIntegerFinder(numbers []int) int {
  min:=0
  min=numbers[0]
  
  for i:=0; len(numbers)>i; i++{
    if min>numbers[i]{
      min=numbers[i]
    }
  }
  return min 
}

/*

Test:

var _ = Describe("Sample Test Cases", func() {
  It("should work for sample tests", func() {
    Expect(Expect(SmallestIntegerFinder([]int{34, 15, 88, 2})).To(Equal(2)))
    Expect(Expect(SmallestIntegerFinder([]int{34, -345, -1, 100})).To(Equal(-345)))
  })
})

*/
