package smallestelement

func SmallestIntegerFinder(numbers []int) int {
	min := 0
	min = numbers[0]

	for i := 0; len(numbers) > i; i++ {
		if min > numbers[i] {
			min = numbers[i]
		}
	}
	return min
}
