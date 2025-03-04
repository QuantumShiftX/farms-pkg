package cmodel

import (
	"time"
)

const TableNameVDailyOperatedUserSummary = "v_daily_operated_user_summaries"

// VDailyOperatedUserSummary 对应视图 `farms-game`.daily_operated_user_summary
// 统计被操作用户每天各类型被操作的总次数
type VDailyOperatedUserSummary struct {
	// 操作日期
	OperationDate time.Time `gorm:"column:operation_date;type:date;not null;primaryKey" json:"operation_date"`

	// 被操作的用户ID
	OpToUserID int64 `gorm:"column:op_to_user_id;not null;primaryKey" json:"op_to_user_id"`

	// 被操作的农场ID
	OpToFarmID int64 `gorm:"column:op_to_farm_id;not null;primaryKey" json:"op_to_farm_id"`

	// 操作类型：1-浇水 2-施肥
	OperationType int8 `gorm:"column:operation_type;not null;primaryKey" json:"operation_type"`

	// 被操作总次数
	TotalReceivedOperations int64 `gorm:"column:total_received_operations;not null" json:"total_received_operations"`

	// 操作者的不同用户数量
	UniqueOperators int64 `gorm:"column:unique_operators;not null" json:"unique_operators"`

	// 被操作的不同植物数量
	UniquePlants int64 `gorm:"column:unique_plants;not null" json:"unique_plants"`
}

// TableName 设置表名
func (*VDailyOperatedUserSummary) TableName() string {
	return TableNameVDailyOperatedUserSummary
}
