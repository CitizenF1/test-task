package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindEmails(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Test one email",
			input:    "Email:  test.email@test.com",
			expected: []string{"test.email@test.com"},
		},
		{
			name:     "Test two email",
			input:    "dasdasd: Email: test.email1@test.com \n\t Email: test.email2@test.com",
			expected: []string{"test.email1@test.com", "test.email2@test.com"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := FindEmail(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestFindIIN(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Test one iin",
			input:    "asdasddwq 932312301546",
			expected: []string{"932312301546"},
		},
		{
			name:     "Test two email",
			input:    "dasdasd: Инн: 902312301548  902312301547",
			expected: []string{"902312301548", "902312301547"},
		},
		{
			name:     "Test tree email",
			input:    "dasdasd: Инн: 902312301548  902312301547 \n\t 902312301543",
			expected: []string{"902312301548", "902312301547", "902312301543"},
		},
		{
			name:     "Test empty",
			input:    "",
			expected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Findiin(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}
