package entity

import (
	"github.com/google/uuid"
)

const (
	ProductTableName = "product"
)

type product struct {
	Product_ID   uuid.UUID `gorm: "type:uuid;primary_key" json:"product_id"`
	Stock_P      int64     `gorm: "type:integer;not_null" json:"stok_p"`
	Product_type string    `gorm: "type:varchar(50);not_null" json:"product_type"`
	Price        int64     `gorm: "type:integer;not_null" json:"price"`
	Auditable
}

func NewProduct(product_id uuid.UUID, product_type string, stock_p, price int) *product {
	return &product{
		Product_ID:   product_id,
		Stock_P:      int64(stock_p),
		Product_type: product_type,
		Price:        int64(price),
		Auditable:    NewAuditable(),
	}
}

//Table specifies table name for Product
func (model *product) TableName() string {
	return ProductTableName
}
