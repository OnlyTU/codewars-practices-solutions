   
/*

Solution Explanation

- Every dragons take 2 bullets that's why 2 times the number of dragons must be equal to or less than the number of bullets.

*/


package kata

func Hero(bullets, dragons int) bool {
  if (dragons * 2 <= bullets ) {
    return true
  }
    return false 
}


// Single line solution

/*

func Hero(bullets, dragons int) bool {
  return dragons * 2 <= bullets
}

*/
