package isTriangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsTriangle(t *testing.T) {
	assert.Equal(t, false, isTriangle(5, 1, 2))
	assert.Equal(t, false, isTriangle(1, 2, 5))
	assert.Equal(t, true, isTriangle(3, 4, 5))
	assert.Equal(t, true, isTriangle(5, 1, 5))
	assert.Equal(t, false, isTriangle(-1, 2, 3))
	assert.Equal(t, true, isTriangle(2, 2, 2))
}
