package domain

import (
	"errors"
	"strings"
)

var (
	// ErrInvalidInput is returned when the name validation fails
	ErrInvalidInput = errors.New("Invalid Input")
	// ErrEmptyName is returned when the name parameter is missing or empty
	ErrEmptyName = errors.New("name parameter is required")
)

// GreetingRequest represents the input for greeting generation
type GreetingRequest struct {
	Name string
}

// GreetingResponse represents a successful greeting response
type GreetingResponse struct {
	Message string `json:"message"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// Validate checks if the name is valid and returns an error if not
func (g *GreetingRequest) Validate() error {
	name := strings.TrimSpace(g.Name)
	if name == "" {
		return ErrEmptyName
	}

	firstChar := name[0]
	if firstChar >= 'A' && firstChar <= 'Z' {
		firstChar += 'a' - 'A'
	}
	if firstChar < 'a' || firstChar > 'z' {
		return ErrInvalidInput
	}

	if firstChar >= 'n' && firstChar <= 'z' {
		return ErrInvalidInput
	}

	return nil
}

// IsFirstHalfAlphabet checks if the first character is in A-M range
func IsFirstHalfAlphabet(name string) bool {
	name = strings.TrimSpace(name)
	if name == "" {
		return false
	}

	firstChar := name[0]
	if firstChar >= 'A' && firstChar <= 'Z' {
		firstChar += 'a' - 'A'
	}
	return firstChar >= 'a' && firstChar <= 'm'
}
