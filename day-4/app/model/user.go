package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64 `gorm:"primaryKey"`
	Name      string
	Password  string
	Email     string `gorm:"index:idx_name,unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
