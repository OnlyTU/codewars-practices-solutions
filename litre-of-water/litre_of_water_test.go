package litreofwater

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LitreWater(t *testing.T) {
	assert.Equal(t, 0, Litres(0))
	assert.Equal(t, 0, Litres(-1))
	assert.Equal(t, 6, Litres(12.3))
	assert.Equal(t, 893, Litres(1787))
	assert.Equal(t, 0, Litres(0.82))
}
