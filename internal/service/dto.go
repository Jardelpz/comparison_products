package service

import "comparison/internal/models"

type CompareProductsResponse struct {
	RequestedIDs []string         `json:"requestedIds"`
	Found        []models.Product `json:"found"`
	NotFound     []string         `json:"notFound"`
	Comparison   CompareByFields  `json:"comparison"`
	Summary      CompareSummary   `json:"summary"`
}

type CompareSummary struct {
	Requested  int `json:"requested"`
	Found      int `json:"found"`
	NotFound   int `json:"notFound"`
	Duplicated int `json:"duplciated"`
}

type CompareByFields struct {
	Fields []string         `json:"fields"`
	Items  []map[string]any `json:"items"`
}

func NewCompareProductsResponse(requestedIDs []string, found []models.Product, notFound []string, duplicated []string, comparison CompareByFields) CompareProductsResponse {
	return CompareProductsResponse{
		RequestedIDs: requestedIDs,
		Found:        found,
		NotFound:     notFound,
		Comparison:   comparison,
		Summary: CompareSummary{
			Requested:  len(requestedIDs),
			Found:      len(found),
			NotFound:   len(notFound),
			Duplicated: len(duplicated),
		},
	}
}
