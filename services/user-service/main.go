package main

import (
	"context"
	database "database"
	"log"
	"net/http"

	"github.com/ErikHanma/toyshop/services/user-service/handlers"
	"github.com/ErikHanma/toyshop/services/user-service/repositories"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Подключение к MongoDB
	client, err := database.NewMongoClient("mongodb://localhost:27017")
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
	// Новый маршрут для поиска пользователя по ID
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetUserByIDHandler(w, r, userRepository) // Обработчик для поиска по ID
	})

	// Запуск сервера
	log.Println("User Service running on port 8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}
