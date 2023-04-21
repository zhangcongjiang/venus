package models

import (
	"gorm.io/gorm"
)

type Cigarette struct {
	gorm.Model
	Id         int    `json:"id"  gorm:"-;primary_key;AUTO_INCREMENT"` // id
	Uuid       string `json:"uuid" gorm:"primary_key" `                // uuid
	Desc       string `json:"desc"`                                    // 描述信息
	Logo       string `json:"logo"`                                    // logo
	Brand      string `json:"brand"`                                   // 香烟品牌
	Type       string `json:"type"`                                    // 型号
	Category   string `json:"category"`                                // 类别
	Price      string `json:"price"`                                   // 香烟价格
	Length     string `json:"length"`                                  // 香烟长度
	Diameter   string `json:"diameter"`                                // 香烟直径
	Nicotine   string `json:"nicotine"`                                // 香烟尼古丁含量
	Tar        string `json:"tar"`                                     // 香烟焦油含量
	Co         string `json:"co"`                                      // 一氧化碳含量
	Packaging  string `json:"packaging"`                               // 香烟包装方式
	Menthol    bool   `json:"menthol"`                                 // 香烟是否含薄荷味
	Flavour    string `json:"flavour"`                                 // 香烟口味
	Origin     string `json:"origin"`                                  // 香烟产地
	ExpiryDate string `json:"expiry_date"`                             // 香烟保质期，格式为"yyyy-mm-dd"
	BarCode    string `json:"bar_code"`                                // 条码
	Url        string `json:"url"`                                     //数据来源URL

}

func (*Cigarette) TableName() string {
	return "cigarette"
}
