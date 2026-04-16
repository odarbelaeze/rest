package rest

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

// JSON writes the given data as a JSON response with the specified HTTP status code.
func JSON(ctx context.Context, w http.ResponseWriter, code int, data any) {
	w.Header().Set("Content-Type", "application/json")
	content, err := json.Marshal(data)
	if err != nil {
		slog.ErrorContext(ctx, "failed to marshal response data", "error", err)
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}
	w.WriteHeader(code)
	_, err = w.Write(content)
	if err != nil {
		slog.ErrorContext(ctx, "failed to write error response", "error", err)
	}
}
