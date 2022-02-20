 /*
 
Solution Explanation

- If we multiply each element of the given array by itself and add up the obtained values, we can find the expected answer.

*/


package kata

func SquareSum(numbers []int) int {
  
  temp := 0
  
  if len(numbers) != 0 {
    for i := 0; len(numbers) > i; i++ {
      temp += numbers[i] * numbers[i]
    }       
    return temp
  }  
  return 0
}
