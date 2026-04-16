package rest

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

// Error represents a structured error response for REST APIs.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// Implements the error interface.
func (e *Error) Error() string {
	return e.Message
}

type ErrorOption func(*Error)

// WithDetails allows adding additional details to the error response.
func WithDetails(details string) ErrorOption {
	return func(e *Error) {
		e.Details = details
	}
}

// NewError creates a new Error instance with the provided code, message, and optional configurations.
func NewError(code int, message string, opts ...ErrorOption) *Error {
	err := &Error{
		Code:    code,
		Message: message,
	}
	for _, opt := range opts {
		opt(err)
	}
	return err
}

// Err writes a structured error response to the http.ResponseWriter with the given status code and message.
func Err(ctx context.Context, w http.ResponseWriter, code int, message string, opts ...ErrorOption) {
	errorMessage := NewError(code, message, opts...)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	content, err := json.Marshal(errorMessage)
	if err != nil {
		slog.ErrorContext(ctx, "failed to marshal error response", "error", err)
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}
	_, err = w.Write(content)
	if err != nil {
		slog.ErrorContext(ctx, "failed to write error response", "error", err)
	}
}
