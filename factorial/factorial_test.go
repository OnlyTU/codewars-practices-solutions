package factorial

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Factorial(t *testing.T) {
	assert.Equal(t, 120, factorial(5))
	assert.Equal(t, 1, factorial(0))
	assert.Equal(t, 1, factorial(1))
	assert.Equal(t, 355687428096000, factorial(17))
}
