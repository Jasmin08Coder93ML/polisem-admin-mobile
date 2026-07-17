package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"polisem-platform/internal/products"
	"polisem-platform/pkg/database" // Предполагаемый путь к вашему пакету БД
)

func main() {
	// 1. Инициализация базы данных (лучшая практика)
	db := database.Connect() 

	// 2. Инициализация слоев (DI)
	productRepo := products.NewRepository(db)
	productService := products.NewService(productRepo)
	productHandler := products.NewHandler(productService)

	// 3. Использование нового роутера (ServeMux)
	mux := http.NewServeMux()
	mux.HandleFunc("POST /products", productHandler.CreateProduct) // Современный синтаксис Go 1.22+

	// 4. Конфигурация сервера
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// 5. Запуск в горутине для корректного завершения
	go func() {
		log.Println("Polisem API server 8080 portunda işläp başlady...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Ожидание сигнала завершения
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	log.Println("Serveri ýapýaryn...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Shutdown(ctx)
}
