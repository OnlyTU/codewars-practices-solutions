package terminalgame

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TerminalGame(t *testing.T) {
	assert.Equal(t, 8, terminalGame(0, 4))
	assert.Equal(t, 15, terminalGame(3, 6))
	assert.Equal(t, 12, terminalGame(2, 5))
}
