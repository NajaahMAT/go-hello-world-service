package service

import (
	"go-hello-world-service/internal/domain"
	"testing"
)

func TestGreetingService_GenerateGreeting(t *testing.T) {
	service := NewGreetingService()

	tests := []struct {
		name        string
		inputName   string
		wantMessage string
		wantErr     error
	}{
		{
			name:        "Valid name - Alice",
			inputName:   "Alice",
			wantMessage: "Hello Alice",
			wantErr:     nil,
		},
		{
			name:        "Valid name - lowercase",
			inputName:   "alice",
			wantMessage: "Hello Alice",
			wantErr:     nil,
		},
		{
			name:        "Valid name - UPPERCASE",
			inputName:   "ALICE",
			wantMessage: "Hello Alice",
			wantErr:     nil,
		},
		{
			name:        "Valid name - Bob",
			inputName:   "bob",
			wantMessage: "Hello Bob",
			wantErr:     nil,
		},
		{
			name:        "Valid name at boundary - Mary",
			inputName:   "Mary",
			wantMessage: "Hello Mary",
			wantErr:     nil,
		},
		{
			name:        "Invalid name - Zane",
			inputName:   "Zane",
			wantMessage: "",
			wantErr:     domain.ErrInvalidInput,
		},
		{
			name:        "Invalid name - Nancy",
			inputName:   "Nancy",
			wantMessage: "",
			wantErr:     domain.ErrInvalidInput,
		},
		{
			name:        "Empty name",
			inputName:   "",
			wantMessage: "",
			wantErr:     domain.ErrEmptyName,
		},
		{
			name:        "Name with spaces only",
			inputName:   "   ",
			wantMessage: "",
			wantErr:     domain.ErrEmptyName,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := service.GenerateGreeting(tt.inputName)

			// Check error
			if err != tt.wantErr {
				t.Errorf("GenerateGreeting() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// If we expect an error, response should be nil
			if tt.wantErr != nil {
				if response != nil {
					t.Errorf("GenerateGreeting() expected nil response on error, got %v", response)
				}
				return
			}

			// Check response message
			if response.Message != tt.wantMessage {
				t.Errorf("GenerateGreeting() message = %v, want %v", response.Message, tt.wantMessage)
			}
		})
	}
}
