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
}

type UserParams struct {
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
	Age      int    `form:"age,omitempty"`
	Tell     string `form:"tell"`
	Gender   string `form:"gender"`
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

	p := &UserParams{}
	err := c.ShouldBind(p)

	var result commonModels.Result
	if err != nil {
		result.Code = commonModels.CODE_ERROR
		result.Msg = "fail"
		result.Data = map[string]string{"err": "参数错误"}
		c.JSON(http.StatusBadRequest, result)
		return
	}
	var user = &models.User{
		Uuid:     uuid.New().String(),
		Name:     p.Name,
		Email:    p.Email,
		Password: p.Password,
		Age:      p.Age,
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

//
////修改单条记录接口
//func UpdateUser(c *gin.Context) {
//	//接值
//	ids := c.Request.FormValue("id")
//	id, _ := strconv.Atoi(ids)
//	name := c.Request.FormValue("name")
//	cgender := c.Request.FormValue("gender")
//	gender, _ := strconv.Atoi(cgender)
//	cage := c.Request.FormValue("age")
//	age, _ := strconv.Atoi(cage)
//	//赋值
//	user := models.User{
//		ID:     id,
//		Name:   name,
//		Gender: gender,
//		Age:    age,
//	}
//	//调用模型中修改方法
//	user.EditUser()
//	c.JSON(http.StatusOK, gin.H{
//		"msg":  "修改成功",
//		"code": 200,
//	})
//}
//

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
	if user.ID != 0 {
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
