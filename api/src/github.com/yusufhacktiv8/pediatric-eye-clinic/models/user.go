package models

import (
	"github.com/jinzhu/gorm"
)

// User is a model for user
type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Role     []Role `gorm:"many2many:user_roles;"`
}
