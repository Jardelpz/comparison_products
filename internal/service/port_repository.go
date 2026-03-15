package service

import (
	"comparison/internal/domain"
	"context"
)

type ProductRepository interface {
	FindProductsByIDs(ctx context.Context, ids []string) (domain.FindProductResult, error)
}
