package spiders

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/google/uuid"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"venus/apps/wine/models"
	"venus/common/database"
)

var (
	db      = database.GetDB()
	baseUrl = "https://www.cnxiangyan.com"
)

func Run() {
	makeWine(baseUrl + "/jiu/article/list_102.html")
	//makeWineDetail("https://www.cnxiangyan.com/jiu/langjiu/42118.html")
}

func makeWine(url string) {
	doc := querySite(url)
	// 获取所有文章的标题和链接
	doc.Find(".price_a").Each(func(i int, s *goquery.Selection) {
		detailUrl, _ := s.Attr("href")
		makeWineDetail(baseUrl + detailUrl)

	})

	doc.Find(".tg_pages li").Each(func(i int, s *goquery.Selection) {
		if s.Find("a").Text() == "下一页" {
			nextPage, _ := s.Find("a").Attr("href")

			//nextPage:pinpai/hongshuangxi/list_45_2.html?ivk_sa=1024320u
			nextPage = strings.Split(nextPage, "?")[0]
			fmt.Println("next page:", nextPage)
			nextUrl := baseUrl + nextPage
			makeWine(nextUrl)
		}
	})
}

func makeWineDetail(url string) {
	doc := querySite(url)
	detailDoc := doc.Find(".txt")
	name := detailDoc.Find(".title h1").Text()
	priceDoc := detailDoc.Find(".price font")
	fmt.Println("name:", name)
	factoryUuid := uuid.New().String()
	factory := &models.Wine{
		Uuid:               factoryUuid,
		Name:               name,
		Brand:              detailDoc.Find(".c1 span").First().Text(),
		Origin:             detailDoc.Find(".c2 span").First().Text(),
		Factory:            detailDoc.Find(".c2 span").Eq(2).Text(),
		Alcohol:            detailDoc.Find(".c1 span").Eq(5).Text(),
		Age:                0,
		ReferencePrice:     handlePrice(priceDoc.First().Text()),
		MarketPrice:        handlePrice(priceDoc.Last().Text()),
		Description:        "",
		Capacity:           strings.Split(detailDoc.Find(".c1 span").Eq(4).Text(), "ml")[0],
		Category:           detailDoc.Find(".c1 span").Eq(1).Text(),
		RawMaterial:        detailDoc.Find(".c2 span").Eq(1).Text(),
		ShelfLife:          0,
		StorageMethod:      "",
		DrinkingSuggestion: "",
		ImageURL:           "",
		Url:                url,
	}
	db.Create(factory)
	time.Sleep(1 * time.Second)
}

func querySite(url string) *goquery.Document {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 用 goquery 解析 HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}
func handlePrice(price string) float64 {
	if price == "517；6参考价：3元" {
		return 517
	}
	if price == "暂无报价" {
		return 0
	}
	price = strings.Replace(price, " ", "", -1)
	price = strings.Replace(price, ",", "", -1)
	price = strings.Replace(price, "元", "", -1)
	price = strings.Replace(price, "坛", "", -1)
	price = strings.Replace(price, "套", "", -1)

	floatVar, err := strconv.ParseFloat(price, 64)
	if err != nil {
		log.Fatal(err)
	}
	return floatVar
}
