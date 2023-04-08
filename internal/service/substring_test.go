package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxSubstring(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Test 1", "abcabcbb", "abc"},
		{"Test 2", "bbbbb", "b"},
		{"Test 3", "pwwkew", "wke"},
		{"Test 4", "", ""},
		{"Test 5", "a", "a"},
		{"Test 6", "ab", "ab"},
		{"Test 7", "aab", "ab"},
		{"Test 8", "abcabcd", "abcd"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := FindMaxSubstring(test.input)
			if result != test.expected {
				assert.Equal(t, result, test.input)
			}
		})

	}
}
