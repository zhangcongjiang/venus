package views

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"venus/apps/wine/models"
	commonModels "venus/common/models"
	"venus/route"
)

func init() { // 初始化时候进行路由的注册
	route.RegisterRoute(Routes) // 把需要注册的传递进去
}

func Routes(r *gin.Engine) {
	r.GET("/wines", wines)
	r.GET("/wines/top", topWines)
}

type ListWineParams struct {
	PageIndex int `form:"page_index"`
	PageSize  int `form:"page_size"`
	MaxPrice  int `form:"max_price"`
	MinPrice  int `form:"min_price"`
}

// @Tags 白酒相关接口
// @Summary 分页查询所有白酒信息
// @Description 这是一个查询白酒列表信息接口
// @Router /wines [get]
// @Param page_index query int false "第几页" default(0)
// @Param page_size query int  false "每页展示条数" default(0)
// @Param min_price query int  false "最低价格" default(0)
// @Param max_price query int  false "最高价格" default(0)
// @Produce json
// @Success 200 {object} commonModels.Result "结果"
func wines(c *gin.Context) {
	p := &ListWineParams{}
	err := c.ShouldBind(p)

	var result commonModels.Result
	if err != nil {
		result.Code = 1
		result.Msg = "fail"
		result.Data = map[string]string{"err": "参数错误"}
		c.JSON(http.StatusBadRequest, result)
		return
	}
	var listWineParams = models.ListWineParams{
		PageSize:  p.PageSize,
		PageIndex: p.PageIndex,
		MinPrice:  p.MinPrice,
		MaxPrice:  p.MaxPrice,
		OrderBy:   "",
		Asc:       true,
	}
	wines := models.ListWine(listWineParams)
	result.Code = 0
	result.Msg = "success"
	result.Data = wines
	c.JSON(http.StatusOK, result)
}

type TopWineParams struct {
	Category string `form:"category"`
	TopNum   int    `form:"top_num"`
}

// @Tags 白酒相关接口
// @Summary 查询白酒top排行相关接口
// @Description 这是一个查询白酒各类排行的接口
// @Router /wines/top [get]
// @Param category query string false "第几页" Enums(market_price,reference_price)
// @Param top_num query int  false "每页展示条数" default(5)
// @Produce json
// @Success 200 {object} commonModels.Result "结果"
func topWines(c *gin.Context) {
	p := &TopWineParams{}
	err := c.ShouldBind(p)
	var result commonModels.Result
	if err != nil {
		result.Code = 1
		result.Msg = "fail"
		result.Data = map[string]string{"err": "参数错误"}
		c.JSON(http.StatusBadRequest, result)
		return
	}
	topParams := models.TopWineParams{
		Category: p.Category,
		TopNum:   p.TopNum,
	}
	wines := models.TopWines(topParams)
	result = commonModels.Result{
		Code: 0,
		Data: wines,
		Msg:  "success",
	}
	c.JSON(http.StatusOK, result)
}
