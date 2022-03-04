package smallestelement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SmallestElement(t *testing.T) {
	assert.Equal(t, 2, SmallestIntegerFinder([]int{34, 15, 88, 2}))
	assert.Equal(t, -345, SmallestIntegerFinder([]int{34, -345, -1, 100}))
}
