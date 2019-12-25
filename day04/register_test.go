package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewRegister(t *testing.T) {
	r, err := NewRegister(307237)
	require.NoError(t, err)
	assert.Equal(t, [lenPass]int{3, 0, 7, 2, 3, 7}, r.data)

	_, err = NewRegister(111)
	require.Error(t, err)
}

func TestRegister_ToDec(t *testing.T) {
	r, err := NewRegister(307237)
	require.NoError(t, err)
	assert.Equal(t, 307237, r.ToDec())
}

func TestRegister_Next(t *testing.T) {
	r, err := NewRegister(307237)
	require.NoError(t, err)

	next := r.Next()
	assert.Equal(t, [lenPass]int{3, 0, 7, 2, 3, 8}, next.data)

	next = next.Next()
	assert.Equal(t, [lenPass]int{3, 0, 7, 2, 3, 9}, next.data)

	next = next.Next()
	assert.Equal(t, [lenPass]int{3, 0, 7, 2, 4, 4}, next.data)

	r, err = NewRegister(199999)
	require.NoError(t, err)

	next = r.Next()
	assert.Equal(t, [lenPass]int{2, 2, 2, 2, 2, 2}, next.data)
}

func TestRegister_IsValidPass(t *testing.T) {
	r, err := NewRegister(111111)
	require.NoError(t, err)
	assert.True(t, r.IsValidPass())

	r, err = NewRegister(223450)
	require.NoError(t, err)
	assert.False(t, r.IsValidPass())

	r, err = NewRegister(123789)
	require.NoError(t, err)
	assert.False(t, r.IsValidPass())
}
