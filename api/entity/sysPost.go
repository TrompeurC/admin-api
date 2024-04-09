package entity

import "admin-api/utils"

type SysPost struct {
	ID         uint        `json:"id" gorm:"column:id;primaryKey;autoIncrement;comment:岗位ID"`
	PostCode   string      `json:"postCode" gorm:"column:post_code;varchar(64);comment:岗位编码"`
	PostName   string      `json:"postName" gorm:"column:post_name;varchar(128);comment:岗位名称"`
	PostStatus uint        `json:"postStatus" gorm:"column:post_status;default:1;comment:''状态(1->正 常 2->停用)"`
	CreateTime utils.HTime `json:"createTime" gorm:"column:create_time;comment:'创建时间'"`
	Remark     string      `json:"remark" gorm:"column:remark;varchar(200)"`
}

// 修改状态参数
type UpdateSysPostStatusDto struct {
	Id         uint // ID
	PostStatus uint // 状态(1->正常 2->停用)
}

// Id参数
type SysPostIdDto struct {
	Id uint `json:"id"` // ID
}

// 删除岗位的参数
type DelSysPostDto struct {
	Ids []uint // Id列表
}

// 岗位下拉列表对象模型
type SysPostVo struct {
	Id       int    `json:"id"`       // ID
	PostName string `json:"postName"` // 岗位名称
}
