package quarteryear

func QuarterOf(month int) int {
	quarter := [12]int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4}
	return quarter[month-1]
}
