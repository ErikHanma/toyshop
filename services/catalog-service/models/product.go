package models

// Product represents a product in the toyshop.
type Product struct {
	ID          string   `bson:"_id,omitempty"`
	Name        string   `bson:"name"`
	Description string   `bson:"description"`
	Price       float64  `bson:"price"`
	Categories  []string `bson:"categories"`
}
