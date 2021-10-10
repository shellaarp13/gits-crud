package entity

import (
	"github.com/google/uuid"
)

const (
	OrderDetailsTableName = "order_details"
)

type order_details struct {
	Order_details_id uuid.UUID `gorm:type:uuid;primary_key" json:"order_details"`
	Order_number     string    `gorm:type:varchar(50);not_null" json:"order_number"`
	Product_id       uuid.UUID `gorm:type:uuid;not_null" json:"product_id"`
	Quantity_product string    `gorm:type:varchar(50);not_null" json:"quantity_product"`
	Auditable
}

func NewOrderDetails(order_details_id, product_id uuid.UUID, order_number, quantity_product string) *order_details {
	return &order_details{
		Order_details_id: order_details_id,
		Order_number:     order_number,
		Product_id:       product_id,
		Quantity_product: quantity_product,
		Auditable:        NewAuditable(),
	}
}

func (model *order_details) TableName() string {
	return OrderDetailsTableName
}
