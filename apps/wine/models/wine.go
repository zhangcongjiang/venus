package models

import (
	"fmt"
	"gorm.io/gorm"
	"math"
	"venus/common/database"
)

type Wine struct {
	gorm.Model
	Id                 int     `json:"id"  gorm:"-;primary_key;AUTO_INCREMENT"` // id
	Uuid               string  `json:"uuid" gorm:"primary_key" `                // uuid
	Name               string  `json:"name"`                                    // 白酒名称
	Brand              string  `json:"brand"`                                   // 白酒品牌
	Origin             string  `json:"origin"`                                  // 白酒产地
	Factory            string  `json:"factory"`                                 // 酒厂
	Alcohol            string  `json:"alcohol"`                                 // 酒精度数
	Age                int     `json:"age"`                                     // 酒的年份
	ReferencePrice     float64 `json:"reference_price"`                         // 参考价格
	MarketPrice        float64 `json:"market_price"`                            // 市面价格
	Description        string  `json:"description"`                             // 酒的描述
	Capacity           string  `json:"capacity"`                                // 酒的容量，单位为毫升
	Category           string  `json:"category"`                                // 酒的分类，例如白酒、黄酒、啤酒等
	RawMaterial        string  `json:"raw_material"`                            // 酒的原料，例如高粱、小麦、玉米等
	ShelfLife          int     `json:"shelf_life"`                              // 酒的保质期，单位为月
	StorageMethod      string  `json:"storage_method"`                          // 酒的存储方式，例如常温、冷藏、避光等
	DrinkingSuggestion string  `json:"drinking_suggestion"`                     // 饮用建议，例如搭配什么食物、什么场合饮用等
	ImageURL           string  `json:"image_url"`                               // 酒的图片链接
	Url                string  `json:"url"`                                     // 数据来源
}

type WineFactory struct {
	Id       int    `json:"id" gorm:"id;-;primary_key;AUTO_INCREMENT"` //id
	Uuid     string `json:"uuid" gorm:"primary_key;uuid"`              //uuid
	Name     string `json:"name"`                                      //名称
	Desc     string `json:"desc"`                                      //描述信息
	Location string `json:"location"`                                  //地址
	Category string `json:"category"`                                  //酒的类型，例如酱香型，清香型
	Url      string `json:"url"`                                       //数据来源Url
	Href     string `json:"href"`                                      //指向的链接
}

type ListWineParams struct {
	PageSize  int
	PageIndex int
	MaxPrice  int
	MinPrice  int
	OrderBy   string //排序key
	Asc       bool   //升序
}

type TopWineParams struct {
	Category string
	TopNum   int
}

var (
	db = database.GetDB()
)

func ListWine(listWineParams ListWineParams) []Wine {
	if listWineParams.PageIndex <= 0 {
		listWineParams.PageIndex = 1
	}
	if listWineParams.PageSize <= 0 {
		listWineParams.PageSize = 100
	}
	if listWineParams.OrderBy == "" {
		listWineParams.OrderBy = "market_price"
	}
	if listWineParams.MinPrice <= 0 {
		listWineParams.MinPrice = 0
	}
	if listWineParams.MaxPrice <= 0 {
		listWineParams.MaxPrice = math.MaxInt
	}
	var wines []Wine
	var order string
	if listWineParams.Asc {
		order = fmt.Sprintf("%s %s", listWineParams.OrderBy, "asc")
	} else {
		order = fmt.Sprintf("%s %s", listWineParams.OrderBy, "desc")
	}
	db.Where("market_price BETWEEN  ? AND ?", listWineParams.MinPrice, listWineParams.MaxPrice).Limit(listWineParams.PageSize).Offset((listWineParams.PageIndex - 1) * listWineParams.PageSize).Order(order).Find(&wines)
	return wines
}

func TopWines(p TopWineParams) []Wine {
	var wines []Wine
	order := fmt.Sprintf("%s desc", p.Category)
	db.Limit(p.TopNum).Order(order).Find(&wines)
	return wines
}
