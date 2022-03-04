package herosurvive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HeroSurvive(t *testing.T) {
	assert.Equal(t, true, Hero(10, 5))
	assert.Equal(t, false, Hero(7, 4))
	assert.Equal(t, true, Hero(100, 40))
	assert.Equal(t, false, Hero(1500, 751))
	assert.Equal(t, false, Hero(0, 1))
}
