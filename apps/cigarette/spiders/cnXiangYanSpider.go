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
	"venus/apps/cigarette/models"
	"venus/common/database"
)

var (
	db = database.GetDB()
)

func Run() {

	makeTobaccoFactory("https://www.cnxiangyan.com")
	time.Sleep(1 * time.Second)

	var factories []models.TobaccoFactory

	db.Find(&factories)

	for index, v := range factories {
		if index >= 105 {
			fmt.Println("index", index)
			fmt.Println("site", v.Href)
			makeCigarette(v.Href)
		}

	}
}

func makeTobaccoFactory(url string) {

	doc := querySite(url)
	// 获取所有文章的标题和链接
	doc.Find(".brandsList a").Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		link, _ := s.Attr("href")
		factoryUuid := uuid.New().String()
		factory := &models.TobaccoFactory{
			Uuid: factoryUuid,
			Name: title,
			Url:  "https://www.cnxiangyan.com",
			Href: "https://www.cnxiangyan.com" + link,
		}
		db.Create(factory)
	})
}

func makeCigarette(url string) {
	// 发送 GET 请求并获取响应
	doc := querySite(url)
	var cigarette models.Cigarette
	db := database.GetDB()
	doc.Find(".xy_ul li").Each(func(i int, s *goquery.Selection) {
		if len(s.Find(".xy_right").Nodes) > 0 {

			var (
				name      string
				category  string
				price     string
				cType     string
				tar       string
				co        string
				nicotine  string
				length    string
				packaging string
				barCode   string
			)
			nameDoc, _ := s.Find(".xy_tit a").Attr("title")
			detail, _ := s.Find(".xy_tit a").Attr("href")
			detailUrl := "https://www.cnxiangyan.com" + detail

			detailDoc := querySite(detailUrl)

			rt := []rune(nameDoc)
			validKeyList := []int{65288, 65289, 40, 41, 125}
			for j := 0; j < len(rt); j++ {
				for _, v := range validKeyList {
					if v == int(rt[j]) {
						rt[j] = 32
					}
				}

			}
			nameDoc = string(rt)
			fmt.Println("name:", nameDoc)
			nameDocList := strings.Split(nameDoc, " ")
			name = nameDocList[0]
			if len(nameDocList) > 1 {
				category = strings.Split(nameDoc, " ")[1]
			}

			detailDoc.Find(".show_p span").Each(func(j int, g *goquery.Selection) {

				if find(g.Text(), "单盒参考价") {
					node := g.Next()
					price = formatStrings(node.Text())
				}
			})

			detailDoc.Find(".show_nt").Each(func(j int, g *goquery.Selection) {
				if find(g.Text(), "类型") {
					node := g.Next()
					cType = formatStrings(node.Text())
				}
				if find(g.Text(), "焦油量") {
					node := g.Next()
					tar = formatStrings(node.Text())
				}
				if find(g.Text(), "烟气烟碱量") {
					node := g.Next()
					nicotine = formatStrings(node.Text())
				}
				if find(g.Text(), "一氧化碳量") {
					node := g.Next()
					co = formatStrings(node.Text())
				}
				if find(g.Text(), "烟长") {
					node := g.Next()
					length = formatStrings(node.Text())
				}
				if find(g.Text(), "包装形式") {
					node := g.Next()
					packaging = formatStrings(node.Text())
				}
				if find(g.Text(), "卷烟条码") {
					node := g.Next()
					barCode = formatStrings(node.Text())
				}

			})
			cigaretteUuid := uuid.New().String()
			cigarette = models.Cigarette{
				Uuid:       cigaretteUuid,
				Desc:       "",
				Logo:       "",
				Brand:      name,
				Type:       cType,
				Category:   category,
				Price:      price,
				Length:     length,
				Diameter:   "",
				Nicotine:   nicotine,
				Tar:        tar,
				Co:         co,
				Packaging:  packaging,
				Menthol:    false,
				Flavour:    "",
				Origin:     "",
				ExpiryDate: "",
				BarCode:    barCode,
				Url:        detailUrl,
			}
			db.Create(&cigarette)
			time.Sleep(1 * time.Second)
		}

	})

	// 判断是否有下一页
	doc.Find(".tg_pages li").Each(func(i int, s *goquery.Selection) {
		if s.Find("a").Text() == "下一页" {
			nextPage, _ := s.Find("a").Attr("href")

			//nextPage:pinpai/hongshuangxi/list_45_2.html?ivk_sa=1024320u
			nextPage = strings.Split(nextPage, "?")[0]
			fmt.Println("next page:", nextPage)
			//nextPage: pinpai/changcheng/list_25_3.html 校验爬取前九页数据，防止死循环
			front := strings.Split(nextPage, ".")[0]
			after := strings.Split(front, "_")
			intValue, err := strconv.Atoi(after[len(after)-1])
			if err != nil {
				return
			}

			if intValue < 10 {
				nextUrl := "https://www.cnxiangyan.com" + nextPage
				makeCigarette(nextUrl)
			} else {
				return
			}

		}
	})
}

func find(values string, value string) bool {
	return strings.Contains(values, value)
}

func formatStrings(s string) string {
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	return s
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
