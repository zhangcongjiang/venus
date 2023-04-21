package models

import "gorm.io/gorm"

type Site struct {
	gorm.Model
	Id         int    `json:"id"  gorm:"-;primary_key;AUTO_INCREMENT"` // id
	Uuid       string `json:"uuid" gorm:"primary_key"`                 // uuid
	Url        string `json:"url"`
	ResourceId string `json:"resource_id"`
}
