package countsheep

import (
	"strconv"
)

func countSheep(num int) string {

	if num == 0 {
		return ""
	}

	temp := ""
	for i := 1; num >= i; i++ {
		s := strconv.Itoa(i)
		temp += s + " sheep..."
	}
	return temp
}
