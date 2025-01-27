package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// CartItem представляет элемент в корзине.
type CartItem struct {
	ProductID primitive.ObjectID `bson:"product_id"` // ID товара
	Quantity  int                `bson:"quantity"`   // Количество товара
}

// Cart представляет корзину пользователя.
type Cart struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"` // ID корзины
	UserID   string             `bson:"user_id"`       // ID пользователя
	Items    []CartItem         `bson:"items"`         // Список товаров
	Total    float64            `bson:"total"`         // Общая сумма
}