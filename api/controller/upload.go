package controller

import (
	"admin-api/api/service"
	"github.com/gin-gonic/gin"
)

// 单图片上传
// @Summary 单图片上传接口
// @Description 单图片上传接口
// @Produce json
// @Accept multipart/form-data
// @Param file formData file true "file" // @Success 200 {object} result.Result // @Router /api/upload [post]
func Upload(c *gin.Context) {
	service.Upload().Upload(c)
}
