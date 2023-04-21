package models

import "gorm.io/gorm"

type RelatePic struct {
	gorm.Model
	Id       int    `json:"id"  gorm:"-;primary_key;AUTO_INCREMENT"` // id
	WineUuid string `json:"wine_uuid"`                               // wine uuid
	Name     string `json:"name"`                                    //图片名字
}
