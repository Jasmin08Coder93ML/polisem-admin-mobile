package warehouse

import "time"

// StockItem представляет единицу товара на складе
type StockItem struct {
	ID        string    `json:"id"`
	ProductID string    `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Warehouse string    `json:"warehouse_name"`
	UpdatedAt time.Time `json:"updated_at"`
}
