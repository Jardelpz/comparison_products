package repository

import (
	"context"
	"testing"

	"comparison/internal/models"

	"github.com/stretchr/testify/require"
)

func TestNewProductRepository_IndexesByID(t *testing.T) {
	products := []models.Product{
		{ID: "1", Name: "A"},
		{ID: "2", Name: "B"},
	}

	repo := NewProductRepository(products)
	require.NotNil(t, repo)
	require.Len(t, repo.products, 2)
	require.Equal(t, "A", repo.products["1"].Name)
	require.Equal(t, "B", repo.products["2"].Name)
}

func TestFindProductsByIDs_SomeNotFound_TracksNotFoundInOrder(t *testing.T) {
	repo := NewProductRepository([]models.Product{
		{ID: "1", Name: "P1"},
		{ID: "3", Name: "P3"},
	})

	res, err := repo.FindProductsByIDs(context.Background(), []string{"1", "2", "3", "999"})
	require.NoError(t, err)

	require.Len(t, res.Found, 2)
	require.Equal(t, []string{"1", "3"}, []string{res.Found[0].ID, res.Found[1].ID})
	require.Equal(t, []string{"2", "999"}, res.NotFound)
}

func TestFindProductsByIDs_DuplicateIDs_TracksDuplicates(t *testing.T) {
	repo := NewProductRepository([]models.Product{
		{ID: "1", Name: "P1"},
	})

	res, err := repo.FindProductsByIDs(context.Background(), []string{"1", "1", "x"})
	require.NoError(t, err)

	// repo agora deduplica durante a busca
	require.Len(t, res.Found, 1)
	require.Equal(t, []string{"1"}, []string{res.Found[0].ID})
	require.Equal(t, []string{"x"}, res.NotFound)
	require.Equal(t, []string{"1"}, res.Duplicated)
}

func TestFindProductsByIDs_EmptyInput(t *testing.T) {
	repo := NewProductRepository([]models.Product{{ID: "1", Name: "P1"}})

	res, err := repo.FindProductsByIDs(context.Background(), []string{})
	require.NoError(t, err)
	require.Empty(t, res.Found)
	require.Empty(t, res.NotFound)
}
