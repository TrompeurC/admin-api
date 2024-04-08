package main

import (
	"admin-api/common/config"
	_ "admin-api/docs"
	"admin-api/pkg/db"
	"admin-api/pkg/log"
	"admin-api/pkg/redis"
	"admin-api/router"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

// @title 通用后台管理系统
// @version 1.0
// @description 后台管理系统API接口文档
// @securityDefinitions.apikey ApiKeyAuth // @in header
// @name Authorization
func main() {
	// 加载日志log
	log := log.Log()
	// 设置启动模式
	gin.SetMode(config.Config.Server.Model)
	// 初始化路由
	router := router.InitRouter()
	srv := &http.Server{
		Addr:    config.Config.Server.Address,
		Handler: router,
	}
	// 启动服务
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Info("listen: %s \n", err)
		} else {
			log.Info("listen: %s \n", srv.Addr)
		}
	}()
	quit := make(chan os.Signal)
	// 监听消息
	<-quit
	log.Info("Shutdown Server ...")
	// 关闭服务
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Info("Server Shutdown:", err)
	}
	log.Info("Server exiting")
}

// 初始化连接
func init() {
	// mysql
	db.SetupDBLink()
	// redis
	redis.SetupRedisDB()
}
