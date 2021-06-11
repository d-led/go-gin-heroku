package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type message struct {
	Message string `json:"message"`
}

func TestGreeting(t *testing.T) {
	response := get("/")
	assert.Equal(t, http.StatusOK, response.Result().StatusCode)
	var m message
	err := json.NewDecoder(response.Body).Decode(&m)
	assert.NoError(t, err)
	assert.Equal(t, "hi", m.Message)
}

func TestVersion(t *testing.T) {
	os.Unsetenv("SOURCE_VERSION")
	response := get("/version")
	assert.Equal(t, http.StatusOK, response.Result().StatusCode)
	assert.Equal(t, "", response.Body.String())

	os.Setenv("SOURCE_VERSION", "42")
	response = get("/version")
	assert.Equal(t, http.StatusOK, response.Result().StatusCode)
	assert.Equal(t, "42", response.Body.String())
}

func get(path string) *httptest.ResponseRecorder {
	request, _ := http.NewRequest(http.MethodGet, path, nil)
	response := httptest.NewRecorder()
	server().ServeHTTP(response, request)
	return response
}
