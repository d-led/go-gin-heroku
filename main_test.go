package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type message struct {
	Message string `json:"message"`
}

func TestGreeting(t *testing.T) {
	assert.Equal(t, 4, 2+2)
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()
	server().ServeHTTP(response, request)
	var m message
	err := json.NewDecoder(response.Body).Decode(&m)
	assert.NoError(t, err)
	assert.Equal(t, "hi", m.Message)
}
