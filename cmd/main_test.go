package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"comparison/internal/handler"
	"comparison/internal/repository"
	"comparison/internal/service"

	"github.com/stretchr/testify/assert"
)

var repo = repository.NewProductRepository(nil)
var svc = service.NewProductService(repo)
var h = handler.NewProductHandler(svc)
var router = handler.NewRouter(h)

func TestPingRoute(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"message":"pong"}`, w.Body.String())
}

func TestRootRoute(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"message":"alive and kicking"}`, w.Body.String())
}

func TestTraceHeaderIsPropagated(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	req.Header.Set(handler.TraceIDHeader, "trace-abc")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "trace-abc", w.Header().Get(handler.TraceIDHeader))
}
