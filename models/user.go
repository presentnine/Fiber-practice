package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nickname string `json:"nickname"`
}
