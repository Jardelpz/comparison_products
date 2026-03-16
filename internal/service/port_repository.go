package service

import (
	"comparison/internal/models"
	"context"
)

type ProductRepository interface {
	FindProductsByIDs(ctx context.Context, ids []string) (models.FindProductResult, error)
}
