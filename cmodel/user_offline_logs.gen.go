package cmodel

const TableNameUserOfflineLog = "user_offline_logs"

// UserOfflineLog 用户断线记录表模型
type UserOfflineLog struct {
	ID        string `gorm:"type:uuid;default:uuid_generate_v4();" json:"id"`                       // UUID，使用 PostgreSQL 的 UUID 类型
	UserID    int64  `gorm:"not null;comment:用户ID" json:"user_id"`                                // 用户 ID
	EventTime int64  `gorm:"default:toUnixTimestamp(now());comment:事件发生时间" json:"event_time"` // 事件发生时间（Unix 时间戳）
	ClientIP  string `gorm:"not null;comment:客户端IP" json:"client_ip"`                            // 客户端 IP
	CreatedAt int64  `gorm:"default:toUnixTimestamp(now());comment:创建时间" json:"created_at"`     // 创建时间（Unix 时间戳）
}

// TableName 指定表名
func (*UserOfflineLog) TableName() string {
	return TableNameUserOfflineLog
}
