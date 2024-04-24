package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/leo-yssa/gin-template/backend/routers"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := routers.NewRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestAdminRoute(t *testing.T) {
	router := routers.NewRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/admin", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "admin", w.Body.String())
}