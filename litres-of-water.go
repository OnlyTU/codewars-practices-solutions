/*

Solution Explanation

- To calculate the amount of water drunk per hour, it is sufficient to multiply the amount of water drunk (0.5 liters) and the time.

*/

package kata

func Litres(time float64) int {
  return int(time*0.5)
}
