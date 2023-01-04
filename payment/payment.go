package payment

import (
	"go-project/product"
	"time"
)

type Payment struct {
	ID        int `json:"id"`
	ProductID int
	Product   product.Product `json:"product_id" gorm:"foreignKey:ProductID;references:ID ;constraint:OnDelete:CASCADE`
	PricePaid float64         `json:"price_paid"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
