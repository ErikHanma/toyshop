package main

import (
	"catalog-service/handlers"
	"catalog-service/repositories"
	"context"
	database "database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Подключение к MongoDB
	database.ConnectMongoDB()

	client := database.MongoClient

	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	// Инициализация репозитория с подключением
	productRepository := repositories.NewProductRepository(client)

	// Роуты
	router := mux.NewRouter()
	router.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetProductsHandler(w, r, productRepository)
	}).Methods(http.MethodGet)
	router.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateProductHandler(w, r, productRepository)
	}).Methods(http.MethodPost)
	router.HandleFunc("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetProductByIDHandler(w, r, productRepository)
	}).Methods(http.MethodGet)
	router.HandleFunc("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateProductHandler(w, r, productRepository)
	}).Methods(http.MethodPut)
	router.HandleFunc("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteProductHandler(w, r, productRepository)
	}).Methods(http.MethodDelete)

	// Запуск сервера
	log.Println("Catalog Service running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
