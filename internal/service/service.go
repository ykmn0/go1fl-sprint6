package service

import (
	"errors"
	"strings"

	"github.com/ykmn0/go1fl-sprint6/pkg/morse"
)

// Convert automatically detects if input is morse code or text and converts accordingly
func Convert(input string) (string, error) {
	if input == "" {
		return "", errors.New("empty input")
	}

	// Simple heuristic: Morse code typically contains dots, dashes and spaces
	// If the string mostly contains these characters, it's likely morse code
	dotsDashes := strings.Count(input, ".") + strings.Count(input, "-")
	spaces := strings.Count(input, " ")
	totalLen := len(input)

	// If more than 30% of characters are dots/dashes, assume it's morse code
	if float64(dotsDashes+spaces)/float64(totalLen) > 0.3 {
		// Convert morse to text
		return morse.ToText(input), nil
	}

	// Otherwise assume it's text and convert to morse
	return morse.ToMorse(input), nil
}
