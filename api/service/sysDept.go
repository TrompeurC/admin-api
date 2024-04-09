package service

import (
	"admin-api/api/dao"
	"admin-api/api/entity"
	"admin-api/common/result"
	"github.com/gin-gonic/gin"
)

type ISysDeptService interface {
	CreateSysDept(c *gin.Context, sysDept entity.SysDept)
	QuerySysDeptVoList(c *gin.Context)
	GetSysDeptById(c *gin.Context, Id int)
	UpdateSysDept(c *gin.Context, dept entity.SysDept)
	DeleteSysDeptById(c *gin.Context, dto entity.SysDeptIdDto)
	GetSysDeptList(c *gin.Context, DeptName string, DeptStatus string)
}
type SysDeptServiceImpl struct{}

// CreateSysDept 创建部门
func (SysDeptServiceImpl) CreateSysDept(c *gin.Context, sysDept entity.SysDept) {
	isSuccess := dao.CreateSysDept(sysDept)
	if !isSuccess {
		result.Failed(c, result.ApiCode.DEPTISDISTRIBUTE, result.ApiCode.GetMessage(result.ApiCode.DEPTISDISTRIBUTE))
		return
	}
	result.Success(c, true)
}

// QuerySysDeptVoList 查询部门下拉列表
func (SysDeptServiceImpl) QuerySysDeptVoList(c *gin.Context) {
	vo := dao.QuerySysDeptVoList()
	result.Success(c, vo)
}

// GetSysDeptById 根据id查询部门
func (SysDeptServiceImpl) GetSysDeptById(c *gin.Context, Id int) {
	dept := dao.GetSysDeptById(Id)
	result.Success(c, dept)
}

// UpdateSysDept 修改部门
func (SysDeptServiceImpl) UpdateSysDept(c *gin.Context, dept entity.SysDept) {
	sysDept := dao.UpdateSysDept(dept)
	result.Success(c, sysDept)
}

// DeleteSysDeptById 删除部门
func (SysDeptServiceImpl) DeleteSysDeptById(c *gin.Context, dto entity.SysDeptIdDto) {
	dao.DeleteSysDeptById(dto)
	result.Success(c, true)
}

// GetSysDeptList 查询部门列表
func (SysDeptServiceImpl) GetSysDeptList(c *gin.Context, DeptName string, DeptStatus string) {
	depts := dao.GetSysDeptList(DeptName, DeptStatus)
	result.Success(c, depts)
}

var sysDeptService = SysDeptServiceImpl{}

func SysDeptService() ISysDeptService {
	return &sysDeptService
}
