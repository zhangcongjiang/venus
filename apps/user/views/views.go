package views

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"venus/apps/user/models"
	"venus/common/database"
	commonModels "venus/common/models"
	"venus/route"
)

var (
	db     = database.GetDB()
	logger = logrus.WithFields(logrus.Fields{
		"module": "user views",
	})
)

func init() {
	route.RegisterRoute(Routes)

}

func Routes(r *gin.Engine) {
	r.POST("/user", AddUser)
	r.GET("/users", ListUsers)
	r.GET("/user/:uuid", Details)
	r.DELETE("/user/:uuid", Delete)
	r.POST("/user/update/:uuid", Update)
}

type UserParam struct {
	Name     string `form:"name" json:"name" example:"张三"`
	Email    string `form:"email" json:"email" example:"8888888@qq.com" binding:"required,email"`
	Password string `form:"password" json:"password" example:"xxxxxxxx"`
	Birthday string `form:"birthday" json:"birthday" example:"1992-09-01" binding:"required,datetime=2006-01-02"`
	Tell     string `form:"tell" json:"tell" example:"133-3333-3333" binding:"required,phone"`
	Gender   string `form:"gender" json:"gender" example:"male"`
}

// @Tags 用户相关接口
// @Summary 创建用户
// @Description 创建用户
// @Accept  json
// @Produce  json
// @Param data body UserParam true "请示参数data"
// @Success 200 {object} commonModels.Result "请求成功"
// @Failure 400 {object} commonModels.Result "请求错误"
// @Failure 500 {object} commonModels.Result "内部错误"
// @Router /user [post]
func AddUser(c *gin.Context) {

	p := &UserParam{}
	err := c.ShouldBind(p)

	var result commonModels.Result
	if err != nil {
		logger.Error(err.Error())
		result.Code = commonModels.CODE_ERROR
		result.Msg = "fail"
		result.Data = map[string]string{"err": "参数错误"}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	user := &models.User{
		Uuid:     uuid.New().String(),
		Name:     p.Name,
		Email:    p.Email,
		Password: p.Password,
		Birthday: p.Birthday,
		Tell:     p.Tell,
		Gender:   p.Gender,
	}
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
// @Router /user/{uuid} [get]
// @Param uuid path string true "UUID"
// @Produce json
// @Success 200 {object} commonModels.Result "结果"
func Details(c *gin.Context) {
	//参数
	id := c.Param("uuid")
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
// @Param uuid path string true "UUID"
// @Param data body UserParam true "请示参数data"
// @Success 200 {object} commonModels.Result "请求成功"
// @Failure 400 {object} commonModels.Result "请求错误"
// @Failure 500 {object} commonModels.Result "内部错误"
// @Router /user/update/{uuid} [post]
func Update(c *gin.Context) {

	p := &UserParam{}
	err := c.ShouldBind(p)

	var result commonModels.Result
	if err != nil {
		result.Code = commonModels.CODE_ERROR
		result.Msg = "fail"
		result.Data = map[string]string{"err": "参数错误"}
		c.JSON(http.StatusBadRequest, result)
		return
	}
	id := c.Param("uuid")
	var user models.User
	db.Where("uuid = ?", id).Take(&user)
	m := map[string]interface{}{}
	if p.Tell != "" {
		m["tell"] = p.Tell
		user.Tell = p.Tell
	}
	if p.Email != "" {
		m["email"] = p.Email
		user.Email = p.Email
	}
	if p.Password != "" {
		m["password"] = p.Password
		user.Password = p.Password
	}
	if p.Name != "" {
		m["name"] = p.Name
		user.Name = p.Name
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
// @Router /user/{uuid} [delete]
// @Param uuid path string true "UUID"
// @Produce json
// @Success 200 {object} commonModels.Result "结果"
func Delete(c *gin.Context) {
	//接值
	id := c.Param("uuid")
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
