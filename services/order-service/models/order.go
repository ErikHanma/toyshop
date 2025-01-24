package models

import "time"

// Order представляет собой заказ.
type Order struct {
	ID            string       `bson:"_id,omitempty" json:"id"`
	UserID        string       `bson:"user_id" json:"user_id"`
	Items         []OrderItem  `bson:"items" json:"items"`
	ShippingAddress Address     `bson:"shipping_address" json:"shipping_address"`
	PaymentInfo    PaymentInfo `bson:"payment_info" json:"payment_info"`
	Status        string       `bson:"status" json:"status"` 
	CreatedAt     time.Time    `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time    `bson:"updated_at" json:"updated_at"`
}

// OrderItem представляет собой позицию в заказе.
type OrderItem struct {
	ProductID   string  `bson:"product_id" json:"product_id"`
	Name        string  `bson:"name" json:"name"` 
	Price       float64 `bson:"price" json:"price"`
	Quantity    int     `bson:"quantity" json:"quantity"`
}

// Address представляет собой адрес доставки.
type Address struct {
	Street     string `bson:"street" json:"street"`
	City       string `bson:"city" json:"city"`
	State      string `bson:"state" json:"state"`
	PostalCode string `bson:"postal_code" json:"postal_code"` 
}

// PaymentInfo  представляет собой информацию об оплате.
type PaymentInfo struct {
	// ... поля для хранения информации об оплате, 
	// например, ID транзакции, способ оплаты 
}
