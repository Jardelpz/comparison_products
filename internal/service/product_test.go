package service

import (
	"context"
	"testing"

	"comparison/internal/models"
	"comparison/internal/repository"

	"github.com/stretchr/testify/require"
)

func TestProductService_CompareProducts_Errors_WhenNoneFound(t *testing.T) {
	repo := repository.NewProductRepository([]models.Product{{ID: "1", Name: "A"}})
	svc := NewProductService(repo)

	_, err := svc.CompareProducts(context.Background(), []string{"999"}, nil)
	require.ErrorIs(t, err, ErrEnoughtProducts)
}

func TestProductService_CompareProducts_Errors_WhenOnlyOneFound(t *testing.T) {
	repo := repository.NewProductRepository([]models.Product{{ID: "1", Name: "A"}})
	svc := NewProductService(repo)

	_, err := svc.CompareProducts(context.Background(), []string{"1", "999"}, nil)
	require.ErrorIs(t, err, ErrEnoughtProductsNotFound)
}

func TestProductService_CompareProducts_BuildsRowsWithNulls(t *testing.T) {
	repo := repository.NewProductRepository([]models.Product{
		{ID: "1", Name: "A", Specs: map[string]any{"battery": "5000mAh"}},
		{ID: "2", Name: "B", Specs: map[string]any{"camera": "48MP"}},
	})
	svc := NewProductService(repo)

	resp, err := svc.CompareProducts(context.Background(), []string{"1", "2"}, nil)
	require.NoError(t, err)
	require.Equal(t, []string{"1", "2"}, resp.RequestedIDs)
	require.Len(t, resp.Found, 2)
	require.Empty(t, resp.NotFound)
}

func TestProductService_CompareProducts_BuildsRowsWithNotFound(t *testing.T) {
	repo := repository.NewProductRepository([]models.Product{
		{ID: "1", Name: "A", Specs: map[string]any{"battery": "5000mAh"}},
		{ID: "2", Name: "B", Specs: map[string]any{"camera": "48MP"}},
		{ID: "3", Name: "C", Specs: map[string]any{"screen": "10px"}},
	})
	svc := NewProductService(repo)

	resp, err := svc.CompareProducts(context.Background(), []string{"1", "2", "44"}, nil)
	require.NoError(t, err)
	require.Len(t, resp.Found, 2)
	require.Len(t, resp.NotFound, 1)
}
