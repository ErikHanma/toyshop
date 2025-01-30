package handlers

import (
	"feedback-service/models"
	"feedback-service/services"
	"encoding/json"
	"net/http"
	"fmt"
)

var feedbackService *services.FeedbackService

// Инициализация сервиса обратной связи
func InitFeedbackService(service *services.FeedbackService) {
	feedbackService = service
}

// Обработчик для отправки сообщения
func SendFeedbackHandler(w http.ResponseWriter, r *http.Request) {
	var feedback models.Feedback
	// Декодируем тело запроса
	err := json.NewDecoder(r.Body).Decode(&feedback)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Валидация данных
	if feedback.Message == "" || feedback.Email == "" {
		http.Error(w, "Email and message are required", http.StatusBadRequest)
		return
	}

	// Добавляем сообщение через сервис
	err = feedbackService.AddFeedback(&feedback)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to save feedback: %v", err), http.StatusInternalServerError)
		return
	}

	// Ответ пользователю
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Feedback submitted successfully"))
}
