package main

import (
	"cart-service/handlers"
	"cart-service/repository"
	"context"
	"log"
	"db"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	// Подключение к MongoDB
	db.ConnectMongoDB()

	client := db.MongoClient

	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	// Инициализация репозитория
	cartRepository := repositories.NewCartRepository(client)

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