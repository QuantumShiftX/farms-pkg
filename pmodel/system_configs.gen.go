// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package pmodel

import "github.com/QuantumShiftX/golib/stores/gormx"

const TableNameSystemConfig = "system_configs"

// SystemConfig mapped from table <system_configs>
type SystemConfig struct {
	ID            int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"`        // 主键
	Type          string `gorm:"column:type;not null;comment:类型(如int、string、json)" json:"type"`       // 类型(如int、string、json)
	Category      string `gorm:"column:category;not null;comment:分组" json:"category"`                 // 分组
	Key           string `gorm:"column:key;not null;comment:键" json:"key"`                            // 键
	Val           string `gorm:"column:val;not null;comment:值" json:"val"`                            // 值
	Status        int8  `gorm:"column:status;not null;default:1;comment:状态：1-启用；2-禁用" json:"status"` // 状态：1-启用；2-禁用
	Metadata      string `gorm:"column:metadata;comment:元数据（额外存储信息，如默认值等）" json:"metadata"`           // 元数据（额外存储信息，如默认值等）
	Rank          int64  `gorm:"column:rank;not null;comment:排序" json:"rank"`                         // 排序
	Remark        string `gorm:"column:remark;not null;comment:备注" json:"remark"`                     // 备注
	//
	gormx.OperationBaseModel
	gormx.Model

}

// TableName SystemConfig's table name
func (*SystemConfig) TableName() string {
	return TableNameSystemConfig
}
