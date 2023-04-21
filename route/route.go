package route

import (
	"github.com/gin-gonic/gin"
	"venus/common/middleware"
)

// 自动注册的类型
type Router func(*gin.Engine)

var routers = []Router{} // 记录自动注册的操作

func RegisterRoute(routes ...Router) { // 这里三个点代表不定传参
	routers = append(routers, routes...) // 把切片展开追缴到数组中

}

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.LoggerToFile(), gin.Recovery())
	for _, route := range routers {
		route(r) // 加载路由
	}
	return r
}
