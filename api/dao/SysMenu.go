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
