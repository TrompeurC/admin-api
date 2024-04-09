package router

import (
	"admin-api/api/controller"
	"admin-api/common/config"
	"admin-api/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	// 当机时候恢复
	router.Use(gin.Recovery())
	// 跨域中间件
	router.Use(middlewares.Cors())
	// 日志中间件
	router.Use(middlewares.Logger())

	// 图片访问路径静态文件夹可直接访问 Static 文件路径也可以展示
	router.StaticFS(config.Config.ImageSettings.UploadDir, http.Dir(config.Config.ImageSettings.UploadDir))
	register(router)
	return router
}

// 路由接口
func register(router *gin.Engine) {
	// todo 后续接口url

	ApiGroup := router.Group("/api")
	{
		ApiGroup.GET("captcha", controller.Captcha)
		ApiGroup.POST("login", controller.Login)
	}
	//岗位
	{
		router.POST("/api/post/add", controller.CreateSysPost)
		router.GET("/api/post/info", controller.GetSysPostById)
		router.PUT("/api/post/update", controller.UpdateSysPost)
		router.DELETE("/api/post/delete", controller.DeleteSysPostById)
		router.DELETE("/api/post/batch/delete", controller.BatchDeleteSysPost)
		router.PUT("/api/post/updateStatus", controller.UpdateSysPostStatus)
		router.GET("/api/post/list", controller.GetSysPostList)
		router.GET("/api/post/vo/list", controller.QuerySysPostVoList)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
