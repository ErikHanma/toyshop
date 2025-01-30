package services

import (
	"feedback-service/models"
	"feedback-service/repositories"
	"fmt"
	"log"
	"feedback-service/mailer"
)

// FeedbackService - сервис для работы с обратной связью
type FeedbackService struct {
	repo repositories.FeedbackRepository
}

// Новый сервис
func NewFeedbackService(repo repositories.FeedbackRepository) *FeedbackService {
	return &FeedbackService{repo: repo}
}

// Метод для добавления нового сообщения
func (service *FeedbackService) AddFeedback(feedback *models.Feedback) error {
	// Сохраняем сообщение в базе данных
	err := service.repo.Save(feedback)
	if err != nil {
		return fmt.Errorf("failed to save feedback: %v", err)
	}

	// Отправляем уведомление администратору
	err = mailer.SendEmail("admin@example.com", "New Feedback Received", fmt.Sprintf("You have a new feedback message: %s", feedback.Message))
	if err != nil {
		log.Printf("Failed to send email: %v", err)
	}

	return nil
}
