package entity

import (
	"github.com/google/uuid"
)

const (
	ProductTableName = "product"
)

type Product struct {
	Product_ID   uuid.UUID `gorm:"type:uuid;primary_key" json:"product_id"`
	Stock_P      int32     `gorm:"type:varchar(50);not_null" json:"stok_p"`
	Product_type string    `gorm:"type:varchar(50);not_null" json:"product_type"`
	Price        int32     `gorm:"type:varchar(50);not_null" json:"price"`
	Auditable
}

func NewProduct(product_id uuid.UUID, stock_p int32, product_type string, price int32) *Product {
	return &Product{
		Product_ID:   product_id,
		Stock_P:      stock_p,
		Product_type: product_type,
		Price:        price,
		Auditable:    NewAuditable(),
	}
}

//Table specifies table name for Product
func (model *Product) TableName() string {
	return ProductTableName
}
