package repositories

import (
    "payment-service/models"
    "time"
    "fmt"
)

var payments = make(map[string]models.Payment)

// CreatePayment создает новый платёж
func CreatePayment(userID, orderID string, amount float64) (models.Payment, error) {
    paymentID := fmt.Sprintf("payment-%d", len(payments)+1)
    payment := models.Payment{
        ID:        paymentID,
        UserID:    userID,
        OrderID:   orderID,
        Amount:    amount,
        Status:    "pending",
        CreatedAt: time.Now().Format(time.RFC3339),
    }
    payments[paymentID] = payment
    return payment, nil
}

// GetPaymentByID возвращает информацию о платеже по ID
func GetPaymentByID(paymentID string) (models.Payment, error) {
    payment, exists := payments[paymentID]
    if !exists {
        return models.Payment{}, fmt.Errorf("Payment not found")
    }
    return payment, nil
}

// UpdatePaymentStatus обновляет статус платежа
func UpdatePaymentStatus(paymentID, status string) error {
    payment, exists := payments[paymentID]
    if !exists {
        return fmt.Errorf("Payment not found")
    }
    payment.Status = status
    payments[paymentID] = payment
    return nil
}
