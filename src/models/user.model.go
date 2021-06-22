package models

import "gorm.io/gorm"

type User struct {
	gorm.Model // instance id,created_at,updated_at,deleted_at
	Username string
	Password string
}