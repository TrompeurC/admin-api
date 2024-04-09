package service

import (
	"admin-api/api/dao"
	"admin-api/api/entity"
	"admin-api/common/result"
	"github.com/gin-gonic/gin"
)

type ISysPostService interface {
	CreateSysPost(c *gin.Context, sysPost entity.SysPost)
	GetSysPostById(c *gin.Context, Id uint)
	UpdateSysPost(c *gin.Context, sysPost entity.SysPost)
	DeleteSysPostById(c *gin.Context, dto entity.SysPostIdDto)
	BatchDeleteSysPost(c *gin.Context, dto entity.DelSysPostDto)
	UpdateSysPostStatus(c *gin.Context, dto entity.UpdateSysPostStatusDto)
	GetSysPostList(c *gin.Context, PageNum, PageSize uint, PostName, PostStatus,
		BeginTime, EndTime string)
	QuerySysPostVoList(c *gin.Context)
}

type SysPostServiceImpl struct {
}

// 岗位下拉列表
func (SysPostServiceImpl) QuerySysPostVoList(c *gin.Context) {
	result.Success(c, dao.QuerySysPostVoList())
}

// 分页查询岗位列表
func (SysPostServiceImpl) GetSysPostList(c *gin.Context, PageNum, PageSize uint, PostName, PostStatus,
	BeginTime, EndTime string) {
	if PageSize < 10 {
		PageSize = 10
	}
	if PageNum < 0 {
		PageNum = 1
	}
	total, data := dao.GetSysPostList(PageNum, PageSize, PostName, PostStatus, BeginTime, EndTime)
	result.Success(c, gin.H{
		"total":    total,
		"list":     data,
		"pageSize": PageSize,
		"pageNum":  PageNum,
	})
}

// 修改岗位状态
func (SysPostServiceImpl) UpdateSysPostStatus(c *gin.Context, dto entity.UpdateSysPostStatusDto) {
	dao.UpdateSysPostStatus(dto)
	result.Success(c, true)
}

// 根据id删除岗位
func (SysPostServiceImpl) DeleteSysPostById(c *gin.Context, dto entity.SysPostIdDto) {
	dao.DeleteSysPostById(dto)
	result.Success(c, true)
}

// 批量删除
func (s SysPostServiceImpl) BatchDeleteSysPost(c *gin.Context, dto entity.DelSysPostDto) {
	dao.BatchDeleteSysPost(dto)
	result.Success(c, true)
}

// 根据id查询岗位
func (s SysPostServiceImpl) GetSysPostById(c *gin.Context, Id uint) {
	result.Success(c, dao.GetSysPostById((Id)))
}

// 新增岗位
func (s SysPostServiceImpl) CreateSysPost(c *gin.Context, sysPost entity.SysPost) {
	bool := dao.CreateSysPost(&sysPost)
	if !bool {
		result.Failed(c, (result.ApiCode.POSTALREADYEXISTS),
			result.ApiCode.GetMessage(result.ApiCode.POSTALREADYEXISTS))
		return
	}
	result.Success(c, true)
}

// 修改岗位
func (s SysPostServiceImpl) UpdateSysPost(c *gin.Context, sysPost entity.SysPost) {
	result.Success(c, dao.UpdateSysPost(&sysPost))
}

var sysPostService = SysPostServiceImpl{}

func SysPostService() ISysPostService {
	return &sysPostService
}
