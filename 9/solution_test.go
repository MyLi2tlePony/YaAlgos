package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsValid(t *testing.T) {
	require.True(t, isValid("()(){}{()}"))
	require.False(t, isValid("()({){}{()}"))
}
