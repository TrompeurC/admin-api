package entity

import "admin-api/utils"

type SysDept struct {
	ID         uint        `json:"id" gorm:"column:id;primaryKey;autoIncrement;comment:'部门ID'"`
	ParentID   uint        `json:"parentId" gorm:"column:parent_id;comment:'父部门ID'"`
	DeptType   uint        `json:"deptType" gorm:"column:dept_type;comment:'部门类型(1：公司 2：中心，3：部门)'"`
	DeptName   string      `json:"deptName" gorm:"column:dept_name;comment:'部门名称';varchar(30);"`
	DeptStatus uint        `json:"deptStatus" gorm:"column:dept_status;comment:'部门状态(1:正常 2:停用)';default:1"`
	CreateTime utils.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"` // 创建时间
	Children   []SysDept   `json:"children" gorm:"-"`
}

// Id参数
type SysDeptIdDto struct {
	Id int `json:"id"` // ID
}

// 部门名称对象
type SysDeptVo struct {
	Id       uint `json:"id"`
	ParentId uint `json:"parentId"` // 父id Label string `json:"label"` // 名称
}
