package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Result 通用结构

type Result struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Success(ctx *gin.Context, data any) {
	if data == nil {
		data = gin.H{}
	}

	r := Result{}
	r.Code = ApiCode.SUCCESS
	r.Message = ApiCode.GetMessage(ApiCode.SUCCESS)
	r.Data = data
	ctx.JSON(http.StatusOK, r)
}

func Failed(ctx *gin.Context, code uint, message string) {
	r := Result{}
	r.Message = message
	r.Code = code
	r.Data = gin.H{}
	ctx.JSON(http.StatusOK, r)
}
