package entity

type SysRoleMenu struct {
	RoleId int `json:"roleId" gorm:"column:role_id;comment:角色ID"`
	MenuId int `json:"menuId" gorm:"column:menu_id;comment:菜单ID"`
}
