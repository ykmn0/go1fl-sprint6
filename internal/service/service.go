package service

import (
	"errors"
	"strings"

	"go1fl-sprint6/pkg/morse"
)

// Convert automatically detects if input is morse code or text and converts accordingly
func Convert(input string) (string, error) {
	if input == "" {
		return "", errors.New("empty input")
	}

	isMorse := !strings.ContainsFunc(input, func(r rune) bool {
		return !strings.ContainsRune(".- /", r)
	})

	if isMorse {
		return morse.ToText(input), nil
	}

	return morse.ToMorse(input), nil
}
