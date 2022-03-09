package models

import "gorm.io/gorm"

type Administrator struct {
	gorm.Model
	Id       uint   `json:"id" gorm:"primaryKey AUTO_INCREMENT"`
	Token    string `json:"token" gorm:"unique"`
	Username string `gorm:"unique"`
	Password string
}
