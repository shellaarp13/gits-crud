package entity

const (
	AccountTableName = "account"
)

//AccountModel is a model for entity.Account
type Account struct {
	Username string `gorm:"type:varchar(50);primary_key" json:"username"`
	Password string `gorm:"type:varchar(150);not_null" json:"password"`
	Auditable
}

func NewAccount(username, password string) *Account {
	return &Account{
		Username:  username,
		Password:  password,
		Auditable: NewAuditable(),
	}
}

func (model *Account) TableName() string {
	return AccountTableName
}
