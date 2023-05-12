package models

import (
	"gorm.io/gorm"
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
	gorm.Model
	Uuid     string `form:"uuid" json:"uuid" gorm:"primary_key;uuid"` //uuid
	Name     string `form:"name" json:"name"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	Birthday string `form:"birthday" json:"birthday"`
	Tell     string `form:"tell" json:"tell" `
	Gender   string `form:"gender" json:"gender"`
}

func (u *User) BeforeSave() error {
	t, err := time.Parse("2006-01-02", u.Birthday)
	if err != nil {
		return err
	}
	u.Birthday = t.Format("2006-01-02")
	return nil
}
