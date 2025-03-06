package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string

	TenantID int    `gorm:"type:int;not null"`
	Role     string `gorm:"not null"` // Admin, Manager, User
}
