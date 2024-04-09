package dao

import (
	"admin-api/api/entity"
	"admin-api/pkg/db"
	"admin-api/utils"
	"time"
)

// 根据编码查询
func GetSysPostByCode(code string) (sysPost *entity.SysPost) {
	db.Db.Where("post_code = ?", code).First(&sysPost)
	return sysPost
}

// 根据名称查询
func GetSysPostByName(name string) (sysPost *entity.SysPost) {
	db.Db.Where("post_name = ?", name).First(&sysPost)
	return sysPost
}

// 新增岗位
func CreateSysPost(sysPost *entity.SysPost) bool {
	sysPostCode := GetSysPostByCode(sysPost.PostCode)
	if sysPostCode.ID > 0 {
		return false
	}
	sysPostName := GetSysPostByName(sysPost.PostName)
	if sysPostName.ID > 0 {
		return false
	}
	sysPost.CreateTime = utils.HTime{Time: time.Now()}
	tx := db.Db.Create(sysPost)

	return tx.RowsAffected > 0
}

// 根据id查询岗位
func GetSysPostById(id uint) (sysPost *entity.SysPost) {
	db.Db.First(sysPost, id)
	return sysPost
}

// 修改岗位
func UpdateSysPost(post *entity.SysPost) (sysPost *entity.SysPost) {
	db.Db.Save(post)
	return post
}

// 根据id删除岗位
func DeleteSysPostById(dto entity.SysPostIdDto) {
	db.Db.Delete(&entity.SysPost{}, dto.Id)
}

// 批量删除岗位
func BatchDeleteSysPost(dto entity.DelSysPostDto) {
	db.Db.Where("id in (?)", dto.Ids).Delete(&entity.SysPost{})
}

// 修改状态
func UpdateSysPostStatus(dto entity.UpdateSysPostStatusDto) {
	var sysPost entity.SysPost
	db.Db.First(&sysPost, dto.Id)
	sysPost.PostStatus = dto.PostStatus
	db.Db.Save(&sysPost)
}

// 分页查询岗位列表
func GetSysPostList(PageNum, PageSize uint, PostName, PostStatus, BeginTime, EndTime string) (sysPost []entity.SysPost, count int64) {
	curDb := db.Db.Table("sys_post")
	if PostName != "" {
		curDb.Where("post_name like ?", PostName)
	}
	if PostStatus != "" {
		curDb = curDb.Where("post_status = ?", PostStatus)
	}
	if BeginTime != "" && EndTime != "" {
		curDb.Where("create_time BETWEEN ? AND ? ", BeginTime, EndTime)
	}
	curDb.Count(&count)
	curDb.Limit(int(PageSize)).Offset(int(PageSize) * (int(PageNum) - 1)).Order("create_time desc").Find(&sysPost)
	return sysPost, count
}

// 岗位下拉列表
func QuerySysPostVoList() (sysPostVo []entity.SysPostVo) {
	db.Db.Table("sys_post").Select("id , post_name").First(&sysPostVo)
	return sysPostVo
}
