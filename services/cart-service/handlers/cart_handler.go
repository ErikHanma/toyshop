package handlers

import (
	"cart-service/models"
	"cart-service/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetCartHandler возвращает корзину по UserID.
func GetCartHandler(w http.ResponseWriter, r *http.Request, cr *repositories.CartRepository) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "Missing user_id parameter", http.StatusBadRequest)
		return
	}

	cart, err := cr.GetCartByUserID(userID)
	if err != nil {
		log.Printf("Failed to get cart: %v", err)
		http.Error(w, fmt.Sprintf("Failed to get cart: %v", err), http.StatusInternalServerError)
		return
	}

	if cart == nil {
		http.Error(w, "Cart not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cart); err != nil {
		log.Printf("Failed to encode cart: %v", err)
		http.Error(w, fmt.Sprintf("Failed to encode cart: %v", err), http.StatusInternalServerError)
	}
}

// CreateCartHandler создает новую корзину.
func CreateCartHandler(w http.ResponseWriter, r *http.Request, cr *repositories.CartRepository) {
	var cart models.Cart
	if err := json.NewDecoder(r.Body).Decode(&cart); err != nil {
		log.Printf("Failed to decode cart: %v", err)
		http.Error(w, fmt.Sprintf("Failed to decode cart: %v", err), http.StatusBadRequest)
		return
	}

	// Проверка user_id
	if cart.UserID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// Генерация ID для корзины
	cart.ID = primitive.NewObjectID()

	// Проверка существования корзины для данного пользователя
	existingCart, err := cr.GetCartByUserID(cart.UserID)
	if err != nil {
		log.Printf("Failed to check existing cart: %v", err)
		http.Error(w, fmt.Sprintf("Failed to check existing cart: %v", err), http.StatusInternalServerError)
		return
	}
	if existingCart != nil {
		http.Error(w, "Cart already exists for this user", http.StatusConflict)
		return
	}

	// Создание корзины
	if err := cr.CreateCart(&cart); err != nil {
		log.Printf("Failed to create cart: %v", err)
		http.Error(w, fmt.Sprintf("Failed to create cart: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cart); err != nil {
		log.Printf("Failed to encode cart: %v", err)
		http.Error(w, fmt.Sprintf("Failed to encode cart: %v", err), http.StatusInternalServerError)
	}
}

// UpdateCartHandler обновляет корзину.
func UpdateCartHandler(w http.ResponseWriter, r *http.Request, cr *repositories.CartRepository) {
	vars := mux.Vars(r)
	cartIDStr := vars["id"]

	cartID, err := primitive.ObjectIDFromHex(cartIDStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid ID format: %v", err), http.StatusBadRequest)
		return
	}

	var cart models.Cart
	if err := json.NewDecoder(r.Body).Decode(&cart); err != nil {
		log.Printf("Failed to decode cart: %v", err)
		http.Error(w, fmt.Sprintf("Failed to decode cart: %v", err), http.StatusBadRequest)
		return
	}

	if err := cr.UpdateCart(cartID, &cart); err != nil {
		log.Printf("Failed to update cart: %v", err)
		http.Error(w, fmt.Sprintf("Failed to update cart: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Cart updated successfully"))
}

// DeleteCartHandler удаляет корзину.
func DeleteCartHandler(w http.ResponseWriter, r *http.Request, cr *repositories.CartRepository) {
	vars := mux.Vars(r)
	cartIDStr := vars["id"]

	cartID, err := primitive.ObjectIDFromHex(cartIDStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid ID format: %v", err), http.StatusBadRequest)
		return
	}

	if err := cr.DeleteCart(cartID); err != nil {
		log.Printf("Failed to delete cart: %v", err)
		http.Error(w, fmt.Sprintf("Failed to delete cart: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Cart deleted successfully"))
}