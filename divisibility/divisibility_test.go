package divisibility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsDivisible(t *testing.T) {
	assert.Equal(t, false, isDivisible(3, 3, 4))
	assert.Equal(t, true, isDivisible(12, 3, 4))
	assert.Equal(t, false, isDivisible(8, 3, 4))
	assert.Equal(t, true, isDivisible(48, 3, 4))
	assert.Equal(t, true, isDivisible(100, 5, 10))
	assert.Equal(t, false, isDivisible(100, 5, 3))
	assert.Equal(t, false, isDivisible(5, 2, 3))
	assert.Equal(t, true, isDivisible(17, 17, 17))
}
