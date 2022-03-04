package upgradearray

func Maps(x []int) []int {
	var s []int
	for i := 0; i < len(x); i++ {
		temp := 2 * x[i]
		s = append(s, temp)
	}
	return s
}
