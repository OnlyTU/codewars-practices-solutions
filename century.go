/*

Solution Explanation

- Calculation of century.

- If the remainder from the division by 10 and 100 of the given year is not 0, that year falls within the next century.

*/


package kata

func century(year int) int {
  
  if year < 100 { 
    return 1 
  }
 
  if year % 10 != 0 || year % 100 != 0 {
    return int (year / 100 + 1)
  }
  
  return int (year / 100)
}
