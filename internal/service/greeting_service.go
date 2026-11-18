package service

import (
	"fmt"
	"go-hello-world-service/internal/domain"
	"strings"
)

// GreetingService defines the interface for greeting operations
type GreetingService interface {
	GenerateGreeting(name string) (*domain.GreetingResponse, error)
}

// greetingService implements GreetingService
type greetingService struct{}

// NewGreetingService creates a new instance of GreetingService
func NewGreetingService() GreetingService {
	return &greetingService{}
}

// GenerateGreeting validates the name and generates an appropriate greeting
func (s *greetingService) GenerateGreeting(name string) (*domain.GreetingResponse, error) {
	name = strings.TrimSpace(name)

	// Create a greeting request
	req := &domain.GreetingRequest{
		Name: name,
	}

	// Validate the request
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// Capitalize the first letter and normalize the rest of the name.
	normalizedName := strings.ToLower(name)
	capitalizedName := strings.ToUpper(normalizedName[:1]) + normalizedName[1:]

	// Generate the greeting message
	message := fmt.Sprintf("Hello %s", capitalizedName)

	return &domain.GreetingResponse{
		Message: message,
	}, nil
}
