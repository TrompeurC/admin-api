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
	// 部门
	{
		router.POST("/api/dept/add", controller.CreateSysDept)
		router.GET("/api/dept/vo/list", controller.QuerySysDeptVoList)
		router.GET("/api/dept/info", controller.GetSysDeptById)
		router.PUT("/api/dept/update", controller.UpdateSysDept)
		router.DELETE("/api/dept/delete", controller.DeleteSysDeptById)
		router.GET("/api/dept/list", controller.GetSysDeptList)
	}
	// 菜单
	{
		router.GET("/api/menu/vo/list", controller.QuerySysMenuVoList)
		router.POST("/api/menu/add", controller.CreateSysMenu)
		router.GET("/api/menu/info", controller.GetSysMenu)
		router.PUT("/api/menu/update", controller.UpdateSysMenu)
		router.DELETE("/api/menu/delete", controller.DeleteSysRoleMenu)
		router.GET("/api/menu/list", controller.GetSysMenuList)
	}
	// 角色
	{
		router.POST("/api/role/add", controller.CreateSysRole)
		router.PUT("/api/role/info", controller.GetSysRoleById)
		router.PUT("/api/role/update", controller.UpdateSysRole)
		router.DELETE("/api/role/delete", controller.DeleteSysRoleById)
		router.PUT("/api/role/updateStatus", controller.UpdateSysRoleStatus)
		router.GET("/api/role/list", controller.GetSysRoleList)
		router.GET("/api/role/vo/list", controller.QuerySysRoleVoList)
		router.GET("/api/role/vo/idList", controller.QueryRoleMenuIdList)
		router.PUT("/api/role/assignPermissions", controller.AssignPermissions)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
