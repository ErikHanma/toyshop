package database

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

// NewMongoClient создает новый клиент MongoDB и возвращает его.
func NewMongoClient(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Проверка соединения
	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("MongoDB ping failed: %w", err)
	}

	fmt.Println("Connected to MongoDB!")
	return client, nil
}
