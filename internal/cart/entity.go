package cart

import "time"

type CartItem struct {
	ID          int       `json:"id" gorm:"autoIncrement,primaryKey"`
	ProductSku  string    `json:"productSku"`
	UserId      string    `json:"userId"`
	Quantity    int       `json:"quantity"`
	TimeAdded   time.Time `json:"timeAdded"`
	LastUpdated time.Time `json:"lastUpdated"`
}
