package countsheep

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CountSheep(t *testing.T) {
	assert.Equal(t, "1 sheep...2 sheep...", countSheep(2))
	assert.Equal(t, "", countSheep(0))
	assert.Equal(t, "1 sheep...", countSheep(1))
	assert.Equal(t, "", countSheep(-1))
}
