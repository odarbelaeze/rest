package rest_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/odarbelaeze/rest"
	"github.com/stretchr/testify/assert"
)

func TestJSON(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := context.Background()

	data := map[string]string{"message": "hello world"}
	rest.JSON(ctx, w, http.StatusOK, data)

	res := w.Result()
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, "application/json", res.Header.Get("Content-Type"))

	var resData map[string]string
	err := json.NewDecoder(res.Body).Decode(&resData)
	assert.NoError(t, err)
	assert.Equal(t, "hello world", resData["message"])
}

func TestJSON_Error(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := context.Background()

	// Use an unmarshable type to trigger the error path
	data := func() {} // Functions cannot be marshaled to JSON
	rest.JSON(ctx, w, http.StatusOK, data)

	res := w.Result()
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
}
