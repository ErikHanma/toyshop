package repositories

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"user-service/models"
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

// Структура репозитория пользователей
type UserRepository struct {
	client *mongo.Client
}

// Конструктор репозитория
func NewUserRepository(client *mongo.Client) *UserRepository {
	return &UserRepository{client: client}
}

type User struct {
	ID       string `bson:"_id,omitempty"`
	Username string `bson:"username"`
	Email    string `bson:"email"`
}

// CreateUser создает нового пользователя с хешированным паролем.
func (ur *UserRepository) CreateUser(user *models.User) error {
    // Хеширование пароля
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return fmt.Errorf("failed to hash password: %w", err)
    }
    user.Password = string(hashedPassword)

    collection := ur.client.Database("toyshop").Collection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err = collection.InsertOne(ctx, user)
    return err
}

func (ur *UserRepository) GetUsers() ([]models.User, error) {
    collection := ur.client.Database("toyshop").Collection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        log.Printf("Error fetching users: %v", err)
        return nil, err
    }
    defer cursor.Close(ctx)

    var users []models.User
    for cursor.Next(ctx) {
        var user models.User
        if err := cursor.Decode(&user); err != nil {
            log.Printf("Error decoding user: %v", err)
            return nil, err
        }
        users = append(users, user)
    }
    return users, nil
}

// GetUserByUsername получает пользователя по имени пользователя.
func (ur *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	collection := ur.client.Database("toyshop").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

