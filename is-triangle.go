/*

Question:

Implement a function that accepts 3 integer values a, b, c. The function should return true if a triangle can be built with the sides of given length and false in any other case.

(In this case, all triangles must have surface greater than 0 to be accepted).

Solution Explanation:

- I used the Triangle Inequality Theorem, which states that the sum of the lengths of two sides of a triangle is always greater than the third side.

*/

package kata


func IsTriangle(a, b, c int) bool {
  return a+b > c && b+c > a && c+a > b
}


