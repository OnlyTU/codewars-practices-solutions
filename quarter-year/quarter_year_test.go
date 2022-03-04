package quarteryear

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_QuarterOfYear(t *testing.T) {
	assert.Equal(t, 1, QuarterOf(3))
	assert.Equal(t, 3, QuarterOf(8))
	assert.Equal(t, 4, QuarterOf(11))
	assert.Equal(t, 2, QuarterOf(6))
}
