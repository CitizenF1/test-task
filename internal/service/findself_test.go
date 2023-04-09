package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSelf(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Test one",
			input:    "content",
			expected: []string{"content", "content"},
		},
		{
			name:     "Test two",
			input:    "matches",
			expected: []string{"matches", "matches", "matches", "matches"},
		},
		{
			name:     "Test tree",
			input:    "left",
			expected: []string{"left", "left", "left", "left"},
		},
		{
			name:     "Test four",
			input:    "FindMaxSubstring",
			expected: []string{"FindMaxSubstring"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := FindSelf(test.input)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, test.expected, result)
		})
	}
}
