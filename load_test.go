package rest_test

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/odarbelaeze/rest"
	"github.com/stretchr/testify/assert"
)

type testData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestLoad(t *testing.T) {
	jsonBody := `{"name": "John Doe", "age": 30}`
	r, err := http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(jsonBody))
	assert.NoError(t, err)

	data, err := rest.Load[testData](r)
	assert.NoError(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, "John Doe", data.Name)
	assert.Equal(t, 30, data.Age)
}

func TestLoad_Error(t *testing.T) {
	jsonBody := `{"name": "John Doe", "age": "thirty"}` // invalid age type
	r, err := http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(jsonBody))
	assert.NoError(t, err)

	_, err = rest.Load[testData](r)
	assert.Error(t, err)
}
