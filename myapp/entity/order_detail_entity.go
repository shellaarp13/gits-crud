package entity

import (
	"github.com/google/uuid"
)

const (
	OderDetailsTableName = "order_details"
)

type order_details struct {
	Order_details_id  	uuid.UUID `gorm:type:uuid;primary_key" json:"order_details"`
	Order_number      	int64     `gorm:type:integer;not_null" json:"order_number"`
	Product_id       	uuid.UUID `gorm:type:uuid;not_null" json:"product_id"`
	Quantity_product 	int64     `gorm:type:integer;not_null" json:"quantity_product"`
	Order            	*Order    `gorm:"foreignKey:Order_number" json:"order_number`
	Product          	*Product  `gorm:"foreignKey:Product_id" json:"product_id"`
	Auditable
}

func NewOrderDetails(order_details_id, product_id uuid.UUID, order_number, quantity_product int) *OrderDetails {
	return &OrderDetails{
		Order_details_id: order_details_id,
		Order_number:     order_number,
		Product_id:       product_id,
		Quantity_product: quantity_product,
		Auditable:        NewAuditable(),
	}
}

func (model *OrderDetails) TableName() string {
	return OrderDetailsTableName
}
