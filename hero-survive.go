   
/*

- Question

A hero is on his way to the castle to complete his mission. However, he's been told that the castle is surrounded with a couple of powerful dragons! each dragon takes 2 bullets to be defeated, our hero has no idea how many bullets he should carry.. Assuming he's gonna grab a specific given number of bullets and move forward to fight another specific given number of dragons, will he survive?

Return True if yes, False otherwise :)

- Solution Explanation

- Every dragons take 2 bullets that's why 2 times the number of dragons must be equal to or less than the number of bullets.

*/


package kata

func Hero(bullets, dragons int) bool {
  return dragons * 2 <= bullets
}


/*

- Test

var _ = Describe("Fixed tests", func() {
   It("should test that the solution returns the correct value", func() {
     Expect(Hero(10, 5)).To(Equal(true))
     Expect(Hero(7, 4)).To(Equal(false))
     Expect(Hero(4, 5)).To(Equal(false))
     Expect(Hero(100, 40)).To(Equal(true))
     Expect(Hero(1500, 751)).To(Equal(false))
     Expect(Hero(0, 1)).To(Equal(false))
   })
})

*/
