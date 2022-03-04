package sumofsquaren

func squareSum(numbers []int) int {
	result := 0

	for _, n := range numbers {
		result += n * n
	}

	return result
}
