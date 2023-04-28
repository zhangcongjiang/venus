package views

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"venus/apps/user/models"
	"venus/common/database"
	commonModels "venus/common/models"
	"venus/route"
)

var (
	db = database.GetDB()
)

func init() {
	route.RegisterRoute(Routes)
}

func Routes(r *gin.Engine) {
	r.POST("/user", AddUser)
	r.GET("/users", ListUsers)
	r.GET("/user/:id", Details)
	r.DELETE("/user/:id", Delete)
	r.POST("/user/update", Update)
}

// @Tags 用户相关接口
// @Summary 创建用户
// @Description 创建用户
// @Accept  json
// @Produce  json
// @Param data body models.User true "请示参数data"
// @Success 200 {object} commonModels.Result "请求成功"
// @Failure 400 {object} commonModels.Result "请求错误"
// @Failure 500 {object} commonModels.Result "内部错误"
// @Router /user [post]
func AddUser(c *gin.Context) {

	user := &models.User{}
	err := c.ShouldBind(user)

	var result commonModels.Result
	if err != nil {
		result.Code = commonModels.CODE_ERROR
		result.Msg = "fail"
		result.Data = map[string]string{"err": "参数错误"}
		c.JSON(http.StatusBadRequest, result)
		return
	}
	user.Uuid = uuid.New().String()
	db.Create(user)
	result.Code = commonModels.CODE_SUCCESS
	result.Msg = "success"
	result.Data = user
	c.JSON(http.StatusOK, result)
}

type ListUserParams struct {
	PageIndex int `form:"page_index"`
	PageSize  int `form:"page_size"`
}

// @Tags 用户相关接口
// @Summary 分页查询所有用户
// @Description 这是一个查询用户列表信息接口
// @Router /users [get]
// @Param page_index query int false "第几页" default(0)
// @Param page_size query int  false "每页展示条数" default(0)
// @Produce json
// @Success 200 {object} commonModels.Result "结果"
func ListUsers(c *gin.Context) {
	p := &ListUserParams{}
	err := c.ShouldBind(p)

	var result commonModels.Result
	if err != nil {
		result.Code = commonModels.CODE_ERROR
		result.Msg = "fail"
		result.Data = map[string]string{"err": "参数错误"}
		c.JSON(http.StatusBadRequest, result)
		return
	}
	var users []models.User
	if p.PageIndex <= 0 {
		p.PageIndex = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
	db.Limit(p.PageSize).Offset((p.PageIndex - 1) * p.PageSize).Order("id").Find(&users)

	result.Code = commonModels.CODE_SUCCESS
	result.Msg = "success"
	result.Data = users
	c.JSON(http.StatusOK, result)
}

// @Tags 用户相关接口
// @Summary 查询指定用户详细信息
// @Description 这是一个查询用户详细信息接口
// @Router /user/{id} [get]
// @Param id path string true "ID"
// @Produce json
// @Success 200 {object} commonModels.Result "结果"
func Details(c *gin.Context) {
	//参数
	id := c.Param("id")
	var user models.User
	db.Where("uuid = ?", id).Take(&user)
	result := commonModels.Result{
		Code: commonModels.CODE_SUCCESS,
		Data: user,
		Msg:  "success",
	}
	c.JSON(http.StatusOK, result)
}

// @Tags 用户相关接口
// @Summary 更新用户信息
// @Description 创建用户
// @Accept  json
// @Produce  json
// @Param data body models.User true "请示参数data"
// @Success 200 {object} commonModels.Result "请求成功"
// @Failure 400 {object} commonModels.Result "请求错误"
// @Failure 500 {object} commonModels.Result "内部错误"
// @Router /user/update [post]
func Update(c *gin.Context) {
	user := &models.User{}
	err := c.ShouldBind(user)

	var result commonModels.Result
	if err != nil {
		result.Code = commonModels.CODE_ERROR
		result.Msg = "fail"
		result.Data = map[string]string{"err": "参数错误"}
		c.JSON(http.StatusBadRequest, result)
		return
	}
	m := map[string]interface{}{}
	if user.Tell != "" {
		m["tell"] = user.Tell
	}
	if user.Email != "" {
		m["email"] = user.Email
	}
	if user.Age != 0 {
		m["age"] = user.Age
	}
	if user.Password != "" {
		m["password"] = user.Password
	}

	db.Model(&user).Updates(m)
	result.Code = commonModels.CODE_SUCCESS
	result.Msg = "success"
	result.Data = user
	c.JSON(http.StatusOK, result)
}

// @Tags 用户相关接口
// @Summary 删除用户信息
// @Description 这是一个删除用户信息接口
// @Router /user/{id} [delete]
// @Param id path string true "ID"
// @Produce json
// @Success 200 {object} commonModels.Result "结果"
func Delete(c *gin.Context) {
	//接值
	id := c.Param("id")
	var user models.User
	db.Where("uuid = ?", id).Take(&user)
	result := new(commonModels.Result)
	if user.Uuid != "" {
		db.Delete(&user)
		result.SetCode(commonModels.CODE_SUCCESS)
		result.SetData(user)
		result.SetMessage("success")
	} else {
		result.SetMessage("Not Exist")
		result.SetCode(commonModels.CODE_ERROR)
	}

	c.JSON(http.StatusOK, result)
}
