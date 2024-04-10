package dao

import (
	"admin-api/api/entity"
	"admin-api/pkg/db"
	"admin-api/utils"
	"time"
)

// 新增登录日志
func CreateSysLoginInfo(username, ipAddress, loginLocation, browser, os, message string, loginStatus int) {
	sysLoginInfo := entity.SysLoginInfo{
		Username:      username,
		IpAddress:     ipAddress,
		LoginLocation: loginLocation,
		Browser:       browser,
		Os:            os,
		Message:       message,
		LoginStatus:   loginStatus,
		LoginTime:     utils.HTime{Time: time.Now()},
	}
	db.Db.Save(&sysLoginInfo)
}

// 分页获取登录日志列表
func GetSysLoginInfoList(Username, LoginStatus, BeginTime, EndTime string, PageSize, PageNum int) (sysLoginInfo []entity.SysLoginInfo, count int64) {
	curDb := db.Db.Table("sys_login_info")
	if Username != "" {
		curDb = curDb.Where("username = ?", Username)
	}
	if BeginTime != "" && EndTime != "" {
		curDb = curDb.Where("`login_time` BETWEEN ? AND ?", BeginTime, EndTime)
	}
	if LoginStatus != "" {
		curDb = curDb.Where("login_status = ?", LoginStatus)
	}
	curDb.Count(&count)
	curDb.Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("login_time desc").Find(&sysLoginInfo)
	return sysLoginInfo, count
}

// 批量删除登录日志
func BatchDeleteSysLoginInfo(dto entity.DelSysLoginInfoDto) {
	db.Db.Where("id in (?)", dto.Ids).Delete(&entity.SysLoginInfo{})
}

// 根据id删除日志
func DeleteSysLoginInfoById(dto entity.SysLoginInfoIdDto) {
	db.Db.Delete(&entity.SysLoginInfo{}, dto.Id)
}

// 清空登录日志
func CleanSysLoginInfo() {
	db.Db.Exec("truncate table sys_login_Info")
}
