package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"comparison/internal/config"
	"comparison/internal/handler"
	"comparison/internal/models"
	"comparison/internal/repository"
	"comparison/internal/service"
)

func main() {
	_ = config.LoadDotEnv(".env")

	filePath := os.Getenv("FILE_PATH")
	if filePath == "" {
		log.Fatal("FILE_PATH is not set")
	}

	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Path not Found")
	}

	var productsLoaded []models.Product
	if err := json.Unmarshal(file, &productsLoaded); err != nil {
		log.Fatal("Error unmarshaling json")
	}

	repo := repository.NewProductRepository(productsLoaded)
	svc := service.NewProductService(repo)
	h := handler.NewProductHandler(svc)
	router := handler.NewRouter(h)

	srv := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("listening on :8080")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}

	router.Run()
}
