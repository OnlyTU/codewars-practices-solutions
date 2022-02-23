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
  for i := 0; i < len(x); i++ {
    temp := 2 * x[i]
    s = append(s, temp)
  }
  return s
}


