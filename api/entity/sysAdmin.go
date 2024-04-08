package entity

import "admin-api/utils"

type SysAdmin struct {
	ID         uint        `json:"id" gorm:"column:id;comment:'主键';primaryKey;autoIncrement"`
	PostId     uint        `json:"postId" gorm:"column:post_id;comment:'岗位id'"`
	DeptId     uint        `json:"dept_id" gorm:"column:dept_id;comment:'部门id'"`
	Username   string      `json:"username" gorm:"column:username;varchar(64);NOT NULL;comment:'用户名'"`
	Password   string      `json:"password" gorm:"column:password;varchar(64);NOT NULL;comment:'密码'"`
	Nickname   string      `json:"nickname" gorm:"column:nickname;varchar(64);comment:'昵称'"`
	Status     string      `json:"status" gorm:"column:status;default:1;NOT NULL;comment:'1->启用,2->禁用';"`
	Icon       string      `json:"icon" gorm:"column:icon;varchar(500);comment:'头像'"`
	Email      string      `json:"email" gorm:"column:email;varchar(64);comment:'邮箱'"`
	Phone      string      `json:"phone" gorm:"column:phone;varchar(64);comment:'手机号'"`
	Note       string      `json:"note" gorm:"column:not;varchar(500);comment:'备注'"`
	CreateTime utils.HTime `json:"create_time" gorm:"column:create_time;autoCreateTime;comment:'创建时间'"`
}

func TableName() string {
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
	Image    string `json:"image" validate:"required"`
	IdKey    string `json:"idKey" validate:"required;min:4;max:6;"`
}
