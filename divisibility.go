 /*
 
- Question

Create a function that checks if a number n is divisible by two numbers x AND y. All inputs are positive, non-zero digits.
 
Examples:
1) n =   3, x = 1, y = 3 =>  true because   3 is divisible by 1 and 3
2) n =  12, x = 2, y = 6 =>  true because  12 is divisible by 2 and 6
3) n = 100, x = 5, y = 3 => false because 100 is not divisible by 3
4) n =  12, x = 7, y = 5 => false because  12 is neither divisible by 7 nor 5

- Solution Explanation

- If a given number is equal to 0 after dividing by x and y, that number is divisible by x and y. 

*/

package kata

func IsDivisible(n, x, y int) bool {
    return n%x==0 && n%y==0 
}

/*

- Test

var _ = Describe("Basic tests", func() {
    It("IsDivisible(3, 3, 4)", func() { Expect(IsDivisible(3, 3, 4)).To(Equal(false)) })
    It("IsDivisible(12, 3, 4)", func() { Expect(IsDivisible(12, 3, 4)).To(Equal(true)) })
    It("IsDivisible(8, 3, 4)", func() { Expect(IsDivisible(8, 3, 4)).To(Equal(false)) })
    It("IsDivisible(48, 3, 4)", func() { Expect(IsDivisible(48, 3, 4)).To(Equal(true)) })
    It("IsDivisible(100, 5, 10)", func() { Expect(IsDivisible(100, 5, 10)).To(Equal(true)) })
    It("IsDivisible(100, 5, 3)", func() { Expect(IsDivisible(100, 5, 3)).To(Equal(false)) })
    It("IsDivisible(4, 4, 2)", func() { Expect(IsDivisible(4, 4, 2)).To(Equal(true)) })
    It("IsDivisible(5, 2, 3)", func() { Expect(IsDivisible(5, 2, 3)).To(Equal(false)) })
    It("IsDivisible(17, 17, 17)", func() { Expect(IsDivisible(17, 17, 17)).To(Equal(true)) })
    It("IsDivisible(17, 1, 17)", func() { Expect(IsDivisible(17, 1, 17)).To(Equal(true)) })
})

*/
