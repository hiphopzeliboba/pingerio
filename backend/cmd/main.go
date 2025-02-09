package main

import (
	"context"
	"log"
	"net/http"
	"pingerio/backend/internal/api/handler"
	"pingerio/backend/internal/api/router"
	"pingerio/backend/internal/db"
	"pingerio/backend/internal/repository/containers"

	service "pingerio/backend/internal/service/containers"
)

func main() {
	ctx := context.Background()
	pg_pool, err := db.NewPostgresPool(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to pg pool: %v", err)
	}
	containerRepo := containers.NewContainerRepository(pg_pool)

	containerRepo.CreateTable(ctx)

	containerService := service.NewContainerService(containerRepo)
	containerHandler := handler.NewContainerHandler(containerService)

	r := router.NewRouter(containerHandler)

	log.Printf("Starting server on :8081")
	if err := http.ListenAndServe(":8081", r.Setup()); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
