package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"venus/apps/cigarette/models"
	_ "venus/apps/user/views"
	wineModels "venus/apps/wine/models"
	cigaretteSpider "venus/apps/wine/spiders"
	wineSpider "venus/apps/wine/spiders"
	_ "venus/apps/wine/views"
	"venus/common/database"
	commonModels "venus/common/models"
	_ "venus/docs"
	"venus/route"
)

// @title 我的swagger
// @version 1.0
// @description 这里写描述信息
// @host localhost:8888
// @BasePath /
func main() {
	database.Setup()

	r := route.InitRouter()
	initRoutes(r)

	if err := r.Run(":8888"); err != nil {
		panic(err)
	}

}

func initRoutes(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func Spider() {
	db := database.GetDB()
	err := db.AutoMigrate(&commonModels.Site{})
	if err != nil {
		panic("failed to migrate database")
	}
	err = db.AutoMigrate(&models.Cigarette{})
	if err != nil {
		panic("failed to migrate database")
	}
	err = db.AutoMigrate(&models.TobaccoFactory{})
	if err != nil {
		panic("failed to migrate database")
	}
	err = db.AutoMigrate(&wineModels.Wine{})
	if err != nil {
		panic("failed to migrate database")
	}
	err = db.AutoMigrate(&wineModels.WineFactory{})
	if err != nil {
		panic("failed to migrate database")
	}
	cigaretteSpider.Run()
	wineSpider.Run()
}
