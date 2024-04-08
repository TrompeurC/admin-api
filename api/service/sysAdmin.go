package service

import (
	"admin-api/api/dao"
	"admin-api/api/entity"
	"admin-api/common/result"
	"admin-api/pkg/jwt"
	"admin-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ISysAdminService interface {
	Login(*gin.Context, entity.LoginDto)
}
type SysAdminServiceImpl struct{}

func (SysAdminServiceImpl) Login(c *gin.Context, dto entity.LoginDto) {
	// 登录参数校验
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, result.ApiCode.MissingLoginParameter, result.ApiCode.GetMessage(result.ApiCode.MissingLoginParameter))
		return
	}
	// 验证码是否过期
	code := utils.RedisStore{}.Get(dto.IdKey, true)
	if len(code) == 0 {
		result.Failed(c, result.ApiCode.VerificationCodeHasExpired,
			result.ApiCode.GetMessage(result.ApiCode.VerificationCodeHasExpired))
		return
	}
	// 校验验证码
	verifyRes := CaptVerify(dto.IdKey, dto.Image)
	if !verifyRes {
		result.Failed(c, result.ApiCode.CAPTCHANOTTRUE,
			result.ApiCode.GetMessage(result.ApiCode.CAPTCHANOTTRUE))
		return
	}
	// 校验 密码
	sysAdmin := dao.SysAdminDetail(dto)
	if sysAdmin.Password != utils.EncryptionMd5(dto.Password) {
		result.Failed(c, result.ApiCode.PASSWORDNOTTRUE,
			result.ApiCode.GetMessage(result.ApiCode.PASSWORDNOTTRUE))
		return
	}
	var status uint = 2
	if sysAdmin.Status == status {
		result.Failed(c, (result.ApiCode.STATUSISENABLE),
			result.ApiCode.GetMessage(result.ApiCode.STATUSISENABLE))
		return
	}
	// 生成Token
	tokenString, _ := jwt.GenerateTokenByAdmin(sysAdmin)
	result.Success(c, map[string]interface{}{"token": tokenString, "sysAdmin": sysAdmin})
}

var sysAdminService = SysAdminServiceImpl{}

func SysAdminService() ISysAdminService {
	return &sysAdminService
}
