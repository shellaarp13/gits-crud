package entity

import (
	"github.com/google/uuid"
)

const (
	CustomerTableName = "customer"
)

// CustomerModel is a model for entity.Customer
type Customer struct {
	Customer_ID uuid.UUID `gorm:"type:uuid;primary_key" json:"customer_id"`
	First_Name  string    `gorm:"type:varchar(50);not_null" json:"first_name"`
	Last_Name   string    `gorm:"type:varchar(50);not_null" json:"last_name"`
	Street      string    `gorm:"type:varchar(100);not_null" json:"street"`
	Zip         string    `gorm:"type:varchar(10);not_null" json:"zip"`
	Phone       string    `gorm:"type:varchar(15);not_null" json:"phone"`
	Auditable
}

func NewCustomer(customer_id uuid.UUID, first_name, last_name, street, zip, phone string) *Customer {
	return &Customer{
		Customer_ID: customer_id,
		First_Name:  first_name,
		Last_Name:   last_name,
		Street:      street,
		Zip:         zip,
		Phone:       phone,
		Auditable:   NewAuditable(),
	}
}

func (model *Customer) TableName() string {
	return CustomerTableName
}
