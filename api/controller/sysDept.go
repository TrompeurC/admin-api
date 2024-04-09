package controller

import (
	"admin-api/api/entity"
	"admin-api/api/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

var sysDept entity.SysDept

// 新增部门
// @Summary 新增部门接口
// @Produce json
// @Description 新增部门接口
// @Param data body entity.SysDept true "data" // @Success 200 {object} result.Result
// @router /api/dept/add [post]
func CreateSysDept(c *gin.Context) {
	_ = c.BindJSON(&sysDept)
	service.SysDeptService().CreateSysDept(c, sysDept)
}

// 部门下拉列表
// @Summary 部门下拉列表接口
// @Produce json
// @Description 部门下拉列表接口
// @Success 200 {object} result.Result
// @router /api/dept/vo/list [get]
func QuerySysDeptVoList(c *gin.Context) {
	service.SysDeptService().QuerySysDeptVoList(c)
}

// 根据id查询部门
// @Summary 根据id查询部门接口
// @Produce json
// @Description 根据id查询部门接口
// @Param id query int true "ID"
// @Success 200 {object} result.Result // @router /api/dept/info [get]
func GetSysDeptById(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	service.SysDeptService().GetSysDeptById(c, Id)
}

// 修改部门
// @Summary 修改部门接口
// @Produce json
// @Description 修改部门接口
// @Param data body entity.SysDept true "data" // @Success 200 {object} result.Result
// @router /api/dept/update [put]
func UpdateSysDept(c *gin.Context) {
	_ = c.BindJSON(&sysDept)
	service.SysDeptService().UpdateSysDept(c, sysDept)
}

// 根据id删除部门
// @Summary 根据id删除部门接口
// @Produce json
// @Description 根据id删除部门接口
// @Param data body entity.SysDeptIdDto true "data" // @Success 200 {object} result.Result
// @router /api/dept/delete [delete]
func DeleteSysDeptById(c *gin.Context) {
	var dto entity.SysDeptIdDto
	_ = c.BindJSON(&dto)
	service.SysDeptService().DeleteSysDeptById(c, dto)
}

// 查询部门列表
// @Summary 查询部门列表接口
// @Produce json
// @Description 查询部门列表接口
// @Param deptName query string false "部门名称"
// @Param deptStatus query string false "部门状态"
// @Success 200 {object} result.Result
// @router /api/dept/list [get]
func GetSysDeptList(c *gin.Context) {
	DeptName := c.Query("deptName")
	DeptStatus := c.Query("deptStatus")
	service.SysDeptService().GetSysDeptList(c, DeptName, DeptStatus)
}
