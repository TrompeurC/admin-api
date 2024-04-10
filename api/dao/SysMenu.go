package dao

import (
	"admin-api/api/entity"
	"admin-api/pkg/db"
)

// 根据菜单名称查询
func GetSysMenuByName(menuName string) (sysMenu entity.SysMenu) {
	db.Db.Where("menu_name = ?", menuName).First(&sysMenu)
	return sysMenu
}

// 查询新增选项列表
func QuerySysMenuVoList() (sysMenuVo []entity.SysMenuVo) {
	db.Db.Model(&entity.SysMenu{}).Select("id , menu_name as label ,parent_id ").Find(&sysMenuVo)
	return sysMenuVo
}

// 新增菜单
func CreateSysMenu(addSysMenu entity.SysMenu) bool {
	tx := db.Db.Create(&addSysMenu)
	return tx.RowsAffected > 0
}

// 根据id查询菜单详情
func GetSysMenu(Id int) (sysMenu entity.SysMenu) {
	db.Db.First(&sysMenu, Id)
	return sysMenu
}

// 修改菜单
func UpdateSysMenu(menu entity.SysMenu) (sysMenu entity.SysMenu) {
	db.Db.First(&sysMenu, menu.ID)
	sysMenu.ParentId = menu.ParentId
	sysMenu.MenuName = menu.MenuName
	sysMenu.Icon = menu.Icon
	sysMenu.Value = menu.Value
	sysMenu.MenuType = menu.MenuType
	sysMenu.Url = menu.Url
	sysMenu.MenuStatus = menu.MenuStatus
	sysMenu.Sort = menu.Sort
	db.Db.Save(&sysMenu)
	return sysMenu
}

func GetSysRoleMenu(id int) (sysRoleMenu entity.SysRoleMenu) {
	db.Db.Where("menu_id = ?", id).First(&sysRoleMenu)
	return sysRoleMenu
}

// 删除菜单
func DeleteSysMenu(dto entity.SysMenuIdDto) bool {
	// 菜单已分配角色不能删除
	sysRoleMenu := GetSysRoleMenu(int(dto.Id))
	if sysRoleMenu.MenuId > 0 {
		return false
	}
	db.Db.Where("parent_id = ?", dto.Id).Delete(&entity.SysMenu{})
	db.Db.Delete(&entity.SysMenu{}, dto.Id)
	return true
}

// 查询菜单列表
func GetSysMenuList(MenuName string, MenuStatus string) (sysMenu []*entity.SysMenu) {
	curDb := db.Db.Table("sys_menu").Order("sort")
	if MenuName != "" {
		curDb = curDb.Where("menu_name = ?", MenuName)
	}
	if MenuStatus != "" {
		curDb = curDb.Where("menu_status = ?", MenuStatus)
	}
	curDb.Find(&sysMenu)
	return sysMenu
}

// 当前登录用户左侧菜单级列表
func QueryMenuVoList(AdminId, MenuId uint) (menuSvo []entity.MenuSvo) {
	const status, menuStatus, menuType uint = 1, 2, 2
	db.Db.Table("sys_menu sm").
		Select("sm.menu_name, sm.icon, sm.url").
		Joins("LEFT JOIN sys_role_menu srm ON sm.id = srm.menu_id").
		Joins("LEFT JOIN sys_role sr ON sr.id = srm.role_id").
		Joins("LEFT JOIN sys_admin_role sar ON sar.role_id = sr.id").
		Joins("LEFT JOIN sys_admin sa ON sa.id = sar.admin_id").
		Where("sr.status = ?", status).
		Where("sm.menu_status = ?", menuStatus).
		Where("sm.menu_type = ?", menuType).
		Where("sm.parent_id = ?", MenuId).
		Where("sa.id = ?", AdminId).
		Order("sm.sort").
		Scan(&menuSvo)
	return menuSvo
}

// 当前登录用户左侧菜单列表
func QueryLeftMenuList(Id uint) (leftMenuVo []entity.LeftMenuVo) {
	const status, menuStatus, menuType uint = 1, 2, 1
	db.Db.Table("sys_menu sm").
		Select("sm.id, sm.menu_name, sm.url, sm.icon").
		Joins("LEFT JOIN sys_role_menu srm ON sm.id = srm.menu_id").
		Joins("LEFT JOIN sys_role sr ON sr.id = srm.role_id").
		Joins("LEFT JOIN sys_admin_role sar ON sar.role_id = sr.id").
		Joins("LEFT JOIN sys_admin sa ON sa.id = sar.admin_id").
		Where("sr.status = ?", status).
		Where("sm.menu_status = ?", menuStatus).
		Where("sm.menu_type = ?", menuType).
		Where("sa.id = ?", Id).
		Order("sm.sort").
		Scan(&leftMenuVo)
	return leftMenuVo
}

// 当前登录用户权限列表
func QueryPermissionList(Id uint) (valueVo []entity.ValueVo) {
	const status, menuStatus, menuType uint = 1, 2, 1
	db.Db.Table("sys_menu sm").
		Select("sm.value").
		Joins("LEFT JOIN sys_role_menu srm ON sm.id = srm.menu_id").
		Joins("LEFT JOIN sys_role sr ON sr.id = srm.role_id").
		Joins("LEFT JOIN sys_admin_role sar ON sar.role_id = sr.id").
		Joins("LEFT JOIN sys_admin sa ON sa.id = sar.admin_id").
		Where("sr.status = ?", status).
		Where("sm.menu_status = ?", menuStatus).
		Not("sm.menu_type = ?", menuType).
		Where("sa.id = ?", Id).
		Scan(&valueVo)
	return valueVo
}
