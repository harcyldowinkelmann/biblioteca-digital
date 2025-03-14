package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
	CreatedAt *time.Time     `gorm:"column:created_at;autoCreateTime" json:"CreatedAt"`
	UpdatedAt *time.Time     `gorm:"column:updated_at;autoUpdateTime" json:"UpdatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"DeletedAt"`
}
