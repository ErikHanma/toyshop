package repositories

import (
	"catalog-service/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// ProductRepository struct for interacting with the products collection.
type ProductRepository struct {
	client *mongo.Client
}

// NewProductRepository creates a new ProductRepository.
func NewProductRepository(client *mongo.Client) *ProductRepository {
	return &ProductRepository{client: client}
}

// GetProducts retrieves all products from the database.
func (pr *ProductRepository) GetProducts() ([]models.Product, error) {
	collection := pr.client.Database("toyshop").Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
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

// ... other functions for product operations (create, update, delete) 
// CreateProduct добавляет новый продукт в базу данных.
func (pr *ProductRepository) CreateProduct(product *models.Product) error {
	collection := pr.client.Database("toyshop").Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, product)
	if err != nil {
		return err
	}
	return nil
}

// GetProductByID получает продукт из базы данных по его ID.
func (pr *ProductRepository) GetProductByID(id string) (*models.Product, error) {
	collection := pr.client.Database("toyshop").Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var product models.Product
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("product not found")
		}
		return nil, err
	}
	return &product, nil
}


// UpdateProduct обновляет существующий продукт в базе данных.
func (pr *ProductRepository) UpdateProduct(id string, product *models.Product) error {
	collection := pr.client.Database("toyshop").Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": product,
	}

	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return err
	}
	return nil
}


// DeleteProduct удаляет продукт из базы данных по его ID.
func (pr *ProductRepository) DeleteProduct(id string) error {
	collection := pr.client.Database("toyshop").Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}
