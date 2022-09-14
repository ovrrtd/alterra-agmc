package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64 `gorm:"primaryKey"`
	Name      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
