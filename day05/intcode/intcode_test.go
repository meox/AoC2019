package intcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOpCode(t *testing.T) {
	assert.Equal(t, HLT, NewOpCode(HLT).Value)
	assert.Equal(t, ADD, NewOpCode(ADD).Value)

	c := NewOpCode(01002)
	assert.Equal(t, MUL, c.Value)
	assert.Equal(t, POSITION_MODE, c.C())
	assert.Equal(t, IMMEDIATE_MODE, c.B())
	assert.Equal(t, POSITION_MODE, c.A())
}
