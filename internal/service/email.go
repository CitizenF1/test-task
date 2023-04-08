package service

import (
	"regexp"
	"strings"
)

func FindEmail(text string) []string {
	// regex для поиска валидных email в строке
	var emailReg = regexp.MustCompile(`\s*([a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,})`)
	// Check Email
	matches := emailReg.FindAllString(text, -1)

	var result []string

	for _, match := range matches {
		buffer := strings.TrimSpace(match)
		result = append(result, buffer)
	}
	return result
}

func Findiin(text string) []string {
	// regex для поиска валидных иин в строке
	var iinReg = regexp.MustCompile(`\b(\d{12})\b`)
	// Check Email
	matches := iinReg.FindAllString(text, -1)

	var result []string

	for _, match := range matches {
		buffer := strings.TrimSpace(match)
		result = append(result, buffer)
	}
	return result
}
