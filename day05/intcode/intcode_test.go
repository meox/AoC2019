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

func TestPosModeJmp(t *testing.T) {
	p, err := NewProgram("3,9,8,9,10,9,4,9,99,-1,8")
	require.NoError(t, err)
	p.SetInput(8)
	err = p.Run()
	require.NoError(t, err)
	assert.Equal(t, 1, p.output)

	p, err = NewProgram("3,9,8,9,10,9,4,9,99,-1,8")
	require.NoError(t, err)
	p.SetInput(5)
	err = p.Run()
	require.NoError(t, err)
	assert.Equal(t, 0, p.output)
}

func TestImModeJmp(t *testing.T) {
	p, err := NewProgram("3,3,1108,-1,8,3,4,3,99")
	require.NoError(t, err)
	p.SetInput(8)
	err = p.Run()
	require.NoError(t, err)
	assert.Equal(t, 1, p.output)

	p, err = NewProgram("3,3,1108,-1,8,3,4,3,99")
	require.NoError(t, err)
	p.SetInput(5)
	err = p.Run()
	require.NoError(t, err)
	assert.Equal(t, 0, p.output)
}

func TestJmp(t *testing.T) {
	// first example
	pr1 := "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9"
	p, err := NewProgram(pr1)
	require.NoError(t, err)
	p.SetInput(0)
	err = p.Run()
	require.NoError(t, err)
	assert.Equal(t, 0, p.output)

	p.SetInput(10)
	err = p.Run()
	require.NoError(t, err)
	assert.NotEqual(t, 0, p.output)

	// second example
	pr2 := "3,3,1105,-1,9,1101,0,0,12,4,12,99,1"
	p, err = NewProgram(pr2)
	require.NoError(t, err)
	p.SetInput(0)
	err = p.Run()
	require.NoError(t, err)
	assert.Equal(t, 0, p.output)

	p.SetInput(10)
	err = p.Run()
	require.NoError(t, err)
	assert.NotEqual(t, 0, p.output)
}

func TestExt(t *testing.T) {
	code := "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99"
	p, err := NewProgram(code)
	require.NoError(t, err)
	p.SetInput(8)
	err = p.Run()
	require.NoError(t, err)
	assert.Equal(t, 1000, p.output)

	p.SetInput(5)
	err = p.Run()
	require.NoError(t, err)
	assert.Equal(t, 999, p.output)

	p.SetInput(11)
	err = p.Run()
	require.NoError(t, err)
	assert.Equal(t, 1001, p.output)
}
