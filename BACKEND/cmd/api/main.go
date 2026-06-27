package main

import (
	"log"
	"net/http"
	"polisem-platform/internal/products"
)

func main() {
	// Repository (DB) we Service (Logika) döretmek
	productRepo := products.NewRepository()
	productService := products.NewService(productRepo)
	productHandler := products.NewHandler(productService)

	// Router sazlamalary
	http.HandleFunc("/products", productHandler.CreateProduct)

	log.Println("Polisem API server 8080 portunda işläp başlady...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
