package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"context"
	"time"
	"github.com/gorilla/mux"
	"order-service/models"
	"order-service/repositories"
)

// OrderHandler структура для обработчиков заказов.
type OrderHandler struct {
	orderRepository *repositories.OrderRepository
	// ... другие зависимости, например,
	// клиенты для взаимодействия с другими сервисами
}

// NewOrderHandler создает новый OrderHandler.
func NewOrderHandler(orderRepository *repositories.OrderRepository) *OrderHandler {
	return &OrderHandler{
		orderRepository: orderRepository,
		// ... инициализация других зависимостей
	}
}

// GetOrdersHandler обрабатывает HTTP-запросы GET на /orders для получения всех заказов.
func (oh *OrderHandler) GetOrdersHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	orders, err := oh.orderRepository.GetOrders(ctx)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get orders: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(orders); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode orders: %v", err), http.StatusInternalServerError)
	}
}

// GetOrderByIDHandler обрабатывает HTTP-запросы GET на /orders/{id}.
func (oh *OrderHandler) GetOrderByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["id"]
	ctx := r.Context();

	order, err := oh.orderRepository.GetOrderByID(ctx, orderID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get order: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(order); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode order: %v", err), http.StatusInternalServerError)
	}
}

func checkProductExists(productID string) (bool, error) {
	catalogServiceURL := fmt.Sprintf("http://localhost:8081/products/%s", productID) // Замени на адрес catalog-service
	resp, err := http.Get(catalogServiceURL)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK, nil
}

// CreateOrderHandler обрабатывает HTTP-запросы POST на /orders для создания нового заказа.
func (oh *OrderHandler) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode order: %v", err), http.StatusBadRequest)
		return
	}

	for _, item := range order.Items {
	// Проверяем, существует ли товар с таким ID в catalog-service
	productExists, err := checkProductExists(item.ProductID) 
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to check product existence: %v", err), http.StatusInternalServerError)
		return
	}

		if !productExists {
			http.Error(w, fmt.Sprintf("Product with ID %s not found", item.ProductID), http.StatusBadRequest)
			return
		}
	}

	// Создаем контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Передаем контекст в CreateOrder
	if err := oh.orderRepository.CreateOrder(ctx, &order); err != nil {
		http.Error(w, fmt.Sprintf("Failed to create order: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Order created successfully"))
}

// UpdateOrderHandler обрабатывает HTTP-запросы PUT на /orders/{id} для обновления заказа.
func (oh *OrderHandler) UpdateOrderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["id"]
	ctx := r.Context()

	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode order: %v", err), http.StatusBadRequest)
		return
	}

	if err := oh.orderRepository.UpdateOrder(ctx, orderID, &order); err != nil {
		http.Error(w, fmt.Sprintf("Failed to update order: %v", err), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Order updated successfully"))
}

// DeleteOrderHandler обрабатывает HTTP-запросы DELETE на /orders/{id} для удаления заказа.
func (oh *OrderHandler) DeleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["id"]
	ctx := r.Context()

	if err := oh.orderRepository.DeleteOrder(ctx, orderID); err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete order: %v", err), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Order deleted successfully"))
}

