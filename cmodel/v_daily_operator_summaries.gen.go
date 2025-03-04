package cmodel

import (
	"time"
)

const TableNameVDailyOperatorSummary = "v_daily_operator_summaries"

// VDailyOperatorSummary 对应视图 `farms-game`.daily_operator_summary
// 统计操作人每天各类型操作的总次数
type VDailyOperatorSummary struct {
	// 操作日期
	OperationDate time.Time `gorm:"column:operation_date;type:date;not null;primaryKey" json:"operation_date"`

	// 操作人用户ID
	UserID int64 `gorm:"column:user_id;not null;primaryKey" json:"user_id"`

	// 操作人农场ID
	FarmID int64 `gorm:"column:farm_id;not null;primaryKey" json:"farm_id"`

	// 操作类型：1-浇水 2-施肥
	OperationType int8 `gorm:"column:operation_type;not null;primaryKey" json:"operation_type"`

	// 操作总次数
	TotalOperations int64 `gorm:"column:total_operations;not null" json:"total_operations"`

	// 操作的不同植物数量
	UniquePlants int64 `gorm:"column:unique_plants;not null" json:"unique_plants"`

	// 操作的不同用户数量
	UniqueUsers int64 `gorm:"column:unique_users;not null" json:"unique_users"`
}

// TableName 设置表名
func (*VDailyOperatorSummary) TableName() string {
	return TableNameVDailyOperatorSummary
}
