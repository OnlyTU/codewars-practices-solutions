/*

Question:

Terminal game move function
In this game, the hero moves from left to right. The player rolls the die and moves the number of spaces indicated by the die two times.

Create a function for the terminal game that takes the current position of the hero and the roll (1-6) and return the new position.

Example:
move(3, 6) should equal 15


Solution Explanation:

- The hero must move 2 times the value of the rolled dice. 

*/


package kata

func Move(position int, roll int) int {
    return position + 2 * roll
}


