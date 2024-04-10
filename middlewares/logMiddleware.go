package middlewares

import (
	"admin-api/api/dao"
	"admin-api/api/entity"
	"admin-api/pkg/jwt"
	"admin-api/utils"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := strings.ToLower(c.Request.Method)
		sysAdmin, _ := jwt.GetAdmin(c)
		if method != "get" {
			log := entity.SysOperationLog{
				AdminId:    sysAdmin.ID,
				Username:   sysAdmin.Username,
				Method:     method,
				Ip:         c.ClientIP(),
				Url:        c.Request.URL.Path,
				CreateTime: utils.HTime{Time: time.Now()},
			}
			dao.CreateSysOperationLog(log)
		}
	}
}
