package models

type TobaccoFactory struct {
	Id       int    `json:"id" gorm:"id;-;primary_key;AUTO_INCREMENT"` //id
	Uuid     string `json:"uuid" gorm:"primary_key;uuid"`              //uuid
	Name     string `json:"name" gorm:"name"`                          //名称
	Desc     string `json:"desc" gorm:"desc"`                          //描述信息
	Location string `json:"location" gorm:"location"`                  //地址
	Url      string `json:"url" gorm:"url"`                            //数据来源Url
	Href     string `json:"href`                                       //指向的链接
}

func (f *TobaccoFactory) TableName() string {
	return "tobacco_factories"
}
