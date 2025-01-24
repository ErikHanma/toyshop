package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
	"user-service/models"
	"user-service/repositories"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// Обновленный обработчик получения пользователей
func GetUsersHandler(w http.ResponseWriter, r *http.Request, ur *repositories.UserRepository) { 
	users, err := ur.GetUsers() // Вызываем GetUsers у userRepository
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch users: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request, ur *repositories.UserRepository) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Валидация пользователя перед сохранением
	if err := user.Validate(); err != nil {
		http.Error(w, fmt.Sprintf("Ошибка валидации: %v", err), http.StatusBadRequest)
		return
	}

	// Хеширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	if err := ur.CreateUser(&user); err != nil {
		http.Error(w, fmt.Sprintf("Failed to create user: %v", err), http.StatusInternalServerError)
		return
	}

	// Генерация JWT
	token, err := generateJWT(user.Username)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}


// generateJWT генерирует JWT токен
func generateJWT(username string) (string, error) {
	// Секретный ключ для подписи токена (хранить в безопасном месте!)
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY")) 

	// Создание токена
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Время истечения - 24 часа
	})

	// Подпись токена секретным ключом
	return token.SignedString(secretKey)
}

// LoginHandler обрабатывает запросы аутентификации.
func LoginHandler(w http.ResponseWriter, r *http.Request, ur *repositories.UserRepository) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := ur.GetUserByUsername(credentials.Username) // Получаем пользователя по имени
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	if err := user.CheckPassword(credentials.Password); err != nil { // Проверяем пароль
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Генерация JWT (добавь свою логику)
	tokenString := "сгенерированный.jwt"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
