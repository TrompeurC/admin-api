package dao

import (
	"admin-api/api/entity"
	"admin-api/pkg/db"
	"admin-api/utils"
	"time"
)

// 根据部门名称查询
func GetSysDeptByName(deptName string) (sysDept entity.SysDept) {
	db.Db.Where("dept_name = ?", deptName).First(&sysDept)
	return sysDept
}

// 新增部门
func CreateSysDept(sysDept entity.SysDept) bool {
	sysDeptByName := GetSysDeptByName(sysDept.DeptName)
	if sysDeptByName.ID > 0 {
		return false
	}
	if res := GetSysDeptByName(sysDept.DeptName); res.ID > 0 {
		return false
	}
	sysDept.CreateTime = utils.HTime{Time: time.Now()}
	tx := db.Db.Save(&sysDept)
	return tx.RowsAffected > 0
}

// 部门下拉列表
func QuerySysDeptVoList() (sysDeptVo []entity.SysDeptVo) {
	db.Db.Model(&entity.SysDept{}).Select("id, parent_id, dept_name").Scan(&sysDeptVo)
	return sysDeptVo
}

// 根据id查询部门
func GetSysDeptById(Id int) (sysDept entity.SysDept) {
	db.Db.Where("id = ?", Id).First(&sysDept)
	return sysDept
}

// 修改部门
func UpdateSysDept(dept entity.SysDept) (sysDept entity.SysDept) {
	db.Db.Save(&dept)
	return dept
}

// 查询部门是否有人
func GetSysAdminDept(id int) (sysAdmin entity.SysAdmin) {
	db.Db.Where("dept_id = ?", id).First(&sysAdmin)
	return sysAdmin
}

// 删除部门
func DeleteSysDeptById(dto entity.SysDeptIdDto) bool {
	sysAdmin := GetSysAdminDept(dto.Id)
	if sysAdmin.ID > 0 {
		return false
	}
	db.Db.Where("parent_id = ? ", dto.Id).Delete(&entity.SysDept{})
	db.Db.Delete(&entity.SysDept{}, dto.Id)
	return true
}

// 查询部门列表
func GetSysDeptList(DeptName string, DeptStatus string) (sysDept []entity.SysDept) {
	curDb := db.Db.Table("sys_dept")
	if DeptName != "" {
		curDb = curDb.Where("dept_name = ?", DeptName)
	}
	if DeptStatus != "" {
		curDb = curDb.Where("dept_status = ?", DeptStatus)
	}
	curDb.Find(&sysDept)
	return sysDept

}
