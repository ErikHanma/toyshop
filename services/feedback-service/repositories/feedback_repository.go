package repositories

import (
	"context"
	// "user-service/models"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	"feedback-service/models"
	"time"
)

type FeedbackRepository struct {
	collection *mongo.Collection
}

func NewFeedbackRepository(client *mongo.Client) *FeedbackRepository {
	return &FeedbackRepository{
		collection: client.Database("toyshop").Collection("feedbacks"),
	}
}

// Сохранение сообщения в базе данных
func (repo *FeedbackRepository) Save(feedback *models.Feedback) error {
	feedback.CreatedAt = time.Now()
	_, err := repo.collection.InsertOne(context.Background(), feedback)
	return err
}

// Получение всех сообщений (или фильтрация)
func (repo *FeedbackRepository) GetAll() ([]models.Feedback, error) {
	cursor, err := repo.collection.Find(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var feedbacks []models.Feedback
	if err = cursor.All(context.Background(), &feedbacks); err != nil {
		return nil, err
	}

	return feedbacks, nil
}
