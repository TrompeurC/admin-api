package entity

import "admin-api/utils"

type SysMenu struct {
	ID         uint        `json:"id" gorm:"column:id;primarykey;autoIncrement;comment:菜单ID"`
	MenuName   string      `json:"name" gorm:"column:name;varchar(100);comment:菜单名称"`
	ParentId   uint        `json:"parentId" gorm:"column:parent_id;comment:父菜单ID"`
	Url        string      `json:"path" gorm:"column:url;varchar(100);comment:路由Url"`
	Icon       string      `json:"icon" gorm:"column:icon;varchar(100);comment:菜单图标"`
	Value      string      `json:"value" gorm:"column:value;varchar(100);comment:权限值"`
	MenuType   uint        `json:"type" gorm:"column:type;comment:菜单类型:1->目录;2->菜单;3->按钮"`        // 菜单类型:1->目录;2->菜单;3->按钮
	MenuStatus uint        `json:"menuStatus" gorm:"column:menu_status;comment:菜单状态:1->禁用;2->启用"` // 菜单状态:1->启用;0->禁用
	CreateTime utils.HTime `json:"createTime" gorm:"column:create_time;autoCreateTime;comment:创建时间"`
	Sort       uint        `json:"sort" gorm:"column:sort;comment:排序"`
	Children   []SysMenu   `json:"children" gorm:"-"`
}

// Id参数
type SysMenuIdDto struct {
	Id uint `json:"id"` // ID
}

// SysMenuVo 对象
type SysMenuVo struct {
	Id       uint   `json:"id"`
	ParentId uint   `json:"parentId"` // 父id
	Label    string `json:"label"`    // 名称
}
