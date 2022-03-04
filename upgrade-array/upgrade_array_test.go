package upgradearray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UpgradeArray(t *testing.T) {
	assert.Equal(t, []int{2, 4, 6}, Maps([]int{1, 2, 3}))
	assert.Equal(t, []int{8, 2, 2, 2, 8}, Maps([]int{4, 1, 1, 1, 4}))
	assert.Equal(t, []int{4, 4, 4, 4, 4, 4}, Maps([]int{2, 2, 2, 2, 2, 2}))
}
