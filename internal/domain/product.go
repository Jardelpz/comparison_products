package domain

import (
	"comparison/internal/models"
)

// avoid circular import
type FindProductResult struct {
	Found      []models.Product
	NotFound   []string
	Duplicated []string
}
