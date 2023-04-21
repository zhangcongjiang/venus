package models

import (
	"gorm.io/gorm"
	"venus/common/database"
)

func init() {
	err := db.AutoMigrate(&User{})
	if err != nil {
		panic("failed to migrate database")
	}
}

var (
	db = database.GetDB()
)

type User struct {
	gorm.Model
	Uuid     string `json:"uuid" gorm:"primary_key;uuid"` //uuid
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Age      int    `json:"age,omitempty"`
	Tell     string `json:"tell"`
	Gender   string `json:"gender"`
}
