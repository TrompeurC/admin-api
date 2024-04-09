package service

import (
	"admin-api/api/dao"
	"admin-api/api/entity"
	"admin-api/common/result"
	"github.com/gin-gonic/gin"
)

// ISysMenuService 接口定义
type ISysMenuService interface {
	QuerySysMenuVoList(c *gin.Context)
	CreateSysMenu(c *gin.Context, menu entity.SysMenu)
	GetSysMenu(c *gin.Context, Id int)
	UpdateSysMenu(c *gin.Context, menu entity.SysMenu)
	DeleteSysMenu(c *gin.Context, dto entity.SysMenuIdDto)
	GetSysMenuList(c *gin.Context, MenuName string, MenuStatus string)
}
type SysMenuServiceImpl struct{}

// 查询菜单列表
func (s SysMenuServiceImpl) GetSysMenuList(c *gin.Context, MenuName string, MenuStatus string) {
	result.Success(c, dao.GetSysMenuList(MenuName, MenuStatus))
}

// 删除菜单
func (s SysMenuServiceImpl) DeleteSysMenu(c *gin.Context, dto entity.SysMenuIdDto) {
	bool := dao.DeleteSysMenu(dto)
	if !bool {
		result.Failed(c, (result.ApiCode.DELSYSMENUFAILED),
			result.ApiCode.GetMessage(result.ApiCode.DELSYSMENUFAILED))
		return
	}
	result.Success(c, true)
}

// 修改菜单
func (s SysMenuServiceImpl) UpdateSysMenu(c *gin.Context, menu entity.SysMenu) {
	result.Success(c, dao.UpdateSysMenu(menu))
}

// 获取详情
func (s SysMenuServiceImpl) GetSysMenu(c *gin.Context, Id int) {
	result.Success(c, dao.GetSysMenu(Id))
}

// 新增菜单
func (s SysMenuServiceImpl) CreateSysMenu(c *gin.Context, sysMenu entity.SysMenu) {
	bool := dao.CreateSysMenu(sysMenu)
	if !bool {
		result.Failed(c, (result.ApiCode.MENUISEXIST),
			result.ApiCode.GetMessage(result.ApiCode.MENUISEXIST))
		return
	}
	result.Success(c, true)
}

// QuerySysMenuVoList 查询新增选项列表
func (s SysMenuServiceImpl) QuerySysMenuVoList(c *gin.Context) {
	result.Success(c, dao.QuerySysMenuVoList())
}

var sysMenuService = SysMenuServiceImpl{}

func SysMenuService() ISysMenuService {
	return &sysMenuService
}
