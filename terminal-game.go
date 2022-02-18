/*

Solution Explanation

- The hero must move 2 times the value of the rolled dice. 

*/


package kata

func Move(position int, roll int) int {
    return position + 2*roll
}
