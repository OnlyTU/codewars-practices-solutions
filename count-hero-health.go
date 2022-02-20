/*

Solution Explanation

- If we subtract the damage done to the hero from the hero's health, we can find our hero's instant health. 

*/


package kata

func combat(health, damage float64) float64 {
  
  if health < damage { return 0.0 }
  
  return health - damage
}
