package repository

import (
	"comparison/internal/models"
	"comparison/internal/trace"
	"context"
	"log"
)

type ProductRepository struct {
	products map[string]models.Product
}

func NewProductRepository(products []models.Product) *ProductRepository {
	// prepare key, value for indexing found and not found products
	productsByID := make(map[string]models.Product, len(products)) // {1:produto, 2 produto... 10}
	for _, p := range products {
		productsByID[p.ID] = p
	}
	return &ProductRepository{products: productsByID}
}

func (r *ProductRepository) FindProductsByIDs(ctx context.Context, ids []string) (models.FindProductResult, error) {
	traceID, _ := trace.TraceIDFromContext(ctx)
	log.Printf("finding products by id: trace_id=%s ids=%v", traceID, ids)

	found := make([]models.Product, 0, len(ids))
	notFound := make([]string, 0)
	duplicated := make([]string, 0)
	// set in go
	seen := make(map[string]struct{})

	for _, id := range ids {
		if _, ok := seen[id]; ok {
			duplicated = append(duplicated, id)
			continue
		}
		seen[id] = struct{}{}
		p, ok := r.products[id]
		if ok {
			found = append(found, p)
			continue
		}
		notFound = append(notFound, id)
	}

	log.Printf("found=%v, notFound=%v, duplicated=%v, trace_id=%s", len(found), len(notFound), len(duplicated), traceID)
	return models.FindProductResult{
		Found:      found,
		NotFound:   notFound,
		Duplicated: duplicated,
	}, nil
}
