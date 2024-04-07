package middlewares

import (
	"admin-api/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		logger := log.Log()
		// 开始时间
		startTime := time.Now()
		// 处理请求
		context.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime) / time.Millisecond

		//请求方式
		method := context.Request.Method
		//请求路由
		reqUri := context.Request.RequestURI
		header := context.Request.Header
		proto := context.Request.Proto

		// 状态码
		statusCode := context.Writer.Status()
		// 请求Ip
		ip := context.ClientIP()
		err := context.Err()
		body, _ := ioutil.ReadAll(context.Request.Body)
		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    ip,
			"req_method":   method,
			"req_uri":      reqUri,
			"header":       header,
			"proto":        proto,
			"err":          err,
			"body":         body,
		}).Info()
	}
}
