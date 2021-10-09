package entity

import (
	"github.com/google/uuid"
)

const (
	OderDetailsTableName = "order_details"
)

type order_details struct {
	Oder_details_id  uuid.UUID `gorm:type:uuid;not_null" json:"oder_details"`
	Oder_number      int64     `gorm:type:integer;not_null" json:"oder_number"`
	Product_id       uuid.UUID `gorm:type:uuid;not_null" json:"product_id"`
	Quantity_product int64     `gorm:type:integer;not_null" json:"quantity_product"`
	Order            *Order    `gorm:type:"foreignKey:Order_number" json:"for_order_number`
	Product          *Product  `gorm:type:"foreignKey:Product_id" json:"for_product_id"`
	Auditable
}

func NewOderDetails(oder_details_id, product_id uuid.UUID, order_number, quantity_product int) *OderDetails {
	return &OrderDetails{
		Order_details_id: order_details_id,
		Order_number:     order_number,
		Product_id:       product_id,
		Quantity_product: quantity_product,
		Auditable:        NewAuditable(),
	}
}

func (model *OderDetails) TableName() string {
	return OderDetailsTableName
}
