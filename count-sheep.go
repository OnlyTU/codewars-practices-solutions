/*

Solution Explanation

- First I returned its value based on the condition of being 0.
- I made a string definition called "temp" to keep the repeat count.
- I found how many times it was written with a for loop.
- I converted every value I found to integer to string, then returned the result.

*/

package kata

import(
  "strconv"
)

func countSheep(num int) string {
  
  if num == 0 {
    return ""
  }
  
  temp := ""
  for i := 1; num >= i; i++ {    
      s := strconv.Itoa(i)
      temp += s + " sheep..."
    }
  return temp 
}
