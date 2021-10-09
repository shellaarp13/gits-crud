package entity

import (
	"time"

	"github.com/google/uuid"
)

const (
	OrderTableName = "dosen"
)

// OrderModel is a model for entity.Order
type Order struct {
	Order_Number  uuid.UUID `gorm:"type:uuid;not_null" json:"order_number"`
	Customer_ID   uuid.UUID `gorm:"type:uuid;not_null" json:"customer_id"`
	Customer_Name string    `gorm:"type:varchar(50);not_null" json:"customer_name"`
	To_street     string    `gorm:"type:varchar(100);not_null" json:"to_street"`
	To_city       string    `gorm:"type:varchar(50);not_null" json:"to_city"`
	To_zip        string    `gorm:"type:varchar(10);not_null" json:"to_zip"`
	Ship_date     time.Time `gorm:"type:timestamptz;not_null" json:"ship_date"`
	Auditable
}

func NewOrder(order_number, customer_id uuid.UUID, customer_name, to_street, to_city, to_zip string, ship_date time.Time) *Order {
	return &Order{
		Order_Number:  order_number,
		Customer_ID:   customer_id,
		Customer_Name: customer_name,
		To_street:     to_street,
		To_city:       to_city,
		To_zip:        to_zip,
		Ship_date:     ship_date,
		Auditable:     NewAuditable(),
	}
}

func (model *Order) TableName() string {
	return OrderTableName
}
