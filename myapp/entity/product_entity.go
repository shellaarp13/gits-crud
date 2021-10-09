package entity

import (
	"github.com/google/uuid"
)

const (
	ProductTableName = "Product"
)

type product struct {
	Product_ID   uuid.UUID `gorm: "type:uuid;primary_key" json:"customer_id"`
	Stock_P      int64     `gorm: "type:integer;not_null" json:"Stok_P"`
	Product_type string    `gorm: "type:varchar(50);not_null" json:"Product_type"`
	Price        int64     `gorm: "type:integer;not_null" json:"Price"`
}

func NewProduct(Product_ID uuid.UUID, Product_type string, Stock_P, Price int) *product {
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
