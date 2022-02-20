 /*
 
Solution Explanation

- If a given number is equal to 0 after dividing by x and y, that number is divisible by x and y. 

*/

package kata

func IsDivisible(n, x, y int) bool {
    return n%x==0 && n%y==0 
}
