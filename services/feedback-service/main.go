package main

import (
	// "fmt"
	"log"
	"os"
	"feedback-service/handlers"
	"feedback-service/repositories"
	"feedback-service/services"
	"feedback-service/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	// Загрузка переменных окружения из файла .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Подключение к базе данных MongoDB
	clientOptions := options.Client().ApplyURI(os.Getenv("mongodb://localhost:27017"))
	client, err := mongo.Connect(nil, clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	defer func() {
		if err := client.Disconnect(nil); err != nil {
			log.Fatal("Error disconnecting from MongoDB:", err)
		}
	}()

	// Инициализация репозитория и сервиса для работы с обратной связью
	feedbackRepository := repositories.NewFeedbackRepository(client)
	feedbackService := services.NewFeedbackService(feedbackRepository)

	// Инициализация Gin роутера
	router := gin.Default()

	// Инициализация маршрутов
	routes.SetupRoutes(router)

	// Инициализация сервиса обратной связи
	handlers.InitFeedbackService(feedbackService)

	// Запуск HTTP сервера
	log.Println("Feedback Service running on port 8084")
	if err := router.Run(":8084"); err != nil {
		log.Fatal("Failed to start the server:", err)
	}
}
