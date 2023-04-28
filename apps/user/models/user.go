package models

import (
	"fmt"
	"time"
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
	Uuid     string     `form:"uuid" json:"uuid" gorm:"primary_key;uuid"` //uuid
	Name     string     `form:"name" json:"name"`
	Email    string     `form:"email" json:"email"`
	Password string     `form:"password" json:"password"`
	Birthday *LocalTime `form:"birthday" json:"birthday"`
	Tell     string     `form:"tell" json:"tell"`
	Gender   string     `form:"gender" json:"gender"`
}
