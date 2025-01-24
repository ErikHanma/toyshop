package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"order-service/handlers"  // Замени на путь к твоим обработчикам
	"order-service/repositories" // Замени на путь к твоим репозиториям
)

func main() {
	// Подключение к MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") // Замени на свой URL
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

	// Создание роутера
	router := mux.NewRouter()

	// Маршруты для order-service
	router.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetOrdersHandler(w, r, orderRepository)
	}).Methods(http.MethodGet)
	// ... (добавь другие маршруты для создания, получения по ID, обновления и удаления заказов)

	// Запуск сервера
	fmt.Println("Order Service listening on port 8083") // Выбери свободный порт
	log.Fatal(http.ListenAndServe(":8083", router))
}
