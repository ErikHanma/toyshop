package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"order-service/handlers"
	"order-service/repositories"
)

func main() {
	// Подключение к MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	// Инициализация репозитория заказов
	orderRepository := repositories.NewOrderRepository(client)

	orderHandler := handlers.NewOrderHandler(orderRepository)

	// Создание роутера
	router := mux.NewRouter()

	// Маршруты для order-service
	router.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		orderHandler.GetOrdersHandler(w, r)
	}).Methods(http.MethodGet)

	// Запуск сервера
	fmt.Println("Order Service listening on port 8082") // Выбери свободный порт
	log.Fatal(http.ListenAndServe(":8082", router))
}
