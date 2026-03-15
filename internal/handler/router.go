package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(urlHandler *ProductHandler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Use(TraceMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "alive and kicking",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/compare/products", urlHandler.CompareProductsHandler)
	}

	return r
}
