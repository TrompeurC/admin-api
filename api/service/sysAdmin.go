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

// ISysAdminService 定义接口
type ISysAdminService interface {
	Login(c *gin.Context, dto entity.LoginDto)
	CreateSysAdmin(c *gin.Context, dto entity.AddSysAdminDto)
	GetSysAdminInfo(c *gin.Context, Id int)
	UpdateSysAdmin(c *gin.Context, dto entity.UpdateSysAdminDto)
	DeleteSysAdminById(c *gin.Context, dto entity.SysAdminIdDto)
	UpdateSysAdminStatus(c *gin.Context, dto entity.UpdateSysAdminStatusDto)
	ResetSysAdminPassword(c *gin.Context, dto entity.ResetSysAdminPasswordDto)
	GetSysAdminList(c *gin.Context, PageSize, PageNum int, Username, Status,
		BeginTime, EndTime string)
	UpdatePersonal(c *gin.Context, dto entity.UpdatePersonalDto)
	UpdatePersonalPassword(c *gin.Context, dto entity.UpdatePersonalPasswordDto)
}
type SysAdminServiceImpl struct{}

// 修改个人密码
func (s SysAdminServiceImpl) UpdatePersonalPassword(c *gin.Context, dto entity.UpdatePersonalPasswordDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, (result.ApiCode.MissingChangePasswordParameter),
			result.ApiCode.GetMessage(result.ApiCode.MissingChangePasswordParameter))
		return
	}
	sysAdmin, _ := jwt.GetAdmin(c)
	dto.Id = sysAdmin.ID
	sysAdminExist := dao.GetSysAdminByUsername(sysAdmin.Username)
	if sysAdminExist.Password != utils.EncryptionMd5(dto.Password) {
		result.Failed(c, (result.ApiCode.PASSWORDNOTTRUE),
			result.ApiCode.GetMessage(result.ApiCode.PASSWORDNOTTRUE))
		return
	}
	if dto.NewPassword != dto.ResetPassword {
		result.Failed(c, (result.ApiCode.RESETPASSWORD),
			result.ApiCode.GetMessage(result.ApiCode.RESETPASSWORD))
		return
	}
	dto.NewPassword = utils.EncryptionMd5(dto.NewPassword)
	sysAdminUpdatePwd := dao.UpdatePersonalPassword(dto)
	tokenString, _ := jwt.GenerateTokenByAdmin(sysAdminUpdatePwd)
	result.Success(c, map[string]interface{}{"token": tokenString, "sysAdmin": sysAdminUpdatePwd})
	return
}

// 修改个人信息
func (s SysAdminServiceImpl) UpdatePersonal(c *gin.Context, dto entity.UpdatePersonalDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c,
			(result.ApiCode.MissingModificationOfPersonalParameters),
			result.ApiCode.GetMessage(result.ApiCode.MissingModificationOfPersonalParameters))
		return
	}
	id, _ := jwt.GetAdminId(c)
	dto.Id = id
	result.Success(c, dao.UpdatePersonal(dto))
}

// 分页查询用户列表
func (s SysAdminServiceImpl) GetSysAdminList(c *gin.Context, PageSize, PageNum int, Username, Status, BeginTime, EndTime string) {
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	sysAdmin, count := dao.GetSysAdminList(PageSize, PageNum, Username, Status,
		BeginTime, EndTime)
	result.Success(c, map[string]interface{}{"total": count, "pageSize": PageSize,
		"pageNum": PageNum, "list": sysAdmin})
	return
}

// 重置密码
func (s SysAdminServiceImpl) ResetSysAdminPassword(c *gin.Context, dto entity.ResetSysAdminPasswordDto) {
	dao.ResetSysAdminPassword(dto)
	result.Success(c, true)
}

// 修改用户状态
func (s SysAdminServiceImpl) UpdateSysAdminStatus(c *gin.Context, dto entity.UpdateSysAdminStatusDto) {
	dao.UpdateSysAdminStatus(dto)
	result.Success(c, true)
}

// 根据id删除用户
func (s SysAdminServiceImpl) DeleteSysAdminById(c *gin.Context, dto entity.SysAdminIdDto) {
	dao.DeleteSysAdminById(dto)
	result.Success(c, true)
}

