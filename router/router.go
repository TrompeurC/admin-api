package router

import (
	"admin-api/common/config"
	"admin-api/middlewares"
	"github.com/gin-gonic/gin"
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
}
