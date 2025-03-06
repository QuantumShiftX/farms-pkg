package cmodel

const TableNameOperationLog = "operation_logs"

// OperationLog 对应 ClickHouse `operation_logs` 表
type OperationLog struct {
	OperationID string `gorm:"column:operation_id;primaryKey" json:"operation_id"` // 操作 ID（主键）
	UserID      int64  `gorm:"column:user_id;index" json:"user_id"`                // 用户 ID
	Status      string `gorm:"column:status" json:"status"`                        // 操作状态 // started, completed, failed
	EntityType  string `gorm:"column:entity_type;index" json:"entity_type"`        // 实体类型
	EntityData  string `gorm:"column:entity_data" json:"entity_data"`              // 实体数据
	ErrorMsg    string `gorm:"column:error_msg" json:"error_msg"`                  // 错误消息
	CreatedAt   int64  `gorm:"column:created_at" json:"created_at"`                // 创建时间（Unix 时间戳）
	UpdatedAt   int64  `gorm:"column:updated_at" json:"updated_at"`                // 更新时间（Unix 时间戳）
}

// TableName 设置 ClickHouse 表名
func (*OperationLog) TableName() string {
	return TableNameOperationLog
}
