package service

import (
	"comparison/pkg"
	"context"
	"log"

	"comparison/internal/models"
	"comparison/internal/trace"
)

const maxIDs = 10

type ProductService struct {
	repository ProductRepository
}

func NewProductService(repo ProductRepository) *ProductService {
	return &ProductService{
		repository: repo,
	}
}

func (ps *ProductService) CompareProducts(ctx context.Context, ids []string, fields []string) (CompareProductsResponse, error) {
	traceID, _ := trace.TraceIDFromContext(ctx)
	log.Printf("compare service: trace_id=%s ids=%v, fields=%v", traceID, ids, fields)

	if len(ids) < 2 {
		return CompareProductsResponse{}, ErrEnoughtProducts
	}

	if len(ids) > maxIDs {
		return CompareProductsResponse{}, ErrMaxProductComparison
	}

	res, err := ps.repository.FindProductsByIDs(ctx, ids)
	if err != nil {
		return CompareProductsResponse{}, err
	}
	if len(res.Found) == 0 {
		return CompareProductsResponse{}, ErrProductsNotFound
	}
	if len(res.Found) == 1 {
		return CompareProductsResponse{}, ErrEnoughtProductsNotFound
	}

	if len(fields) == 0 {
		fields = models.GetProductDefaultFields()
	}

	comparison := compareProductsByFields(res.Found, fields)
	return NewCompareProductsResponse(ids, res.Found, res.NotFound, res.Duplicated, CompareByFields{
		Fields: fields,
		Items:  comparison,
	}), nil
}

func compareProductsByFields(products []models.Product, fields []string) []map[string]any {
	items := make([]map[string]any, 0, len(products))

	for _, p := range products {
		item := map[string]any{
			"id": p.ID,
		}

		for _, field := range fields {
			switch field {
			case "id":
				continue
			case "name":
				item[field] = utils.EmptyToNil(p.Name)
			case "category":
				item[field] = utils.EmptyToNil(p.Category)
			case "description":
				item[field] = utils.EmptyToNil(p.Description)
			case "price":
				item[field] = utils.EmptyToNil(p.Price)
			case "size":
				item[field] = utils.EmptyToNil(p.Size)
			case "weight":
				item[field] = utils.EmptyToNil(p.Weight)
			case "color":
				item[field] = utils.EmptyToNil(p.Color)
			default:
				if p.Specs == nil {
					item[field] = nil
					continue
				}

				if v, ok := p.Specs[field]; ok {
					item[field] = v
				} else {
					item[field] = nil
				}
			}
		}
		items = append(items, item)

	}
	return items
}
