package handler

import (
	"comparison/internal/service"
	"context"
)

type ProductService interface {
	CompareProducts(ctx context.Context, ids []string, fields []string) (service.CompareProductsResponse, error)
}
