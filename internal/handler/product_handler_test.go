package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"comparison/internal/models"
	"comparison/internal/repository"
	"comparison/internal/service"

	"github.com/stretchr/testify/require"
)

func TestCompareProductsHandler_MissingIDs_Returns400(t *testing.T) {
	repo := repository.NewProductRepository([]models.Product{{ID: "1", Name: "A"}, {ID: "2", Name: "B"}})
	svc := service.NewProductService(repo)
	h := NewProductHandler(svc)
	r := NewRouter(h)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/v1/compare/products", nil)
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
	require.Contains(t, w.Body.String(), service.ErrNoIds.Error())
}

func TestCompareProductsHandler_EmptyIDValue_Returns400(t *testing.T) {
	repo := repository.NewProductRepository([]models.Product{{ID: "1", Name: "A"}, {ID: "2", Name: "B"}})
	svc := service.NewProductService(repo)
	h := NewProductHandler(svc)
	r := NewRouter(h)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/v1/compare/products?ids=1,,2", nil)
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
	require.Contains(t, w.Body.String(), service.ErrEmptyIds.Error())
}

func TestCompareProductsHandler_Error_MaxItems_Returns400(t *testing.T) {
	repo := repository.NewProductRepository([]models.Product{{ID: "1", Name: "A"}, {ID: "2", Name: "B"}})
	svc := service.NewProductService(repo)
	h := NewProductHandler(svc)
	r := NewRouter(h)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/v1/compare/products?ids=1,2,3,4,5,6,7,8,9,10,11", nil)
	req.Header.Set(TraceIDHeader, "trace-test-123")
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
	require.Equal(t, "trace-test-123", w.Header().Get(TraceIDHeader))
	require.Contains(t, w.Body.String(), service.ErrMaxProductComparison.Error())
}

func TestCompareProductsHandler_Success_Returns200AndTraceHeader(t *testing.T) {
	repo := repository.NewProductRepository([]models.Product{{ID: "1", Name: "A"}, {ID: "2", Name: "B"}})
	svc := service.NewProductService(repo)
	h := NewProductHandler(svc)
	r := NewRouter(h)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/v1/compare/products?ids=1,2", nil)
	req.Header.Set(TraceIDHeader, "trace-test-123")
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "trace-test-123", w.Header().Get(TraceIDHeader))
	require.Contains(t, w.Body.String(), "requestedIds")
}
