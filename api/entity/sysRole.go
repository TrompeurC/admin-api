package entity

import "admin-api/utils"

type SysRole struct {
	ID          uint        `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	RoleName    string      `json:"roleName" gorm:"role_name;varchar(100);comment:角色名称"`
	RoleKey     string      `json:"roleKey" gorm:"role_key;varchar(100);comment:角色权限字符串"`
	Status      uint        `json:"status" gorm:"status;comment:角色状态:1->启用;2->禁用"`
	Description string      `json:"description" gorm:"description;varchar(500);comment:描述"`
	CreateTime  utils.HTime `json:"createTime" gorm:"column:create_time;autoCreateTime;comment:创建时间"`
}

// 新增参数
type AddSysRoleDto struct {
	RoleName    string // 角色名称
	RoleKey     string // 角色key
	Status      uint   // 状态:1->启用,2->禁用
	Description string // 描述
}

// Id参数
type SysRoleIdDto struct {
	Id uint `json:"id"` // ID
}

// 新增参数
type UpdateSysRoleDto struct {
	Id          uint   // Id
	RoleName    string // 角色名称
	RoleKey     string // 角色key
	Status      int    // 状态:1->启用,2->禁用
	Description string // 描述
}

// UpdateSysRoleStatusDto 设置状态参数
type UpdateSysRoleStatusDto struct {
	Id     uint // ID
	Status int  // 状态:1->启用,2->禁用
}

// 角色下拉列表
type SysRoleVo struct {
	Id       int    `json:"id"`       // ID
	RoleName string `json:"roleName"` // 角色名称
}

// IdVo 当前角色的菜单权限id
type IdVo struct {
	Id int `json:"id"` // ID
}

// RoleMenu 角色id,菜单id视图
type RoleMenu struct {
	Id      uint   `json:"id" binding:"required"`      // ID
	MenuIds []uint `json:"menuIds" binding:"required"` // 菜单id列表
}
