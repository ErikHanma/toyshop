package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"catalog-service/models"
	"catalog-service/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetProductsHandler retrieves all products.
func GetProductsHandler(w http.ResponseWriter, r *http.Request, pr *repositories.ProductRepository) {
	id := r.URL.Query().Get("id") // Получаем параметр id из URL
	if id != "" {
		// Если передан ID, возвращаем конкретный продукт
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid ID format: %v", err), http.StatusBadRequest)
			return
		}

		product, err := pr.GetProductByID(objectID)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to get product: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(product); err != nil {
			http.Error(w, fmt.Sprintf("Failed to encode product: %v", err), http.StatusInternalServerError)
		}
		return
	}

	// Если ID не передан, возвращаем все продукты
	products, err := pr.GetProducts()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get products: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode products: %v", err), http.StatusInternalServerError)
	}
}


// CreateProductHandler обрабатывает HTTP-запросы POST на /products для создания продукта.
func CreateProductHandler(w http.ResponseWriter, r *http.Request, pr *repositories.ProductRepository) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode product: %v", err), http.StatusBadRequest)
		return
	}

	if err := pr.CreateProduct(&product); err != nil {
		http.Error(w, fmt.Sprintf("Failed to create product: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Product created successfully"))
}

// GetProductByIDHandler обрабатывает HTTP-запросы GET на /products/{id} для получения продукта по ID.
func GetProductByIDHandler(w http.ResponseWriter, r *http.Request, pr *repositories.ProductRepository) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Получить ID продукта из URL
	vars := mux.Vars(r)
	productIDStr := vars["id"]

	// Преобразуем строковый ID в ObjectId
	productID, err := primitive.ObjectIDFromHex(productIDStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid ID format: %v", err), http.StatusBadRequest)
		return
	}

	product, err := pr.GetProductByID(productID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get product: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(product); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode product: %v", err), http.StatusInternalServerError)
	}
}

// UpdateProductHandler обрабатывает HTTP-запросы PUT на /products/{id} для обновления продукта.
func UpdateProductHandler(w http.ResponseWriter, r *http.Request, pr *repositories.ProductRepository) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Получить ID продукта из URL
	vars := mux.Vars(r)
	productIDStr := vars["id"]

	// Преобразуем строковый ID в ObjectId
	productID, err := primitive.ObjectIDFromHex(productIDStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid ID format: %v", err), http.StatusBadRequest)
		return
	}

	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode product: %v", err), http.StatusBadRequest)
		return
	}

	if err := pr.UpdateProduct(productID, &product); err != nil {
		http.Error(w, fmt.Sprintf("Failed to update product: %v", err), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Product updated successfully"))
}

// DeleteProductHandler обрабатывает HTTP-запросы DELETE на /products/{id} для удаления продукта.
func DeleteProductHandler(w http.ResponseWriter, r *http.Request, pr *repositories.ProductRepository) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Получить ID продукта из URL
	vars := mux.Vars(r)
	productIDStr := vars["id"]

	// Преобразуем строковый ID в ObjectId
	productID, err := primitive.ObjectIDFromHex(productIDStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid ID format: %v", err), http.StatusBadRequest)
		return
	}

	if err := pr.DeleteProduct(productID); err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete product: %v", err), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Product deleted successfully"))
}
