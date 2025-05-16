package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	CreatedBy string     `gorm:"column:created_by;not null" json:"created_by"`
	UpdatedAt time.Time  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	UpdatedBy *string    `gorm:"column:updated_by" json:"updated_by,omitempty"`
}

func (model *BaseModel) BeforeCreate(tx *gorm.DB) error {
	model.CreatedBy = "System"
	return nil
}

func (model *BaseModel) BeforeUpdate(tx *gorm.DB) error {
	user := "System"
	model.UpdatedBy = &user
	return nil
}
