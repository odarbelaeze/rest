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

func TestError_Error(t *testing.T) {
	err := &rest.Error{Message: "test error"}
	assert.Equal(t, "test error", err.Error())
}

func TestNewError(t *testing.T) {
	err := rest.NewError(http.StatusBadRequest, rest.WithDetails("some details"))
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest, err.Code)
	assert.Equal(t, "bad request", err.Message)
	assert.Equal(t, "some details", err.Details)
}

func TestErr(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := context.Background()

	rest.Err(ctx, w, http.StatusNotFound, rest.WithDetails("could not find resource"))

	res := w.Result()
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
	assert.Equal(t, "application/json", res.Header.Get("Content-Type"))

	var errResp rest.Error
	err := json.NewDecoder(res.Body).Decode(&errResp)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusNotFound, errResp.Code)
	assert.Equal(t, "not found", errResp.Message)
	assert.Equal(t, "could not find resource", errResp.Details)
}
