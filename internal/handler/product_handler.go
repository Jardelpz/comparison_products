package handler

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"comparison/internal/service"
	"comparison/internal/trace"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service ProductService
}

func NewProductHandler(repo ProductService) *ProductHandler {
	return &ProductHandler{
		service: repo,
	}
}

// - /v1/compare/products?ids=1,2,3&fields=price,rating,brand
func (ph *ProductHandler) CompareProductsHandler(c *gin.Context) {
	traceID, _ := trace.TraceIDFromContext(c.Request.Context())
	log.Printf("compare handler: trace_id=%s", traceID)

	ids, fields, err := GetFromRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()
	result, err := ph.service.CompareProducts(ctx, ids, fields)
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			c.JSON(http.StatusGatewayTimeout, gin.H{"error": "timeout"})
			return
		}
		if errors.Is(ctx.Err(), context.Canceled) {
			c.JSON(http.StatusRequestTimeout, gin.H{"error": "canceled"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func GetFromRequest(c *gin.Context) ([]string, []string, error) {
	rawIds := strings.TrimSpace(c.Query("ids"))
	if rawIds == "" {
		return nil, nil, service.ErrNoIds
	}

	parts := strings.Split(rawIds, ",")
	ids := make([]string, 0, len(parts))

	for _, p := range parts {
		v := strings.TrimSpace(p)
		if v == "" {
			return nil, nil, service.ErrEmptyIds
		}
		ids = append(ids, v)
	}

	rawFields := strings.TrimSpace(c.Query("fields"))
	fields := make([]string, 0)
	if rawFields != "" {
		for _, f := range strings.Split(rawFields, ",") {
			fv := strings.TrimSpace(f)
			if fv == "" {
				return nil, nil, service.ErrEmptyFields
			}
			fields = append(fields, fv)
		}
	}
	return ids, fields, nil
}
