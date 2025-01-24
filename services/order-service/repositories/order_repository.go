package repositories

import (
	"context"
	"fmt"
	// "time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"catalog-service/models"
	// "order-service/models"
)


// OrderRepository структура для взаимодействия с коллекцией заказов.
type OrderRepository struct {
	client *mongo.Client
}

// NewOrderRepository создает новый OrderRepository.
func NewOrderRepository(client *mongo.Client) *OrderRepository {
	return &OrderRepository{client: client}
}

// GetOrders получает все заказы.
func (or *OrderRepository) GetOrders(ctx context.Context) ([]models.Order, error) {
	collection := or.client.Database("toyshop").Collection("orders")

	cursor, err := collection.Find(ctx, bson.M{})
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

// GetOrderByID получает заказ по ID.
func (or *OrderRepository) GetOrderByID(ctx context.Context, orderID string) (*models.Order, error) {
	collection := or.client.Database("toyshop").Collection("orders")

	objID, err := primitive.ObjectIDFromHex(orderID) 
	if err != nil {
		return nil, err
	}

	var order models.Order
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&order)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("order not found")
		}
		return nil, err
	}
	return &order, nil
}

// CreateOrder создает новый заказ.
func (or *OrderRepository) CreateOrder(ctx context.Context, order *models.Order) error {
	collection := or.client.Database("toyshop").Collection("orders")

	result, err := collection.InsertOne(ctx, order)
	if err != nil {
		return err
	}

	// Обновляем ID заказа сгенерированным MongoDB ID
	order.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

// UpdateOrder обновляет существующий заказ.
func (or *OrderRepository) UpdateOrder(ctx context.Context, orderID string, order *models.Order) error {
	collection := or.client.Database("toyshop").Collection("orders")

	objID, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": order,
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return err
	}
	return nil
}

// DeleteOrder удаляет заказ по ID.
func (or *OrderRepository) DeleteOrder(ctx context.Context, orderID string) error {
	collection := or.client.Database("toyshop").Collection("orders")

	objID, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}
	return nil
}
