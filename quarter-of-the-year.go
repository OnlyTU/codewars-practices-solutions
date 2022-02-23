/*

- Question

Given a month as an integer from 1 to 12, return to which quarter of the year it belongs as an integer number.

For example: month 2 (February), is part of the first quarter; month 6 (June), is part of the second quarter; and month 11 (November), is part of the fourth quarter.


- Solution Explanation

- If we create an array and subtract 1 from the given month value, we can find which quarter we are in..

*/


package kata

func QuarterOf(month int) int {
  quarter := [12] int {1,1,1,2,2,2,3,3,3,4,4,4} 
  return quarter[month - 1]
}


