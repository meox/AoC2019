package intcode

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestNewOpCode(t *testing.T) {
	assert.Equal(t, HLT, NewOpCode(HLT).Value)
	assert.Equal(t, ADD, NewOpCode(ADD).Value)

	c := NewOpCode(1002)
	assert.Equal(t, MUL, c.Value)
	assert.Equal(t, POSITION_MODE, c.ModeParamC())
	assert.Equal(t, IMMEDIATE_MODE, c.ModeParamB())
	assert.Equal(t, POSITION_MODE, c.ModeParamA())
}

func TestBasicProgram(t *testing.T) {
	p, err := NewProgram("1101,100,-1,4,0")
	require.NoError(t, err)
	err = p.Run()
	require.NoError(t, err)
}

func TestExampleProgram(t *testing.T) {
	p, err := NewProgram("1002,4,3,4,33")
	require.NoError(t, err)
	err = p.Run()
	require.NoError(t, err)
	assert.Equal(t, 99, p.memory[4])
}
