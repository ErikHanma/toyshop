package repositories

import (
	"cart-service/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CartRepository struct {
	collection *mongo.Collection
}

// NewCartRepository создает новый экземпляр CartRepository.
func NewCartRepository(client *mongo.Client) *CartRepository {
	return &CartRepository{
		collection: client.Database("toyshop").Collection("carts"),
	}
}

// GetCartByUserID возвращает корзину по UserID.
func (cr *CartRepository) GetCartByUserID(userID string) (*models.Cart, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var cart models.Cart
	err := cr.collection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&cart)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &cart, nil
}

// CreateCart создает новую корзину.
func (cr *CartRepository) CreateCart(cart *models.Cart) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := cr.collection.InsertOne(ctx, cart)
	return err
}

// UpdateCart обновляет существующую корзину.
func (cr *CartRepository) UpdateCart(cartID primitive.ObjectID, updatedCart *models.Cart) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{"$set": updatedCart}
	_, err := cr.collection.UpdateOne(ctx, bson.M{"_id": cartID}, update)
	return err
}

// DeleteCart удаляет корзину по ее ID.
func (cr *CartRepository) DeleteCart(cartID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := cr.collection.DeleteOne(ctx, bson.M{"_id": cartID})
	return err
}