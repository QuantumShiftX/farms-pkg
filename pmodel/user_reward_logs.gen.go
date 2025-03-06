package pmodel

import (
	"time"

	"encoding/json"
)

const TableNameUserRewardLog = "user_reward_logs"

// UserRewardLog 结构体对应 user_reward_logs 表
type UserRewardLog struct {
	ID           int64           `gorm:"primaryKey;autoIncrement;column:id" json:"id"`                                                                 // 自增主键ID
	UserID       int64           `gorm:"not null;index:idx_user_reward_logs_user_id;column:user_id" json:"user_id"`                                    // 用户ID
	RewardType   int16           `gorm:"not null;default:2;index:idx_user_reward_logs_reward_type;column:reward_type" json:"reward_type"`              // 奖励类型
	RewardAmount int64           `gorm:"not null;column:reward_amount" json:"reward_amount"`                                                           // 奖励数量
	Source       string          `gorm:"type:varchar(100);index:idx_user_reward_logs_source;column:source" json:"source"`                              // 奖励来源
	Description  string          `gorm:"type:text;column:description" json:"description"`                                                              // 奖励描述
	CreatedAt    time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP;index:idx_user_reward_logs_created_at;column:created_at" json:"created_at"` // 记录创建时间
	Metadata     json.RawMessage `gorm:"type:jsonb;column:metadata" json:"metadata"`                                                                   // 额外属性 (JSON)
}

// TableName 设置表名
func (*UserRewardLog) TableName() string {
	return "user_reward_logs"
}
