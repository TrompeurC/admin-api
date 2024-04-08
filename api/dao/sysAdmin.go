package dao

import (
	"admin-api/api/entity"
	"admin-api/pkg/db"
)

// SysAdminDetail 用户详情
func SysAdminDetail(dto entity.LoginDto) entity.SysAdmin {
	var sysAdmin = entity.SysAdmin{}
	username := dto.Username
	db.Db.Where("username = ? ", username).First(&sysAdmin)
	return sysAdmin
}
