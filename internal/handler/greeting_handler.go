package handler

import (
	"go-hello-world-service/internal/domain"
	"go-hello-world-service/internal/service"
	"go-hello-world-service/pkg/response"
	"net/http"
)

// GreetingHandler handles HTTP requests for greeting operations
type GreetingHandler struct {
	greetingService service.GreetingService
}

// NewGreetingHandler creates a new instance of GreetingHandler
func NewGreetingHandler(greetingService service.GreetingService) *GreetingHandler {
	return &GreetingHandler{
		greetingService: greetingService,
	}
}

// HelloWorld handles the /hello-world endpoint
// GET /hello-world?name=<name>
func (h *GreetingHandler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	// Only allow GET requests
	if r.Method != http.MethodGet {
		response.Error(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Get the name query parameter
	name := r.URL.Query().Get("name")

	// Call the service to generate greeting
	greeting, err := h.greetingService.GenerateGreeting(name)

	if err != nil {
		// Check the type of error and respond accordingly
		switch err {
		case domain.ErrEmptyName:
			response.Error(w, http.StatusBadRequest, "Invalid Input")
		case domain.ErrInvalidInput:
			response.Error(w, http.StatusBadRequest, "Invalid Input")
		default:
			response.Error(w, http.StatusInternalServerError, "Internal Server Error")
		}
		return
	}

	// Return successful response
	response.Success(w, greeting)
}
