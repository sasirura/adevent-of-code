package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T)  {

	b, err := os.ReadFile("test.txt")

	assert.NoError(t, err)
	tests := []struct {
		expected int
		input []byte
	}{
		{
			expected: 142,
			input: b, 
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, part1(test.input), test.input)
	}
	
}
