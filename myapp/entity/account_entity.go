package entity

import (
	"github.com/google/uuid"
)

const (
	AccountTableName = "Account"
)

//AccountModel is a model for entity.Account
type Account struct {
	Username uuid.UUID `gorm:"type:uuid;primary_key" json:"username"`
	Password string    `gorm:"type:varchar(50);not_null" json:"password"`
}

func NewAccount(Username uuid.UUID, password string) *Account {
	return &Account{
		Username:  username,
		Password:  password,
		Auditable: NewAuditable(),
	}
}

func (model *Account) TableName() string {
	return AccountTableName
}
