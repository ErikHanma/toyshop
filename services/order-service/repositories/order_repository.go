package repositories

import (
	"context"
	"fmt"
	"order-service/models"
	// "time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	collection *mongo.Collection
}

// NewOrderRepository создаёт новый экземпляр OrderRepository.
func NewOrderRepository(client *mongo.Client) *OrderRepository {
	return &OrderRepository{
		collection: client.Database("toyshop").Collection("orders"),
	}
}

// GetOrders возвращает все заказы.
func (or *OrderRepository) GetOrders(ctx context.Context) ([]models.Order, error) {
	cursor, err := or.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []models.Order
	if err = cursor.All(ctx, &orders); err != nil {
		return nil, err
	}

	return orders, nil
}

// GetOrderByID возвращает заказ по его ID.
func (or *OrderRepository) GetOrderByID(ctx context.Context, orderID string) (*models.Order, error) {
	objID, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		return nil, err
	}

	var order models.Order
	err = or.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&order)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("order not found")
		}
		return nil, err
	}
	return &order, nil
}

// CreateOrder добавляет новый заказ.
func (or *OrderRepository) CreateOrder(ctx context.Context, order *models.Order) error {
	result, err := or.collection.InsertOne(ctx, order)
	if err != nil {
		return err
	}
	order.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

// UpdateOrder обновляет заказ.
func (or *OrderRepository) UpdateOrder(ctx context.Context, orderID string, order *models.Order) error {
	objID, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		return err
	}

	update := bson.M{"$set": order}
	_, err = or.collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	return err
}

// DeleteOrder удаляет заказ.
func (or *OrderRepository) DeleteOrder(ctx context.Context, orderID string) error {
	objID, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		return err
	}

	_, err = or.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
