package century

func century(year int) int {
	if year < 100 {
		return 1
	}

	if year%10 != 0 || year%100 != 0 {
		return int(year/100 + 1)
	}

	return int(year / 100)
}
