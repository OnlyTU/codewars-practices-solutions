/*

Solution Explanation

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
