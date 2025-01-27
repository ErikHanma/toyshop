package db

import (
	"context"
	"fmt"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client



func ConnectMongoDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Укажите строку подключения
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to MongoDB: %v", err))
	}

	// Проверка соединения
	if err = client.Ping(ctx, nil); err != nil {
		panic(fmt.Sprintf("MongoDB ping failed: %v", err))
	}

	fmt.Println("Connected to MongoDB!")
	MongoClient = client
}
