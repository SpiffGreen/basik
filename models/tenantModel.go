package models

import "gorm.io/gorm"

type Tenant struct {
	gorm.Model
	Name  string `json:"name"`
	Users []User `json:"users" gorm:"foreignKey:TenantID"`
	Tasks []Task `json:"tasks" gorm:"foreignKey:TenantID"`
}
