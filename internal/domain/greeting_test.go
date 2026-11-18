package domain

import (
	"testing"
)

func TestGreetingRequest_Validate(t *testing.T) {
	tests := []struct {
		name        string
		requestName string
		wantErr     error
	}{
		{
			name:        "Valid name starting with A",
			requestName: "Alice",
			wantErr:     nil,
		},
		{
			name:        "Valid name starting with M",
			requestName: "Mary",
			wantErr:     nil,
		},
		{
			name:        "Valid name starting with lowercase a",
			requestName: "alice",
			wantErr:     nil,
		},
		{
			name:        "Invalid name starting with N",
			requestName: "Nancy",
			wantErr:     ErrInvalidInput,
		},
		{
			name:        "Invalid name starting with Z",
			requestName: "Zane",
			wantErr:     ErrInvalidInput,
		},
		{
			name:        "Invalid name starting with lowercase z",
			requestName: "zoe",
			wantErr:     ErrInvalidInput,
		},
		{
			name:        "Empty name",
			requestName: "",
			wantErr:     ErrEmptyName,
		},
		{
			name:        "Name with only spaces",
			requestName: "   ",
			wantErr:     ErrEmptyName,
		},
		{
			name:        "Valid name at boundary (M)",
			requestName: "Michael",
			wantErr:     nil,
		},
		{
			name:        "Invalid name at boundary (N)",
			requestName: "Nathan",
			wantErr:     ErrInvalidInput,
		},
		{
			name:        "Invalid name starting with a digit",
			requestName: "1Alice",
			wantErr:     ErrInvalidInput,
		},
		{
			name:        "Invalid name starting with a symbol",
			requestName: "@Alice",
			wantErr:     ErrInvalidInput,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &GreetingRequest{
				Name: tt.requestName,
			}

			err := req.Validate()

			if err != tt.wantErr {
				t.Errorf("GreetingRequest.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsFirstHalfAlphabet(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{
			name: "First half - lowercase a",
			arg:  "alice",
			want: true,
		},
		{
			name: "First half - uppercase A",
			arg:  "Alice",
			want: true,
		},
		{
			name: "First half - m boundary",
			arg:  "Mary",
			want: true,
		},
		{
			name: "Second half - n boundary",
			arg:  "Nancy",
			want: false,
		},
		{
			name: "Second half - z",
			arg:  "Zane",
			want: false,
		},
		{
			name: "Empty string",
			arg:  "",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFirstHalfAlphabet(tt.arg); got != tt.want {
				t.Errorf("IsFirstHalfAlphabet() = %v, want %v", got, tt.want)
			}
		})
	}
}
