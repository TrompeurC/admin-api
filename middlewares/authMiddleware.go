package middlewares

import (
	"admin-api/common/result"
	"admin-api/constant"
	"admin-api/pkg/jwt"
	"github.com/gin-gonic/gin"
	"strings"
)

// AuthMiddleware 鉴权
func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		authToken := context.Request.Header.Get("Authorization")
		if authToken == "" {
			result.Failed(context, result.ApiCode.NOAUTH, result.ApiCode.GetMessage(result.ApiCode.NOAUTH))
			context.Abort()
			return
		} else {
			parts := strings.SplitN(authToken, " ", 2)
			if !(len(parts) == 2 && parts[0] == "Bearer") {
				result.Failed(context, result.ApiCode.AUTHFORMATERROR, result.ApiCode.GetMessage(result.ApiCode.AUTHFORMATERROR))
				context.Abort()
				return
			} else {
				// todo 鉴权
				var token = parts[1]

				mc, err := jwt.ValidateToken(token)
				if err != nil {
					result.Failed(context, (result.ApiCode.INVALIDTOKEN),
						result.ApiCode.GetMessage(result.ApiCode.INVALIDTOKEN))
					context.Abort()
					return
				}
				// 存用户信息
				context.Set(constant.ContextKeyUserObj, mc)
				context.Next()
			}
		}
	}
}
