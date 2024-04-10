package entity

import "admin-api/utils"

type SysAdmin struct {
	ID         uint        `json:"id" gorm:"column:id;comment:'主键';primaryKey;autoIncrement"`
	PostId     uint        `json:"postId" gorm:"column:post_id;comment:'岗位id'"`
	DeptId     uint        `json:"dept_id" gorm:"column:dept_id;comment:'部门id'"`
	Username   string      `json:"username" gorm:"column:username;varchar(64);NOT NULL;comment:'用户名'"`
	Password   string      `json:"password" gorm:"column:password;varchar(64);NOT NULL;comment:'密码'"`
	Nickname   string      `json:"nickname" gorm:"column:nickname;varchar(64);comment:'昵称'"`
	Status     uint        `json:"status" gorm:"column:status;default:1;NOT NULL;comment:'1->启用,2->禁用';"`
	Icon       string      `json:"icon" gorm:"column:icon;varchar(500);comment:'头像'"`
	Email      string      `json:"email" gorm:"column:email;varchar(64);comment:'邮箱'"`
	Phone      string      `json:"phone" gorm:"column:phone;varchar(64);comment:'手机号'"`
	Note       string      `json:"note" gorm:"column:not;varchar(500);comment:'备注'"`
	CreateTime utils.HTime `json:"create_time" gorm:"column:create_time;autoCreateTime;comment:'创建时间'"`
}

func (SysAdmin) TableName() string {
	return "sys_admin"
}

// 鉴权用户结构体
type JwtAdmin struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Icon     string `json:"icon"`
	Phone    string `json:"phone"`
	Note     string `json:"note"`
}

// 登录对象
type LoginDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Image    string `json:"image" validate:"required,min=4,max=6"`
	IdKey    string `json:"idKey" validate:"required"`
}

// AddSysAdminDto 新增参数
type AddSysAdminDto struct {
	PostId   uint   `validate:"required"` // 岗位id
	RoleId   uint   `validate:"required"` // 角色id
	DeptId   uint   `validate:"required"` // 部门id
	Username string `validate:"required"` // 用户名
	Password string `validate:"required"` // 密码
	Nickname string `validate:"required"` // 昵称
	Phone    string `validate:"required"` // 手机号
	Email    string `validate:"required"` // 邮箱
	Note     string // 备注
	Status   uint   `validate:"required"` // 状态：1->启用,2->禁用
}

// 详情视图
type SysAdminInfo struct {
	ID       uint   `json:"id"`       // ID
	Username string `json:"username"` // 用户名
	Nickname string `json:"nickname"` // 昵称
	Status   uint   `json:"status"`   // 状态：1->启用,2->禁用
	PostId   uint   `json:"postId"`   // 岗位id
	DeptId   uint   `json:"deptId"`   // 部门id
	RoleId   uint   `json:"roleId" `  // 角色id
	Email    string `json:"email"`    // 邮箱
	Phone    string `json:"phone"`    // 手机号
	Note     string `json:"note"`     // 备注
}

// 修改参数
type UpdateSysAdminDto struct {
	Id       uint   // ID
	PostId   uint   // 岗位id
	DeptId   uint   // 部门id
	RoleId   uint   // 角色id
	Username string // 用户名
	Nickname string // 昵称
	Phone    string // 手机号
	Email    string // 邮箱
	Note     string // 备注
	Status   uint   // 状态：1->启用,2->禁用
}

// Id参数
type SysAdminIdDto struct {
	Id uint `json:"id"` // ID
}

// 设置状态参数
type UpdateSysAdminStatusDto struct {
	Id     uint // ID
	Status int  // 状态：1->启用,2->禁用
}

// 重置密码参数
type ResetSysAdminPasswordDto struct {
	Id       uint   // ID
	Password string //密码
}

// 用户列表的vo视图
type SysAdminVo struct {
	ID         uint        `json:"id"`         // ID
	Username   string      `json:"username"`   // 用户名
	Nickname   string      `json:"nickname"`   // 昵称
	Status     int         `json:"status"`     // 状态：1->启用,2->禁用
	PostId     int         `json:"postId"`     // 岗位id
	DeptId     int         `json:"deptId"`     // 部门id
	RoleId     uint        `json:"roleId" `    // 角色id
	PostName   string      `json:"postName"`   // 岗位名称
	DeptName   string      `json:"deptName"`   // 部门名称
	RoleName   string      `json:"roleName"`   // 角色名称
	Icon       string      `json:"icon"`       // 头像
	Email      string      `json:"email"`      // 邮箱
	Phone      string      `json:"phone"`      // 电话
	Note       string      `json:"note"`       // 备注
	CreateTime utils.HTime `json:"createTime"` // 创建时间
}

// 修改个人信息参数
type UpdatePersonalDto struct {
	Id       uint   //ID
	Icon     string // 头像
	Username string `validate:"required"` //用户名
	Nickname string `validate:"required"` // 昵称
	Phone    string `validate:"required"` // 电话
	Email    string `validate:"required"` // 邮箱
	Note     string `validate:"required"` // 备注
}

// 修改个人密码
type UpdatePersonalPasswordDto struct {
	Id            uint   //ID
	Password      string `validate:"required"` // 密码
	NewPassword   string `validate:"required"` // 新密码
	ResetPassword string `validate:"required"` // 重复密码
}
