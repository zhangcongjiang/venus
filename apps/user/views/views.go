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
// @Param data body User true "请示参数data"
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

////查询列表接口
//func ListUsers(c *gin.Context) {
//	//参数
//	search := c.Request.FormValue("search")
//	num := c.Request.FormValue("pageno")
//	pageno, err := strconv.Atoi(num)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	//得到数据集
//	datalist := models.GetUserList(pageno, 1, search)
//	//得到记录的总数
//	count := models.GetUserNum(search)
//	c.JSON(http.StatusOK, gin.H{
//		"datalist": datalist,
//		"count":    count,
//		"pagesize": 1,
//		"pageno":   pageno,
//	})
//}
//
////查询单条记录接口
//func FindUser(c *gin.Context) {
//	//参数
//	ids := c.Request.FormValue("id")
//	id, _ := strconv.Atoi(ids)
//	res := models.GetUserById(id)
//	c.JSON(http.StatusOK, gin.H{
//		"res":  res,
//		"code": 200,
//	})
//}
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
////删除某条记录接口
//func DeleteUser(c *gin.Context) {
//	//接值
//	ids := c.Request.FormValue("id")
//	id, _ := strconv.Atoi(ids)
//	models.DeleteUser(id)
//	c.JSON(http.StatusOK, gin.H{
//		"msg":  "删除成功",
//		"code": 200,
//	})
//}
