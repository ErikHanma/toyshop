package models

// Payment - модель для хранения информации о платеже
type Payment struct {
    ID         string  `json:"id"`
    UserID     string  `json:"user_id"`
    OrderID    string  `json:"order_id"`
    Amount     float64 `json:"amount"`
    Status     string  `json:"status"`
    CreatedAt  string  `json:"created_at"`
}
