package whichevennumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_WhichEvenNumber(t *testing.T) {
	assert.Equal(t, 0, NthEven(1))
	assert.Equal(t, 4, NthEven(3))
	assert.Equal(t, 198, NthEven(100))
	assert.Equal(t, 2597466, NthEven(1298734))

}
