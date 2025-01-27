package repositories

import (
	"catalog-service/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductRepository struct {
	collection *mongo.Collection
}

// NewProductRepository создаёт новый экземпляр ProductRepository.
func NewProductRepository(client *mongo.Client) *ProductRepository {
	return &ProductRepository{
		collection: client.Database("toyshop").Collection("products"),
	}
}

// GetProducts возвращает все продукты из коллекции.
func (pr *ProductRepository) GetProducts() ([]models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := pr.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []models.Product
	if err = cursor.All(ctx, &products); err != nil {
		return nil, err
	}

	return products, nil
}

// CreateProduct добавляет новый продукт.
func (pr *ProductRepository) CreateProduct(product *models.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := pr.collection.InsertOne(ctx, product)
	return err
}

// GetProductByID получает продукт по его ID.
func (pr *ProductRepository) GetProductByID(id primitive.ObjectID) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var product models.Product
	err := pr.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("product not found")
		}
		return nil, err
	}
	return &product, nil
}

// UpdateProduct обновляет продукт.
func (pr *ProductRepository) UpdateProduct(id primitive.ObjectID, product *models.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{"$set": product}
	_, err := pr.collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	return err
}

// DeleteProduct удаляет продукт по его ID.
func (pr *ProductRepository) DeleteProduct(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := pr.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
