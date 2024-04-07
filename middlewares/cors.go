package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 跨域中间件设置
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-type,Authorization")
		context.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS")
		context.Header("Access-Control-Expose-Headers", `Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type`)
		context.Header("Access-Control-Allow-Credentials", "true")
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		context.Next()
	}
}
