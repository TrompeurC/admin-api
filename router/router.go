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
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// jwt鉴权接口
	jwt := router.Group("/api", middlewares.AuthMiddleware())
	//岗位
	{
		jwt.POST("/post/add", controller.CreateSysPost)
		jwt.GET("/post/info", controller.GetSysPostById)
		jwt.PUT("/post/update", controller.UpdateSysPost)
		jwt.DELETE("/post/delete", controller.DeleteSysPostById)
		jwt.DELETE("/post/batch/delete", controller.BatchDeleteSysPost)
		jwt.PUT("/post/updateStatus", controller.UpdateSysPostStatus)
		jwt.GET("/post/list", controller.GetSysPostList)
		jwt.GET("/post/vo/list", controller.QuerySysPostVoList)
	}
	// 部门
	{
		jwt.POST("/dept/add", controller.CreateSysDept)
		jwt.GET("/dept/vo/list", controller.QuerySysDeptVoList)
		jwt.GET("/dept/info", controller.GetSysDeptById)
		jwt.PUT("/dept/update", controller.UpdateSysDept)
		jwt.DELETE("/dept/delete", controller.DeleteSysDeptById)
		jwt.GET("/dept/list", controller.GetSysDeptList)
	}
	// 菜单
	{
		jwt.GET("/menu/vo/list", controller.QuerySysMenuVoList)
		jwt.POST("/menu/add", controller.CreateSysMenu)
		jwt.GET("/menu/info", controller.GetSysMenu)
		jwt.PUT("/menu/update", controller.UpdateSysMenu)
		jwt.DELETE("/menu/delete", controller.DeleteSysRoleMenu)
		jwt.GET("/menu/list", controller.GetSysMenuList)
	}
	// 角色
	{
		jwt.POST("/role/add", controller.CreateSysRole)
		jwt.PUT("/role/info", controller.GetSysRoleById)
		jwt.PUT("/role/update", controller.UpdateSysRole)
		jwt.DELETE("/role/delete", controller.DeleteSysRoleById)
		jwt.PUT("/role/updateStatus", controller.UpdateSysRoleStatus)
		jwt.GET("/role/list", controller.GetSysRoleList)
		jwt.GET("/role/vo/list", controller.QuerySysRoleVoList)
		jwt.GET("/role/vo/idList", controller.QueryRoleMenuIdList)
		jwt.PUT("/role/assignPermissions", controller.AssignPermissions)
	}
	// 用户
	{
		jwt.POST("/admin/add", controller.CreateSysAdmin)
		jwt.GET("/admin/info", controller.GetSysAdminInfo)
		jwt.PUT("/admin/update", controller.UpdateSysAdmin)
		jwt.DELETE("/admin/delete", controller.DeleteSysAdminById)
		jwt.PUT("/admin/updateStatus", controller.UpdateSysAdminStatus)
		jwt.PUT("/admin/updatePassword", controller.ResetSysAdminPassword)
		jwt.PUT("/admin/updatePersonal", controller.UpdatePersonal)
		jwt.PUT("/admin/updatePersonalPassword",
			controller.UpdatePersonalPassword)
		jwt.GET("/admin/list", controller.GetSysAdminList)
	}

	// 文件上传
	{
		jwt.POST("/upload", controller.Upload)
	}
}
