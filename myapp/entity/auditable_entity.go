package entity

import (
	"time"

	"gorm.io/gorm"
)

type Auditable struct {
	CreatedAt time.Time      `gorm:"type:timestamptz;not_null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:timestamptz;not_null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func NewAuditable() Auditable {
	return Auditable{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
