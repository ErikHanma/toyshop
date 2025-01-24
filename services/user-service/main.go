package main

import (
	"context"
	"log"
	"net/http"
	"db"
	"user-service/handlers"
	"user-service/repositories"
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
	userRepository := repositories.NewUserRepository(client)

	// Роуты
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetUsersHandler(w, r, userRepository)
	})
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		handlers.RegisterHandler(w, r, userRepository) // Регистрация обработчика регистрации
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.LoginHandler(w, r, userRepository) // Регистрация обработчика аутентификации
	})

	// Запуск сервера
	log.Println("User Service running on port 8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}
