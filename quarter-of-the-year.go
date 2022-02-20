/*

Solution Explanation

- In order to calculate which quarter the given month belongs to, it will be sufficient to find between which months that month belongs.

*/


package kata

func QuarterOf(month int) int {
  if month > 9 && month < 13{ return 4 }
  if month > 6 && month < 10{ return 3 }
  if month > 3 && month < 7{ return 2 }
  return 1
}
