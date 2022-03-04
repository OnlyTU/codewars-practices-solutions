package divisibility

func isDivisible(n, x, y int) bool {
	return n%x == 0 && n%y == 0
}
