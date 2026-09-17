package respond

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

type APIError struct {
	Type      string `json:"type"`
	Title     string `json:"title"`
	Status    int    `json:"status"`
	Detail    string `json:"detail"`
	RequestID string `json:"requestId"`
}

func Unhandled(message string) *APIError {
	return &APIError{
		Type:   "/internal-server-error",
		Title:  "Internal Server Error",
		Status: http.StatusInternalServerError,
		Detail: message,
	}
}

func NotFound(resource string) *APIError {
	return &APIError{
		Type:   "/not-found",
		Title:  "Resource Not Found",
		Status: http.StatusNotFound,
		Detail: resource,
	}
}

func Error(w http.ResponseWriter, r *http.Request, apiError *APIError) {
	if apiError == nil {
		return
	}

	apiError.RequestID = middleware.GetReqID(r.Context())

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(apiError.Status)
	if err := json.NewEncoder(w).Encode(apiError); err != nil {
		return
	}
}
