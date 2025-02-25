package cmodel

import (
	"github.com/google/uuid"
	"time"
)

const TableNameUserOnlineDurationFlow = "user_online_duration_flows"

// UserOnlineDurationFlow 用户在线时长流水表模型
type UserOnlineDurationFlow struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"` // UUID，使用 PostgreSQL 的 UUID 类型
	UserID    int64     `gorm:"not null;comment:用户ID" json:"user_id"`           // 用户 ID
	StatDate  time.Time `gorm:"type:date;comment:统计日期" json:"stat_date"`        // 统计日期
	Duration  int64     `gorm:"not null;comment:本次统计的在线时长(秒)" json:"duration"`  // 在线时长（秒）
	StartTime int64     `gorm:"not null;comment:开始统计时间" json:"start_time"`      // 开始统计时间
	EndTime   int64     `gorm:"not null;comment:结束统计时间" json:"end_time"`        // 结束统计时间
	CreatedAt int64     `gorm:"default:0;comment:创建时间" json:"created_at"`       // 创建时间
}

// TableName 指定表名
func (*UserOnlineDurationFlow) TableName() string {
	return TableNameUserOnlineDurationFlow
}
