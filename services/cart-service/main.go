package main

import (
	"cart-service/handlers"
	repository "cart-service/repository"
	"context"
	"log"
	database "database"
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

	// Инициализация репозитория
	cartRepository := repository.NewCartRepository(client)

	// Роуты
	router := mux.NewRouter()
	router.HandleFunc("/carts", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetCartHandler(w, r, cartRepository)
	}).Methods(http.MethodGet)
	router.HandleFunc("/carts", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateCartHandler(w, r, cartRepository)
	}).Methods(http.MethodPost)
	router.HandleFunc("/carts/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateCartHandler(w, r, cartRepository)
	}).Methods(http.MethodPut)
	router.HandleFunc("/carts/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteCartHandler(w, r, cartRepository)
	}).Methods(http.MethodDelete)

	// Запуск сервиса
	log.Println("Cart Service running on port 8084")
	log.Fatal(http.ListenAndServe(":8084", router))
}
