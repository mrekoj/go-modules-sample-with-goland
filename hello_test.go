package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloRoute(t *testing.T) {
	r := setupRouter()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/hello", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"message\":\"Hello, world.\"}", w.Body.String())
}

func TestConcurrencyRoute(t *testing.T) {
	r := setupRouter()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/concurrency", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Concurrency is not parallelism.", w.Body.String())
}
