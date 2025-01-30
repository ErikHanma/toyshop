package handlers

import (
    "encoding/json"
    "fmt"
    "net/http"
    "payment-service/repositories"
)

// ProcessPaymentHandler обрабатывает платёж через API
func ProcessPaymentHandler(w http.ResponseWriter, r *http.Request) {
    var request struct {
        PaymentID string `json:"payment_id"`
    }

    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    payment, err := repositories.GetPaymentByID(request.PaymentID)
    if err != nil {
        http.Error(w, fmt.Sprintf("Payment not found: %v", err), http.StatusNotFound)
        return
    }

    // Вызов User Service через HTTP API
    user, err := getUserFromUserService(payment.UserID)
    if err != nil || user.Balance < payment.Amount {
        http.Error(w, "Insufficient balance", http.StatusBadRequest)
        return
    }

    // Обновление через API User Service
    err = updateUserBalanceInUserService(payment.UserID, -payment.Amount)
    if err != nil {
        http.Error(w, "Failed to update balance", http.StatusInternalServerError)
        return
    }

    err = repositories.UpdatePaymentStatus(payment.ID, "paid")
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to update status: %v", err), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(payment)
}