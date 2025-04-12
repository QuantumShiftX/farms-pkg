package pmodel

const TableNameGrowthRecord = "growth_records"

// GrowthRecord 成长值变动记录
type GrowthRecord struct {
	ID           int64  `gorm:"id" json:"id"`
	UserID       int64  `gorm:"user_id" json:"user_id"`
	ActionType   int8   `gorm:"action_type" json:"action_type"`     // '获取成长值的行为类型，1：存款、2：下级存款、3：收获、4：邀请下级、5：每日登录、6：在线时长';
	GrowthChange int64  `gorm:"growth_change" json:"growth_change"` // '成长值变动数量';
	Remark       string `gorm:"remark" json:"remark"`
	CreatedAt    int64  `gorm:"created_at" json:"created_at"`
}

// TableName GrowthRule's table name
func (*GrowthRecord) TableName() string {
	return TableNameGrowthRecord
}
