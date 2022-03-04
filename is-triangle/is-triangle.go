package isTriangle

func isTriangle(a, b, c int) bool {
	return a+b > c && b+c > a && c+a > b
}
