package entity

import (
	"github.com/google/uuid"
)

const (
	OrderDetailsTableName = "order_details"
)

type Order_Details struct {
	Order_details_id uuid.UUID `gorm:"type:uuid;primary_key" json:"order_details"`
	Order_Number     uuid.UUID `gorm:"type:uuid;not_null" json:"order_number"`
	Product_ID       uuid.UUID `gorm:"type:uuid;not_null" json:"product_id"`
	Quantity_product int32     `gorm:"type:integer;not_null" json:"quantity_product"`

	Order   *Order   `gorm:"foreignKey:Order_Number" json:"for_order_number"`
	Product *Product `gorm:"foreignKey:Product_ID" json:"for_product_id"`
	Auditable
}

func NewOrderDetails(order_details_id, order_number, product_id uuid.UUID, quantity_product int32) *Order_Details {
	return &Order_Details{
		Order_details_id: order_details_id,
		Order_Number:     order_number,
		Product_ID:       product_id,
		Quantity_product: quantity_product,
		Auditable:        NewAuditable(),
	}
}

func (model *Order_Details) TableName() string {
	return OrderDetailsTableName
}
