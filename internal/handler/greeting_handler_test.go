package handler

import (
	"encoding/json"
	"go-hello-world-service/internal/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGreetingHandler_HelloWorld(t *testing.T) {
	// Initialize dependencies
	greetingService := service.NewGreetingService()
	handler := NewGreetingHandler(greetingService)

	tests := []struct {
		name           string
		method         string
		queryParam     string
		expectedStatus int
		expectedBody   map[string]string
	}{
		{
			name:           "Valid request - Alice",
			method:         http.MethodGet,
			queryParam:     "Alice",
			expectedStatus: http.StatusOK,
			expectedBody:   map[string]string{"message": "Hello Alice"},
		},
		{
			name:           "Valid request - lowercase alice",
			method:         http.MethodGet,
			queryParam:     "alice",
			expectedStatus: http.StatusOK,
			expectedBody:   map[string]string{"message": "Hello Alice"},
		},
		{
			name:           "Valid request - Bob",
			method:         http.MethodGet,
			queryParam:     "bob",
			expectedStatus: http.StatusOK,
			expectedBody:   map[string]string{"message": "Hello Bob"},
		},
		{
			name:           "Valid request - Mary (valid boundary)",
			method:         http.MethodGet,
			queryParam:     "Mary",
			expectedStatus: http.StatusOK,
			expectedBody:   map[string]string{"message": "Hello Mary"},
		},
		{
			name:           "Invalid request - Zane",
			method:         http.MethodGet,
			queryParam:     "Zane",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   map[string]string{"error": "Invalid Input"},
		},
		{
			name:           "Invalid request - Nancy",
			method:         http.MethodGet,
			queryParam:     "Nancy",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   map[string]string{"error": "Invalid Input"},
		},
		{
			name:           "Missing name parameter",
			method:         http.MethodGet,
			queryParam:     "",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   map[string]string{"error": "Invalid Input"},
		},
		{
			name:           "Invalid name starting with a digit",
			method:         http.MethodGet,
			queryParam:     "1Alice",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   map[string]string{"error": "Invalid Input"},
		},
		{
			name:           "Invalid HTTP method - POST",
			method:         http.MethodPost,
			queryParam:     "Alice",
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   map[string]string{"error": "Method not allowed"},
		},
		{
			name:           "Invalid HTTP method - PUT",
			method:         http.MethodPut,
			queryParam:     "Alice",
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   map[string]string{"error": "Method not allowed"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create request
			url := "/hello-world"
			if tt.queryParam != "" {
				url += "?name=" + tt.queryParam
			}
			req := httptest.NewRequest(tt.method, url, nil)

			// Create response recorder
			w := httptest.NewRecorder()

			// Call handler
			handler.HelloWorld(w, req)

			// Check status code
			if w.Code != tt.expectedStatus {
				t.Errorf("HelloWorld() status = %v, want %v", w.Code, tt.expectedStatus)
			}

			// Check response body
			var response map[string]string
			if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
				t.Fatalf("Failed to decode response body: %v", err)
			}

			// Compare expected body with actual response
			for key, expectedValue := range tt.expectedBody {
				if actualValue, ok := response[key]; !ok || actualValue != expectedValue {
					t.Errorf("HelloWorld() body[%s] = %v, want %v", key, actualValue, expectedValue)
				}
			}
		})
	}
}

func TestGreetingHandler_HelloWorld_Integration(t *testing.T) {
	// This test simulates the complete flow
	greetingService := service.NewGreetingService()
	handler := NewGreetingHandler(greetingService)

	// Test successful scenario
	t.Run("Integration test - successful greeting", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/hello-world?name=Alice", nil)
		w := httptest.NewRecorder()

		handler.HelloWorld(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", w.Code)
		}

		var response map[string]string
		json.NewDecoder(w.Body).Decode(&response)

		if response["message"] != "Hello Alice" {
			t.Errorf("Expected message 'Hello Alice', got '%s'", response["message"])
		}
	})

	// Test error scenario
	t.Run("Integration test - invalid name", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/hello-world?name=Zane", nil)
		w := httptest.NewRecorder()

		handler.HelloWorld(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status 400, got %d", w.Code)
		}

		var response map[string]string
		json.NewDecoder(w.Body).Decode(&response)

		if response["error"] != "Invalid Input" {
			t.Errorf("Expected error 'Invalid Input', got '%s'", response["error"])
		}
	})
}
