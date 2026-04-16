package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Load decodes the JSON request body into a struct of type T and returns a pointer to it.
func Load[T any](r *http.Request) (*T, error) {
	var data T
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode request body: %w", err)
	}
	return &data, nil
}