// 修改用户
func (s SysAdminServiceImpl) UpdateSysAdmin(c *gin.Context, dto entity.UpdateSysAdminDto) {
	result.Success(c, dao.UpdateSysAdmin(dto))
}

// 根据id查询用户信息
func (s SysAdminServiceImpl) GetSysAdminInfo(c *gin.Context, Id int) {
	result.Success(c, dao.GetSysAdminInfo(Id))
}

// 新增用户
func (s SysAdminServiceImpl) CreateSysAdmin(c *gin.Context, dto entity.AddSysAdminDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, (result.ApiCode.MissingNewAdminParameter),
			result.ApiCode.GetMessage(result.ApiCode.MissingNewAdminParameter))
		return
	}
	bool := dao.CreateSysAdmin(dto)
	if !bool {
		result.Failed(c, (result.ApiCode.USERNAMEALREADYEXISTS),
			result.ApiCode.GetMessage(result.ApiCode.USERNAMEALREADYEXISTS))
		return
	}
	result.Success(c, bool)
	return
}

// Login 登录
func (SysAdminServiceImpl) Login(c *gin.Context, dto entity.LoginDto) {
	// 登录参数校验
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, result.ApiCode.MissingLoginParameter, result.ApiCode.GetMessage(result.ApiCode.MissingLoginParameter))
		return
	}
	ip := c.ClientIP()
	// 校验验证码
	verifyRes := CaptVerify(dto.IdKey, dto.Image)
	if !verifyRes {
		dao.CreateSysLoginInfo(dto.Username, ip, utils.GetRealAddressByIP(ip), utils.GetBrowser(c), utils.GetOs(c), "验证码不正确", 2)
		result.Failed(c, (result.ApiCode.CAPTCHANOTTRUE), result.ApiCode.GetMessage(result.ApiCode.CAPTCHANOTTRUE))
		return
	}
	// 校验
	sysAdmin := dao.SysAdminDetail(dto)
	if sysAdmin.Password != utils.EncryptionMd5(dto.Password) {
		dao.CreateSysLoginInfo(dto.Username, ip, utils.GetRealAddressByIP(ip), utils.GetBrowser(c), utils.GetOs(c), "密码不正确", 2)
		result.Failed(c, (result.ApiCode.PASSWORDNOTTRUE), result.ApiCode.GetMessage(result.ApiCode.PASSWORDNOTTRUE))
		return
	}
	const status int = 2
	if sysAdmin.Status == uint(status) {
		dao.CreateSysLoginInfo(dto.Username, ip, utils.GetRealAddressByIP(ip), utils.GetBrowser(c), utils.GetOs(c), "账号已停用", 2)
		result.Failed(c, (result.ApiCode.STATUSISENABLE), result.ApiCode.GetMessage(result.ApiCode.STATUSISENABLE))
		return
	}
	// 生成token
	dao.CreateSysLoginInfo(dto.Username, ip, utils.GetRealAddressByIP(ip), utils.GetBrowser(c), utils.GetOs(c), "登录成功", 1)
	var leftMenuVo []entity.LeftMenuVo
	leftMenuList := dao.QueryLeftMenuList(sysAdmin.ID)
	for _, value := range leftMenuList {
		menuSvoList := dao.QueryMenuVoList(sysAdmin.ID, value.Id)
		item := entity.LeftMenuVo{}
		item.MenuSvoList = menuSvoList
		item.Id = value.Id
		item.MenuName = value.MenuName
		item.Icon = value.Icon
		item.Url = value.Url
		leftMenuVo = append(leftMenuVo, item)
	}
	permissionList := dao.QueryPermissionList(sysAdmin.ID)
	var stringList = make([]string, 0)
	for _, value := range permissionList {
		stringList = append(stringList, value.Value)
	}
	// 生成Token
	tokenString, _ := jwt.GenerateTokenByAdmin(sysAdmin)
	result.Success(c, map[string]interface{}{"token": tokenString, "sysAdmin": sysAdmin, "leftMenuList": leftMenuVo, "permissionList": stringList})
}

var sysAdminService = SysAdminServiceImpl{}

func SysAdminService() ISysAdminService {
	return &sysAdminService
}

//func SysAdminService() ISysAdminService {
//	return &sysAdminService
//}
