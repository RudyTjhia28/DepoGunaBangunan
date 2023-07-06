package model

import "github.com/jinzhu/gorm"

// Users struct
type Users struct {
	gorm.Model
	ID       int64  `gorm:"primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"uniqueIndex"`
	FullName string `json:"fullname" gorm:"fullname"`
}

// UserLogin struct
type UserLogin struct {
	Username string `json:"username"`
	Pass     string `json:"password"`
}
