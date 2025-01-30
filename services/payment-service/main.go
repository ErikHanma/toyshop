package main

import (
    "log"
    "net/http"
    "payment-service/handlers"
    // "github.com/gin-gonic/gin"
    "db"
    "payment-service/repositories"
    "context"
)

func main() {
    // Подключение к MongoDB
	client, err := db.NewMongoClient("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	// Инициализация репозитория с подключением
	paymentRepository := repositories.CreatePayment(client)

    http.HandleFunc("/payments", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreatePaymentHandler(w, r, paymentRepository)
	})

    http.HandleFunc("/payments/process", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreatePaymentHandler(w, r, paymentRepository)
	})


    log.Println("User Service running on port 8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}
