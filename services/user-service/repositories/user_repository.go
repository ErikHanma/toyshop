package repositories

import (
	"context"
	"fmt"
	"log"
	"time"
	"user-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository struct {
	collection *mongo.Collection
}

// NewUserRepository создаёт новый экземпляр UserRepository.
func NewUserRepository(client *mongo.Client) *UserRepository {
	return &UserRepository{
		collection: client.Database("toyshop").Collection("users"),
	}
}

// CreateUser добавляет нового пользователя с хешированным паролем.
func (ur *UserRepository) CreateUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	user.Password = string(hashedPassword)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = ur.collection.InsertOne(ctx, user)
	return err
}

// GetUsers возвращает всех пользователей.
func (ur *UserRepository) GetUsers() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := ur.collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []models.User
	if err = cursor.All(ctx, &users); err != nil {
		log.Printf("Error decoding users: %v", err)
		return nil, err
	}

	return users, nil
}

// GetUserByUsername возвращает пользователя по имени.
func (ur *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := ur.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}


// В user_repository.go добавим функцию GetUserByID
func (ur *UserRepository) GetUserByID(id string) (*models.User, error) {
	// Преобразуем строку ID в ObjectId
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err = ur.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
