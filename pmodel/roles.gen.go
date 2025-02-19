// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package pmodel

import (
	"github.com/QuantumShiftX/golib/stores/gormx"
	"gorm.io/datatypes"
)

const TableNameRole = "roles"

// Role mapped from table <roles>
type Role struct {
	ID            int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"`              // 主键
	Name          string `gorm:"column:name;comment:角色名" json:"name"`                                       // 角色名
	Code          string `gorm:"column:code;comment:角色码" json:"code"`                                       // 角色码
	MenusID       datatypes.JSONSlice[int64] `gorm:"column:menus_id;comment:权限菜单ID数组" json:"menus_id"`                          // 权限菜单ID数组
	Status        int64  `gorm:"column:status;comment:状态：1-启用，2-禁用" json:"status"`                          // 状态：1-启用，2-禁用
	Comment       string `gorm:"column:comment;comment:备注" json:"comment"`                                  // 备注
	Sort          int64  `gorm:"column:sort;comment:排序" json:"sort"`                                        // 排序
	//
	gormx.OperationBaseModel
	gormx.Model
}

// TableName Role's table name
func (*Role) TableName() string {
	return TableNameRole
}
