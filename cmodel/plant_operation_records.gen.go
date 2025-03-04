package cmodel

import (
	"time"
)

const TableNamePlantOperationRecord = "plant_operation_records"

// PlantOperationRecord 对应ClickHouse表`farms-game`.plant_operation_records
// 种植物操作记录表(浇水、施肥等)
type PlantOperationRecord struct {
	// UUID 主键
	ID string `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`

	// 用户ID（操作人用户ID）
	UserID int64 `gorm:"column:user_id;not null;index:idx_user_id" json:"user_id"`

	// 农场ID（操作人农场ID）
	FarmID int64 `gorm:"column:farm_id;not null;index:idx_farm_id" json:"farm_id"`

	// 被操作的用户农场ID
	OpToFarmID int64 `gorm:"column:op_to_farm_id;not null;index:idx_op_to_farm_id" json:"op_to_farm_id"`

	// 被操作的用户种植记录ID
	OpToPlantID int64 `gorm:"column:op_to_plant_id;not null;index:idx_op_to_plant_id" json:"op_to_plant_id"`

	// 被操作的用户ID
	OpToUserID int64 `gorm:"column:op_to_user_id;not null;index:idx_op_to_user_id" json:"op_to_user_id"`

	// 操作类型：1-浇水 2-施肥
	OperationType int8 `gorm:"column:operation_type;not null;index:idx_operation_type" json:"operation_type"`

	// 操作日期
	OperationDate time.Time `gorm:"column:operation_date;type:date;not null;default:CURRENT_DATE;index:idx_operation_date" json:"operation_date"`

	// 操作时间（Unix时间戳）
	OperatedAt int64 `gorm:"column:operated_at;not null;default:extract(epoch from now())" json:"operated_at"`

	// 操作时间（转换为Go时间类型，用于便捷操作）
	OperatedAtTime time.Time `gorm:"-" json:"operated_at_time"`
}

// TableName 设置表名
func (*PlantOperationRecord) TableName() string {
	return TableNamePlantOperationRecord
}
