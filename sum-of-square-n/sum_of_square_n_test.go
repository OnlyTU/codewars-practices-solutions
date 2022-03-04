package sumofsquaren

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SumOfSquareN(t *testing.T) {
	assert.Equal(t, 5, squareSum([]int{1, 2}))
	assert.Equal(t, 50, squareSum([]int{0, 3, 4, 5}))
	assert.Equal(t, 0, squareSum([]int{}))

}
