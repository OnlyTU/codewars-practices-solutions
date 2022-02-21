/*

- Question

Given a month as an integer from 1 to 12, return to which quarter of the year it belongs as an integer number.

For example: month 2 (February), is part of the first quarter; month 6 (June), is part of the second quarter; and month 11 (November), is part of the fourth quarter.


- Solution Explanation

- In order to calculate which quarter the given month belongs to, it will be sufficient to find between which months that month belongs.

*/


package kata

func QuarterOf(month int) int {
  if month > 9 && month < 13{ return 4 }
  if month > 6 && month < 10{ return 3 }
  if month > 3 && month < 7{ return 2 }
  return 1
}

/*

- Test

var _ = Describe("Test Example", func() {
  It("should test example values", func() {
    Expect(QuarterOf(3)).To(Equal(1))
    Expect(QuarterOf(8)).To(Equal(3))
    Expect(QuarterOf(11)).To(Equal(4))
  })
})


*/
