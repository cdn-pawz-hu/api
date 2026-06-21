package models

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"uniqueIndex;not null;size:255"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"type:varchar(20);default:'user'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
