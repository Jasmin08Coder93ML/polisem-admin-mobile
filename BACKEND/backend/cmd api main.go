package main

import (
	"log"
	"net/http"
	"os"
	"polisem-platform/pkg/database"
	"polisem-platform/pkg/middleware"
	"polisem-platform/internal/products"
)

func main() {
	// 1. DB birikmesi
	db, err := database.NewPostgresDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("DB error: %v", err)
	}
	defer db.Close()

	// 2. Modullary (Dependency Injection) gurmak
	productRepo := products.NewRepository(db)
	productService := products.NewService(productRepo)
	productHandler := products.NewHandler(productService)

	// 3. Router
	mux := http.NewServeMux()
	
	// Middleware bilen goralan API
	mux.Handle("/products", middleware.AuthMiddleware(http.HandlerFunc(productHandler.CreateProduct)))

	log.Println("Polisem API server 8080 portunda işledi.")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
